package supervisor

import (
	"github.com/juju/errors"

	. "go-kit/actor"
)

// Installer exposes an Supervisor's Install method.
type Installer interface {
	Install(name string, manifold Manifold) error
}

// Install is a convenience function for installing multiple manifolds into an
// Installer at once. It returns the first error it encounters (and installs no
// more manifolds).
func Install(installer Installer, manifolds Manifolds) error {
	for name, manifold := range manifolds {
		if err := installer.Install(name, manifold); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// Validate will return an error if the dependency graph defined by the supplied
// manifolds contains any cycles.
func Validate(manifolds Manifolds) error {
	inputs := make(map[string][]string)
	for name, manifold := range manifolds {
		inputs[name] = manifold.Inputs
	}
	return validator{
		inputs: inputs,
		doing:  make(map[string]bool),
		done:   make(map[string]bool),
	}.run()
}

// validator implements a topological sort of the nodes defined in inputs; it
// doesn't actually produce sorted nodes, but rather exists to return an error
// if it determines that the nodes cannot be sorted (and hence a cycle exists).
type validator struct {
	inputs map[string][]string
	doing  map[string]bool
	done   map[string]bool
}

func (v validator) run() error {
	for node := range v.inputs {
		if err := v.visit(node); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (v validator) visit(node string) error {
	if v.doing[node] {
		return errors.Errorf("cycle detected at %q (considering: %v)", node, v.doing)
	}
	if !v.done[node] {
		v.doing[node] = true
		for _, input := range v.inputs[node] {
			if err := v.visit(input); err != nil {
				// Tracing this error will not help anyone.
				return err
			}
		}
		v.done[node] = true
		v.doing[node] = false
	}
	return nil
}

// SelfManifold returns a manifold exposing a running dependency supervisor's
// Installer and Reporter services. The returned manifold is intended for
// installation into the supervisor it wraps; installing it into other supervisors
// may have surprising effects.
func SelfManifold(supervisor Supervisor) Manifold {
	return Manifold{
		Start: func(_ Context) (Actor, error) {
			return supervisor, nil
		},
		Output: func(in Actor, out interface{}) error {
			if in != supervisor {
				return errors.New("unexpected input actor")
			}
			switch outPtr := out.(type) {
			case *Installer:
				*outPtr = supervisor
			case *Reporter:
				*outPtr = supervisor
			default:
				return errors.Errorf("out should be a *Installer or a *Reporter; is %#v", out)
			}
			return nil
		},
	}
}
