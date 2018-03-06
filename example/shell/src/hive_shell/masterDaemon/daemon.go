package masterDaemon

import (
	"github.com/juju/loggo"
	"gopkg.in/tomb.v1"
	"hive_shell/config"
	"net/http"
	"time"

	"vega/go-kit/actor"
	"vega/go-kit/actor/master"
	"vega/go-kit/actor/supervisor"
	cmdutil "vega/go-kit/util"
)

const SupervisorName = config.SupervisorName

var (
	logger          = loggo.GetLogger("hive_shell.daemon")
	daemonManifolds = Manifolds
)

type Daemon struct {
	tomb      *tomb.Tomb
	master    actor.Master
	conf      *config.Items
	shellName string
}

func NewDaemon(shellName string, conf *config.Items) *Daemon {
	// todo new log handler
	return &Daemon{
		tomb:      new(tomb.Tomb),
		master:    master.NewMaster(cmdutil.IsFatal, cmdutil.MoreImportant, master.RestartDelay),
		conf:      conf,
		shellName: shellName,
	}
}

func (c *Daemon) Run() error {

	// For debug goroutine, http://172.16.24.100:6061/debug/pprof
	go func() {
		logger.Infof("%s", http.ListenAndServe("0.0.0.0:6061", nil))
	}()

	defer c.tomb.Done()

	c.master.StartSupervisor(SupervisorName, c.DaemonSupervisors)
	err := c.master.Wait()
	c.tomb.Kill(err)
	return err
}

func (c *Daemon) Stop() error {
	c.master.Kill()
	return c.tomb.Wait()
}

func (c *Daemon) DaemonSupervisors() (actor.Supervisor, error) {
	superConfig := supervisor.SupervisorConfig{
		IsFatal:     cmdutil.IsFatal,
		WorstError:  cmdutil.MoreImportantError,
		ErrorDelay:  3 * time.Second,
		BounceDelay: 10 * time.Millisecond,
	}

	sup, err := supervisor.NewSupervisorImpl(superConfig)
	if err != nil {
		return nil, err
	}

	manifolds := daemonManifolds(ManifoldsConfig{
		ConfigItems: c.conf,
		ShellName:   c.shellName,
	})

	if err := supervisor.Install(sup, manifolds); err != nil {
		logger.Debugf("failed with installing manifold to supervisor \n")
		if err := actor.Stop(sup); err != nil {
			logger.Errorf("while stopping supervisor with bad manifolds: %v", err)
		}
		return nil, err
	}
	return sup, nil
}
