package config

import (
	"github.com/segmentio/conf"
	"time"
)

const (
	HDBD     = "hdbd"
	HENGINED = "hengined"

	SupervisorName  = "top-supervisor"
	FrontActorName  = "front"
	WorkerActorName = "worker"

	WorkerTickTime = 1 * time.Second
)

type Items struct {
	LogLevel     string `conf:"log_level"`
	HdbdAddr     string `conf:"hdbd_addr" help:"listen addr" validate:"nonzero"`
	HenginedAddr string `conf:"hengined_addr" help:"listen addr" validate:"nonzero"`

	DbAddr string `conf:"db_addr" help:"monitor addr" validate:"nonzero"`
}

func Load() *Items {
	var config = &Items{}
	conf.Load(config)
	return config
}
