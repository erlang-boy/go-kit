package binder

import (
	"fmt"
	"hive_shell/sdk"
)

type HEngined struct {
	startDbUrl   string
	stopDbUrl    string
	keepAliveUrl string
	statusUrl    string
}

const (
	keepAlive = "/module/hive_engine/keep_alive"
	startDb   = "/module/hive_engine/start"
	stopDb    = "/module/hive_engine/stop"
	status    = "/module/hive_engine/get_status"
)

func NewHEngined(spec *BinderSpec) *HEngined {
	return &HEngined{
		startDbUrl:   makeUrl(spec.BindAddr, startDb),
		stopDbUrl:    makeUrl(spec.BindAddr, stopDb),
		keepAliveUrl: makeUrl(spec.BindAddr, keepAlive),
		statusUrl:    makeUrl(spec.BindAddr, status),
	}
}

func (this *HEngined) Start() error {

	resp, err := sdk.NewClient(BindTimeout).R().Post(this.startDbUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("start db failed, response code: %d", resp.StatusCode)
	}

	return nil
}

func (this *HEngined) Stop() error {

	resp, err := sdk.NewClient(BindTimeout).R().Post(this.stopDbUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("start db failed, response code: %d", resp.StatusCode)
	}
	return nil
}

func (this *HEngined) KeepAlive() error {

	resp, err := sdk.NewClient(BindTimeout).R().Post(this.keepAliveUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("start db failed, response code: %d", resp.StatusCode)
	}
	return nil
}

type statusInfo struct {
	status string `json:"status"`
}

func (this *HEngined) Status() (string, error) {
	res := &statusInfo{}
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
