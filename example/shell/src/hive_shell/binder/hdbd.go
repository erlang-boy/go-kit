package binder

import (
	"fmt"
	"hive_shell/sdk"
)

type Hdbd struct {
	startDbUrl   string
	stopDbUrl    string
	keepAliveUrl string
	statusUrl    string
}

const (
	KeepAlive = "/module/hive_db/keep_alive"
	StartDb   = "/module/hive_db/start"
	StopDb    = "/module/hive_db/stop"
	Status    = "/module/hive_db/get_status"
)

func NewHdbd(spec *BinderSpec) *Hdbd {
	return &Hdbd{
		startDbUrl:   makeUrl(spec.BindAddr, StartDb),
		stopDbUrl:    makeUrl(spec.BindAddr, StopDb),
		keepAliveUrl: makeUrl(spec.BindAddr, KeepAlive),
		statusUrl:    makeUrl(spec.BindAddr, Status),
	}
}

func (this *Hdbd) Start() error {

	resp, err := sdk.NewClient(BindTimeout).R().Post(this.startDbUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("start db failed, response code: %d", resp.StatusCode)
	}

	return nil
}

func (this *Hdbd) Stop() error {

	resp, err := sdk.NewClient(BindTimeout).R().Post(this.stopDbUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("start db failed, response code: %d", resp.StatusCode)
	}
	return nil
}

func (this *Hdbd) KeepAlive() error {

	resp, err := sdk.NewClient(BindTimeout).R().Post(this.keepAliveUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("start db failed, response code: %d", resp.StatusCode)
	}

	return nil
}

type resInfo struct {
	status string `json:"status"`
}

func (this *Hdbd) Status() (string, error) {
	res := &resInfo{}
	_, err := sdk.NewClient(BindTimeout).R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetResult(res).Post(this.statusUrl)

	if err != nil {
		return res.status, err
	}

	return res.status, nil
}

func makeUrl(addr, action string) string {
	return fmt.Sprintf("http://%s%s", addr, action)
}
