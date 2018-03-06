package supervisor

import (
	"math/rand"
	"time"

	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/utils/set"
	"gopkg.in/tomb.v1"

	. "vega/go-kit/actor"
)

var logger = loggo.GetLogger("actor.supervisor")

type SupervisorConfig struct {
	IsFatal     IsFatalFunc
	WorstError  WorstErrorFunc
	Filter      FilterFunc
	ErrorDelay  time.Duration
	BounceDelay time.Duration
}

func (config *SupervisorConfig) Validate() error {
	if config.IsFatal == nil {
		return errors.New("IsFatal not specified")
	}
	if config.WorstError == nil {
		return errors.New("WorstError not specified")
	}
	if config.ErrorDelay < 0 {
		return errors.New("ErrorDelay is negative")
	}
	if config.BounceDelay < 0 {
		return errors.New("BounceDelay is negative")
	}
	return nil
}

func NewSupervisorImpl(config SupervisorConfig) (Supervisor, error) {
	if err := config.Validate(); err != nil {
		return nil, errors.Annotatef(err, "invalid config")
	}
	s := &SupervisorImpl{
		config: config,

		manifolds:  Manifolds{},
		dependents: map[string][]string{},
		current:    map[string]actorInfo{},

		install:     make(chan installTicket),
		started:     make(chan startedTicket),
		stopped:     make(chan stoppedTicket),
		report:      make(chan reportTicket),
		ReqCurrentC: make(chan ReqCurrent),
	}
	go func() {
		defer s.tomb.Done()
		s.tomb.Kill(s.loop())
	}()
	return s, nil
}

type SupervisorImpl struct {
	config     SupervisorConfig
	tomb       tomb.Tomb
	worstError error
	manifolds  Manifolds

	// dependents holds, for each named manifold, those that depend on it.
	dependents map[string][]string

	// current holds the active actor information for each installed manifold.
	current map[string]actorInfo

	// install, started, report and stopped each communicate requests and changes into
	// the loop goroutine.
	install chan installTicket
	started chan startedTicket
	stopped chan stoppedTicket
	report  chan reportTicket

	ReqCurrentC chan ReqCurrent
}

type ReqCurrent struct {
	Done chan struct{}
}

func (s *SupervisorImpl) loop() error {
	oneShotDying := s.tomb.Dying()
	for {
		select {
		case <-oneShotDying:
			oneShotDying = nil
			for name := range s.current {
				s.requestStop(name)
			}
		case ticket := <-s.report:
			// This is safe so long as the Report method reads the result.
			ticket.result <- s.liveReport()
		case ticket := <-s.install:
			// This is safe so long as the Install method reads the result.
			ticket.result <- s.gotInstall(ticket.name, ticket.manifold)
		case ticket := <-s.started:
			s.gotStarted(ticket.name, ticket.actor, ticket.resourceLog)
		case ticket := <-s.stopped:
			s.gotStopped(ticket.name, ticket.error, ticket.resourceLog)
		case req := <-s.ReqCurrentC:
			<-req.Done
		}
		if s.isDying() {
			if s.allOthersStopped() {
				return tomb.ErrDying
			}
		}
	}
}

func (s *SupervisorImpl) Kill() {
	s.tomb.Kill(nil)
}

func (s *SupervisorImpl) Wait() error {
	if tombError := s.tomb.Wait(); tombError != nil {
		return tombError
	}
	err := s.worstError
	if s.config.Filter != nil {
		return s.config.Filter(err)
	}
	return err
}

func (s *SupervisorImpl) Report() map[string]interface{} {
	report := make(chan map[string]interface{})
	select {
	case s.report <- reportTicket{report}:
		// This is safe so long as the loop sends a result.
		return <-report
	case <-s.tomb.Dead():
		// Note that we don't abort on Dying as we usually would; the
		// oneShotDying approach in loop means that it can continue to
		// process requests until the last possible moment. Only once
		// loop has exited do we fall back to this report.
		report := map[string]interface{}{
			KeyState:     "stopped",
			KeyManifolds: s.manifoldsReport(),
		}
		if err := s.Wait(); err != nil {
			report[KeyError] = err.Error()
		}
		return report
	}
}

