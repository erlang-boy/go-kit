package worker

import (
	"github.com/juju/loggo"
	"gopkg.in/tomb.v1"
	"hive_shell/binder"
	"hive_shell/config"
	"time"
)

var (
	logger = loggo.GetLogger("hive_shell.worker")
)

type WorkerParams struct {
	ShellName   string
	MonitorAddr string
}

type Worker struct {
	Binder binder.Binder
	tomb   tomb.Tomb
	stop   chan bool
}

func NewWorker(workerParams *WorkerParams) (*Worker, error) {

	d := &Worker{
		Binder: binder.NewBinder(&binder.BinderSpec{
			BindServerName: workerParams.ShellName,
			BindAddr:       workerParams.MonitorAddr,
		}),
	}

	go func() {
		defer d.tomb.Done()
		d.tomb.Kill(d.loop())
	}()

	return d, nil
}

func (d *Worker) loop() (err error) {

	// 1. start db server
	err = d.Binder.Start()
	if err != nil {
		logger.Errorf("start db server failed, reason: %+v", err)
		return err
	}

	tick := time.NewTicker(config.WorkerTickTime)

	// 2. loop keep alive
	for {
		select {
		case <-tick.C:
			err = d.Binder.KeepAlive()
			if err != nil {
				return err
			}

		case <-d.stop:
			tick.Stop()
			return nil
		}
	}
}

func (d *Worker) Start() error {
	return d.Binder.Start()
}

func (d *Worker) Stop() error {
	// not block util db stop
	err := d.Binder.Stop()
	d.stop <- true
	return err
}

func (d *Worker) KeepAlive() error {
	return nil
}

func (d *Worker) Status() (string, error) {
	return d.Binder.Status()
}

func (d *Worker) Kill() {
	d.tomb.Kill(nil)
}

func (d *Worker) Wait() error {
	return d.tomb.Wait()
}
