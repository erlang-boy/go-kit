package master

import (
	"time"

	"github.com/juju/errors"
	"github.com/juju/loggo"
	"gopkg.in/tomb.v1"

	. "vega/go-kit/actor"
	"vega/go-kit/actor/registry"
)

// RestartDelay holds the length of time that a supervisor
// will wait between exiting and restarting.
var logger = loggo.GetLogger("actor.master")

const RestartDelay = 3 * time.Second

// master runs a set of supervisors, restarting them as necessary
// when they fail.
type master struct {
	tomb          tomb.Tomb
	startc        chan startReq
	stopc         chan string
	donec         chan doneInfo
	startedc      chan startInfo
	supervisors   map[string]*supervisorInfo
	isFatal       func(error) bool
	moreImportant func(err0, err1 error) bool

	// restartDelay holds the length of time that a supervisor
	// will wait between exiting and restarting.
	restartDelay time.Duration
}

type startReq struct {
	id    string
	start func() (Supervisor, error)
}

type startInfo struct {
	id         string
	supervisor Supervisor
}

type doneInfo struct {
	id  string
	err error
}

// NewMaster creates a new Master.  When a supervisor finishes, if its error
// is deemed fatal (determined by calling isFatal), all the other supervisors
// will be stopped and the master itself will finish.  Of all the fatal errors
// returned by the stopped supervisors, only the most important one,
// determined by calling moreImportant, will be returned from
// Master.Wait. Non-fatal errors will not be returned.
//
// The function isFatal(err) returns whether err is a fatal error.  The
// function moreImportant(err0, err1) returns whether err0 is considered
// more important than err1.
func NewMaster(isFatal func(error) bool, moreImportant func(err0, err1 error) bool, restartDelay time.Duration) Master {
	m := &master{
		startc:        make(chan startReq),
		stopc:         make(chan string),
		donec:         make(chan doneInfo),
		startedc:      make(chan startInfo),
		supervisors:   make(map[string]*supervisorInfo),
		isFatal:       isFatal,
		moreImportant: moreImportant,
		restartDelay:  restartDelay,
	}
	go func() {
		defer m.tomb.Done()
		m.tomb.Kill(m.run())
	}()
	return m
}

var ErrDead = errors.New("supervisor master is not running")

// StartSupervisor starts a supervisor running associated with the given id.
// The startFunc function will be called to create the supervisor;
// when the supervisor exits, it will be restarted as long as it
// does not return a fatal error.
//
// If there is already a supervisor with the given id, nothing will be done.
//
// StartSupervisor returns ErrDead if the master is not running.
func (m *master) StartSupervisor(id string, startFunc func() (Supervisor, error)) error {
	registry.Watch(id)
	select {
	case m.startc <- startReq{id, startFunc}:
		return nil
	case <-m.tomb.Dead():
	}
	return ErrDead
}

// StopSupervisor stops the supervisor associated with the given id.
// It does nothing if there is no such supervisor.
//
// StopSupervisor returns ErrDead if the master is not running.
func (m *master) StopSupervisor(id string) error {
	//fmt.Printf("master StoppSupervisor, id: %+v \n", id)
	select {
	case m.stopc <- id:
		return nil
	case <-m.tomb.Dead():
	}
	return ErrDead
}

func (m *master) Wait() error {
	return m.tomb.Wait()
}

func (m *master) Kill() {
	logger.Debugf("killing master %p", m)
	m.tomb.Kill(nil)
}

type supervisorInfo struct {
	start        func() (Supervisor, error)
	supervisor   Supervisor
	restartDelay time.Duration
	stopping     bool
}