func (s *SupervisorImpl) liveReport() map[string]interface{} {
	var reportError error
	state := "started"
	if s.isDying() {
		state = "stopping"
		if tombError := s.tomb.Err(); tombError != nil {
			reportError = tombError
		} else {
			reportError = s.worstError
		}
	}
	report := map[string]interface{}{
		KeyState:     state,
		KeyManifolds: s.manifoldsReport(),
	}
	if reportError != nil {
		report[KeyError] = reportError.Error()
	}
	return report
}

func (s *SupervisorImpl) manifoldsReport() map[string]interface{} {
	manifolds := map[string]interface{}{}
	for name, info := range s.current {
		report := map[string]interface{}{
			KeyState:       info.state(),
			KeyInputs:      s.manifolds[name].Inputs,
			KeyResourceLog: resourceLogReport(info.resourceLog),
		}
		if info.err != nil {
			report[KeyError] = info.err.Error()
		}
		if reporter, ok := info.actor.(Reporter); ok {
			if reporter != s {
				report[KeyReport] = reporter.Report()
			}
		}
		manifolds[name] = report
	}
	return manifolds
}

func (s *SupervisorImpl) Install(name string, manifold Manifold) error {
	result := make(chan error)
	select {
	case <-s.tomb.Dying():
		return errors.New("supervisor is shutting down")
	case s.install <- installTicket{name, manifold, result}:
		// This is safe so long as the loop sends a result.
		return <-result
	}
}

func (s *SupervisorImpl) Child(name string) (Actor, error) {
	done := make(chan struct{})
	defer close(done)
	req := ReqCurrent{
		Done: done,
	}

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	select {
	case s.ReqCurrentC <- req:
		if info, found := s.current[name]; found && info.actor != nil {
			return info.actor, nil
		}
		return nil, errors.New("child not found")
	case <-timer.C:
		return nil, errors.New("get child timeout")
	}
}

func (s *SupervisorImpl) gotInstall(name string, manifold Manifold) error {
	logger.Tracef("installing %q manifold...", name)
	if _, found := s.manifolds[name]; found {
		return errors.Errorf("%q manifold already installed", name)
	}
	if err := s.checkAcyclic(name, manifold); err != nil {
		return errors.Annotatef(err, "cannot install %q manifold", name)
	}
	s.manifolds[name] = manifold
	for _, input := range manifold.Inputs {
		s.dependents[input] = append(s.dependents[input], name)
	}
	s.current[name] = actorInfo{}
	s.requestStart(name, 0)
	return nil
}

func (s *SupervisorImpl) uninstall(name string) {
	for dName, dependents := range s.dependents {
		depSet := set.NewStrings(dependents...)
		depSet.Remove(name)
		s.dependents[dName] = depSet.Values()
	}
	delete(s.current, name)
	delete(s.manifolds, name)
}

func (s *SupervisorImpl) checkAcyclic(name string, manifold Manifold) error {
	manifolds := Manifolds{name: manifold}
	for name, manifold := range s.manifolds {
		manifolds[name] = manifold
	}
	return Validate(manifolds)
}

func (s *SupervisorImpl) requestStart(name string, delay time.Duration) {
	// Check preconditions.
	manifold, found := s.manifolds[name]
	if !found {
		s.tomb.Kill(errors.Errorf("fatal: unknown manifold %q", name))
	}

	// Copy current info and check more preconditions.
	info := s.current[name]
	if !info.stopped() {
		s.tomb.Kill(errors.Errorf("fatal: trying to start a second %q manifold actor", name))
	}

	// Final check that we're not shutting down yet...
	if s.isDying() {
		logger.Tracef("not starting %q manifold actor (shutting down)", name)
		return
	}

	info.starting = true
	info.abort = make(chan struct{})
	s.current[name] = info
	context := s.context(name, manifold.Inputs, info.abort)

	// Always fuzz the delay a bit to help randomise the order of actors starting,
	// which should make bugs more obvious
	if delay > time.Duration(0) {
		delay += time.Duration(rand.Int31n(60)) * time.Millisecond
	}

	go s.runActor(name, delay, manifold.Start, context)
}

