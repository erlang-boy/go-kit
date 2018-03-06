package front

import (
	"github.com/juju/loggo"
	"gopkg.in/tomb.v1"
)

var logger = loggo.GetLogger("hive_shell.front")

type HttpdParams struct {
	ShellName  string
	ListenAddr string
}

type Httpd struct {
	Config *HttpdParams
	tomb   tomb.Tomb
}

func NewHttpd(httpdParams *HttpdParams) (*Httpd, error) {
	d := &Httpd{
		Config: httpdParams,
	}

	go func() {
		defer d.tomb.Done()
		d.tomb.Kill(d.loop())
	}()

	return d, nil
}

func (d *Httpd) loop() (err error) {

	router := NewRouter(d.Config)
	router.Run(d.Config.ListenAddr)
	return nil
}

func (d *Httpd) Kill() {
	d.tomb.Kill(nil)
}

func (d *Httpd) Wait() error {
	return d.tomb.Wait()
}
