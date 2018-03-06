package actor

import (
	"github.com/juju/errors"
)

type Actor interface {
	Kill()
	Wait() error
}

type Supervisor interface {
	Actor

	Install(name string, manifold Manifold) error
	Report() map[string]interface{}
	Child(name string) (Actor, error)
}

type Master interface {
	Actor

	StartSupervisor(id string, startFunc func() (Supervisor, error)) error
	StopSupervisor(id string) error
}

type Manifold struct {
	Inputs []string
	Start  StartFunc
	Filter FilterFunc
	Output OutputFunc
}

type Manifolds map[string]Manifold

type Context interface {
	Abort() <-chan struct{}
	Get(name string, out interface{}) error
}

type StartFunc func(context Context) (Actor, error)
type FilterFunc func(error) error
type OutputFunc func(in Actor, out interface{}) error
type IsFatalFunc func(err error) bool
type WorstErrorFunc func(err0, err1 error) error

var ErrMissing = errors.New("dependency not available")
var ErrBounce = errors.New("restart immediately")
var ErrUninstall = errors.New("resource permanently unavailable")