func (s *SupervisorImpl) context(name string, inputs []string, abort <-chan struct{}) *context {
	outputs := map[string]OutputFunc{}
	actors := map[string]Actor{}
	for _, resourceName := range inputs {
		outputs[resourceName] = s.manifolds[resourceName].Output
		actors[resourceName] = s.current[resourceName].actor
	}
	return &context{
		clientName: name,
		abort:      abort,
		expired:    make(chan struct{}),
		actors:     actors,
		outputs:    outputs,
	}
}

func (s *SupervisorImpl) runActor(name string, delay time.Duration, start StartFunc, context *context) {

	errAborted := errors.New("aborted before delay elapsed")

	startAfterDelay := func() (Actor, error) {
		// NOTE: the context will expire *after* the actor is started.
		// This is tolerable because
		//  1) we'll still correctly block access attempts most of the time
		//  2) failing to block them won't cause data races anyway
		//  3) it's not worth complicating the interface for every client just
		//     to eliminate the possibility of one harmlessly dumb interaction.
		defer context.expire()
		logger.Tracef("starting %q manifold actor in %s...", name, delay)
		select {
		case <-s.tomb.Dying():
			return nil, errAborted
		case <-context.Abort():
			return nil, errAborted
		// TODO(fwereade): 2016-03-17 lp:1558657
		case <-time.After(delay):
		}
		logger.Tracef("starting %q manifold actor", name)
		return start(context)
	}

	startActorAndWait := func() error {
		actor, err := startAfterDelay()
		switch errors.Cause(err) {
		case errAborted:
			return nil
		case nil:
			logger.Tracef("running %q manifold actor", name)
		default:
			logger.Tracef("failed to start %q manifold actor: %v", name, err)
			return err
		}
		select {
		case <-s.tomb.Dying():
			logger.Tracef("stopping %q manifold actor (shutting down)", name)
			// Doesn't matter whether actor == s: if we're already Dying
			// then cleanly Kill()ing ourselves again won't hurt anything.
			actor.Kill()
		case s.started <- startedTicket{name, actor, context.accessLog}:
			logger.Tracef("registered %q manifold actor", name)
		}
		if actor == s {
			// We mustn't Wait() for ourselves to complete here, or we'll
			// deadlock. But we should wait until we're Dying, because we
			// need this func to keep running to keep the self manifold
			// accessible as a resource.
			<-s.tomb.Dying()
			return tomb.ErrDying
		}

		return actor.Wait()
	}

	// We may or may not send on started, but we *must* send on stopped.
	s.stopped <- stoppedTicket{name, startActorAndWait(), context.accessLog}
}

func (s *SupervisorImpl) gotStarted(name string, actor Actor, resourceLog []resourceAccess) {
	// Copy current info; check preconditions and abort the actors if we've
	// already been asked to stop it.
	info := s.current[name]
	switch {
	case info.actor != nil:
		s.tomb.Kill(errors.Errorf("fatal: unexpected %q manifold actor start", name))
		fallthrough
	case info.stopping, s.isDying():
		logger.Tracef("%q manifold actor no longer required", name)
		actor.Kill()
	default:
		// It's fine to use this actor; update info and copy back.
		logger.Debugf("%q manifold actor started", name)
		s.current[name] = actorInfo{
			actor:       actor,
			resourceLog: resourceLog,
		}
		// Any manifold that declares this one as an input needs to be restarted.
		s.bounceDependents(name)
	}
}

