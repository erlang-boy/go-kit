package registry

import (
	"fmt"
	"github.com/juju/loggo"
	"sync"
	. "vega/go-kit/actor"
)

var logger = loggo.GetLogger("actor.registry")

type registry struct {
	supervisors map[string]Supervisor
	watchChans  map[string]*watchChan
	mu          sync.RWMutex
}

type watchChan struct {
	startedc chan bool
	deletedc chan bool
}

func (r *registry) Set(name string, supervisor Supervisor) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.supervisors[name] = supervisor
	//	watchChan := r.watchChans[name]
	//	watchChan.startedc <- true
}

func (r *registry) Get(name string) (Supervisor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	supervisor, ok := r.supervisors[name]
	if !ok {
		logger.Errorf("The '%s' is not presented", name)
		return nil, ErrActorNotFound
	}
	return supervisor, nil
}

func (r *registry) Delete(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.supervisors, name)

	watchChan, ok := r.watchChans[name]
	if !ok {
		return nil
	}
	close(watchChan.deletedc)
	delete(r.watchChans, name)

	return nil
}

func (r *registry) List() ([]Supervisor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res := make([]Supervisor, len(r.supervisors))
	idx := 0
	for _, supervisor := range r.supervisors {
		res[idx] = supervisor
		idx++
	}

	return res, nil
}

func (r *registry) WatchChan(supervisorName string) (*watchChan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ch, ok := r.watchChans[supervisorName]
	if !ok {
		logger.Errorf("The '%s' watch chan is not exist", supervisorName)
		return nil, ErrActorWatchChanNotFound
	}
	return ch, nil
}

func (r *registry) Watch(supervisorName string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	w := &watchChan{
		startedc: make(chan bool, 1),
		deletedc: make(chan bool, 1),
	}
	r.watchChans[supervisorName] = w
}

var (
	r    *registry
	once sync.Once
)

func SupervisorRegistry() *registry {
	once.Do(func() {
		r = &registry{
			supervisors: make(map[string]Supervisor),
			watchChans:  make(map[string]*watchChan),
		}
	})

	return r
}

func Whereis(supervisorName, actorName string) (Actor, error) {
	register := SupervisorRegistry()
	supervisor, err := register.Get(supervisorName)
	if err != nil {
		return nil, fmt.Errorf("supervisor %s not found", supervisorName)
	}

	actor, err := supervisor.Child(actorName)
	if err != nil {
		return nil, fmt.Errorf("actor %s not found", actorName)
	}

	return actor, nil
}

func Watch(supervisorName string) {
	register := SupervisorRegistry()
	register.Watch(supervisorName)
}

func StartedChan(supervisorName string) (chan bool, error) {
	register := SupervisorRegistry()
	watchChan, err := register.WatchChan(supervisorName)
	if err != nil {
		return nil, err
	}
	return watchChan.startedc, nil
}

func DeletedChan(supervisorName string) (chan bool, error) {
	register := SupervisorRegistry()
	watchChan, err := register.WatchChan(supervisorName)
	if err != nil {
		return nil, err
	}
	return watchChan.deletedc, nil
}
