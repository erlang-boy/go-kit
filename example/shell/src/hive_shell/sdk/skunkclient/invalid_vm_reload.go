package skunkclient

import (
	"fmt"


	"automation/sdk"
)

type ReloadInvalidVmsParam struct {
	Address      string `json:"address"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Port      int `json:"port"`
}

func ReloadInvalidVms(skunkUrl string, p *ReloadInvalidVmsParam) error {
	client := sdk.NewClient(SkunkTimeout).SetRetryCount(1)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetBody(p).Post(skunkUrl)
	if err != nil {
		logger.Errorf("reload invalid vms addr %s failed, err is %+v, params is %+v", skunkUrl, err, p)
		return err
	}

  statusCode := resp.StatusCode()
  switch statusCode {
  case 200:
    return nil
  default:
		return fmt.Errorf("reload invalid vms %s failed, status code %d, param is %+v, response body is %s", skunkUrl, statusCode, p, resp)
  }
}