func (m *master) run() error {
	// supervisors holds the current set of supervisors.  All supervisors with a
	// running goroutine have an entry here.
	//supervisors := make(map[string]*supervisorInfo)
	var finalError error

	// isDying holds whether the master is currently dying.  When it
	// is dying (whether as a result of being killed or due to a
	// fatal error), all existing supervisors are killed, no new supervisors
	// will be started, and the loop will exit when all existing
	// supervisors have stopped.
	isDying := false
	tombDying := m.tomb.Dying()
	for {
		if isDying && len(m.supervisors) == 0 {
			return finalError
		}
		select {
		case <-tombDying:
			logger.Infof("master is dying")
			isDying = true
			killAll(m.supervisors)
			tombDying = nil
		case req := <-m.startc:
			if isDying {
				logger.Infof("ignoring start request for %q when dying", req.id)
				break
			}
			info := m.supervisors[req.id]
			if info == nil {
				m.supervisors[req.id] = &supervisorInfo{
					start:        req.start,
					restartDelay: m.restartDelay,
				}
				go m.runSupervisor(0, req.id, req.start)
				break
			}
			if !info.stopping {
				// The supervisor is already running, so leave it alone
				break
			}
			// The supervisor previously existed and is
			// currently being stopped.  When it eventually
			// does stop, we'll restart it immediately with
			// the new start function.
			info.start = req.start
			info.restartDelay = 0
		case id := <-m.stopc:
			logger.Debugf("stop %q", id)
			if info := m.supervisors[id]; info != nil {
				//fmt.Printf("Master is killing Supervisor......... id = %s\n", id)
				killSupervisor(id, info)
			} else {
				//fmt.Printf("Supervisor not exist, id: %+v \n", id)
				registry := registry.SupervisorRegistry()
				registry.Delete(id)
			}
		case info := <-m.startedc:
			logger.Debugf("%q started", info.id)
			supervisorInfo := m.supervisors[info.id]
			supervisorInfo.supervisor = info.supervisor

			if isDying || supervisorInfo.stopping {
				killSupervisor(info.id, supervisorInfo)
			}

			registry := registry.SupervisorRegistry()
			registry.Set(info.id, info.supervisor)
		case info := <-m.donec:
			logger.Debugf("%q done: %v", info.id, info.err)
			supervisorInfo := m.supervisors[info.id]
			if !supervisorInfo.stopping && info.err == nil {
				logger.Debugf("removing %q from known supervisors", info.id)
				delete(m.supervisors, info.id)

				registry := registry.SupervisorRegistry()
				registry.Delete(info.id)

				break
			}
			if info.err != nil {
				if m.isFatal(info.err) {
					logger.Errorf("fatal %q: %v", info.id, info.err)
					if finalError == nil || m.moreImportant(info.err, finalError) {
						finalError = info.err
					}
					delete(m.supervisors, info.id)
					if !isDying {
						isDying = true
						killAll(m.supervisors)
					}
					break
				} else {
					logger.Errorf("exited %q: %v", info.id, info.err)
				}
			}
			if supervisorInfo.start == nil {
				logger.Debugf("no restart, removing %q from known supervisors", info.id)

				// The supervisor has been deliberately stopped;
				// we can now remove it from the list of supervisors.
				delete(m.supervisors, info.id)
				registry := registry.SupervisorRegistry()
				registry.Delete(info.id)
				break
			}
			go m.runSupervisor(supervisorInfo.restartDelay, info.id, supervisorInfo.start)
			supervisorInfo.restartDelay = m.restartDelay
		}
	}
}

func killAll(supervisors map[string]*supervisorInfo) {
	for id, info := range supervisors {
		killSupervisor(id, info)
	}
}

func killSupervisor(id string, info *supervisorInfo) {
	if info.supervisor != nil {
		logger.Debugf("killing %q", id)
		info.supervisor.Kill()
		info.supervisor = nil
	} else {
		logger.Debugf("couldn't kill %q, not yet started", id)
	}
	info.stopping = true
	info.start = nil
}

// runSupervisor starts the given supervisor after waiting for the given delay.
func (m *master) runSupervisor(delay time.Duration, id string, start func() (Supervisor, error)) {
	if delay > 0 {
		logger.Infof("restarting %q in %v", id, delay)
		select {
		case <-m.tomb.Dying():
			m.donec <- doneInfo{id, nil}
			return
		case <-time.After(delay):
		}
	}
	supervisor, err := start()
	if err == nil {
		m.startedc <- startInfo{id, supervisor}
		err = supervisor.Wait()
	}
	logger.Infof("stopped %q, err: %v", id, err)
	m.donec <- doneInfo{id, err}
}
