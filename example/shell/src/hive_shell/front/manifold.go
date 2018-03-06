package front

import (
	"github.com/juju/errors"
	"vega/go-kit/actor"
)

type ManifoldConfig struct {
	ShellName  string
	ListenAddr string
}

func Manifold(config ManifoldConfig) actor.Manifold {
	return actor.Manifold{
		Inputs: []string{},
		Start: func(context actor.Context) (actor.Actor, error) {

			// Collect all required resources.
			//			var r daemon.Daemon
			//			if err := context.Get(config.DaemonName, &r); err != nil {
			//				return nil, err
			//			}

			httpd, err := NewHttpd(&HttpdParams{
				ShellName:  config.ShellName,
				ListenAddr: config.ListenAddr,
			})

			if err != nil {
				return nil, errors.Trace(err)
			}
			return httpd, nil
		},
	}
}
