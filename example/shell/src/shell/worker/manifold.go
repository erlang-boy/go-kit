package worker

import (
	"github.com/juju/errors"
	"go-kit/actor"
)

type ManifoldConfig struct {
	ShellName   string
	MonitorAddr string
}

func Manifold(config ManifoldConfig) actor.Manifold {
	return actor.Manifold{
		Inputs: []string{},
		Start: func(context actor.Context) (actor.Actor, error) {

			worker, err := NewWorker(&WorkerParams{
				ShellName:   config.ShellName,
				MonitorAddr: config.MonitorAddr,
			})

			if err != nil {
				return nil, errors.Trace(err)
			}
			return worker, nil
		},
	}
}
