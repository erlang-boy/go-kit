package masterDaemon

import (
	"shell/config"
	"shell/front"
	"shell/worker"

	"go-kit/actor"
)

const (
	frontName  = config.FrontActorName
	workerName = config.WorkerActorName
)

type ManifoldsConfig struct {
	ConfigItems *config.Items
	ShellName   string
}

func Manifolds(conf ManifoldsConfig) actor.Manifolds {

	var listenAddr string

	switch conf.ShellName {
	case config.HDBD:
		listenAddr = conf.ConfigItems.HdbdAddr
	case config.HENGINED:
		listenAddr = conf.ConfigItems.HenginedAddr
	}

	dbAddr := conf.ConfigItems.DbAddr

	return actor.Manifolds{

		// worker
		workerName: worker.Manifold(worker.ManifoldConfig{
			ShellName:   conf.ShellName,
			MonitorAddr: dbAddr,
		}),

		// front
		frontName: front.Manifold(front.ManifoldConfig{
			ShellName:  conf.ShellName,
			ListenAddr: listenAddr,
		}),
	}
}