func (s *SupervisorImpl) gotStopped(name string, err error, resourceLog []resourceAccess) {
	logger.Debugf("%q manifold actor stopped: %v", name, err)
	if filter := s.manifolds[name].Filter; filter != nil {
		err = filter(err)
	}

	// Copy current info and check for reasons to stop the s.
	info := s.current[name]
	if info.stopped() {
		s.tomb.Kill(errors.Errorf("fatal: unexpected %q manifold actor stop", name))
	} else if s.config.IsFatal(err) {
		s.worstError = s.config.WorstError(err, s.worstError)
		s.tomb.Kill(nil)
	}

	// Reset s info; and bail out if we can be sure there's no need to bounce.
	s.current[name] = actorInfo{
		err:         err,
		resourceLog: resourceLog,
	}
	if s.isDying() {
		logger.Tracef("permanently stopped %q manifold actor (shutting down)", name)
		return
	}

	// If we told the actor to stop, we should start it again immediately,
	// whatever else happened.
	if info.stopping {
		s.requestStart(name, s.config.BounceDelay)
	} else {
		// If we didn't stop it ourselves, we need to interpret the error.
		switch errors.Cause(err) {
		case nil:
			// Nothing went wrong; the task completed successfully. Nothing
			// needs to be done (unless the inputs change, in which case it
			// gets to check again).
		case ErrMissing:
			// The task can't even start with the current state. Nothing more
			// can be done (until the inputs change, in which case we retry
			// anyway).
		case ErrBounce:
			// The task exited but wanted to restart immediately.
			s.requestStart(name, s.config.BounceDelay)
		case ErrUninstall:
			// The task should never run again, and can be removed completely.
			s.uninstall(name)
		default:
			// Something went wrong but we don't know what. Try again soon.
			logger.Errorf("%q manifold actor returned unexpected error: %v", name, err)
			s.requestStart(name, s.config.ErrorDelay)
		}
	}

	// Manifolds that declared a dependency on this one only need to be notified
	// if the actor has changed; if it was already nil, nobody needs to know.
	if info.actor != nil {
		s.bounceDependents(name)
	}
}

func (s *SupervisorImpl) requestStop(name string) {
	// If already stopping or stopped, just don't do anything.
	info := s.current[name]
	if info.stopping || info.stopped() {
		return
	}

	// Update info, kill actor if present, and copy info back to s.
	info.stopping = true
	if info.abort != nil {
		close(info.abort)
		info.abort = nil
	}
	if info.actor != nil {
		info.actor.Kill()
	}
	s.current[name] = info
}

func (s *SupervisorImpl) isDying() bool {
	select {
	case <-s.tomb.Dying():
		return true
	default:
		return false
	}
}

func (s *SupervisorImpl) allOthersStopped() bool {
	for _, info := range s.current {
		if !info.stopped() && info.actor != s {
			return false
		}
	}
	return true
}

func (s *SupervisorImpl) bounceDependents(name string) {
	logger.Tracef("restarting dependents of %q manifold", name)
	for _, dependentName := range s.dependents[name] {
		if s.current[dependentName].stopped() {
			s.requestStart(dependentName, s.config.BounceDelay)
		} else {
			s.requestStop(dependentName)
		}
	}
}

type actorInfo struct {
	starting    bool
	stopping    bool
	abort       chan struct{}
	actor       Actor
	err         error
	resourceLog []resourceAccess
}

func (info actorInfo) stopped() bool {
	switch {
	case info.actor != nil:
		return false
	case info.starting:
		return false
	}
	return true
}

func (info actorInfo) state() string {
	switch {
	case info.starting:
		return "starting"
	case info.stopping:
		return "stopping"
	case info.actor != nil:
		return "started"
	}
	return "stopped"
}

type installTicket struct {
	name     string
	manifold Manifold
	result   chan<- error
}

type startedTicket struct {
	name        string
	actor       Actor
	resourceLog []resourceAccess
}

type stoppedTicket struct {
	name        string
	error       error
	resourceLog []resourceAccess
}

type reportTicket struct {
	result chan map[string]interface{}
}
