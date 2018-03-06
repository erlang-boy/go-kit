package skunkclient

import (
	"fmt"


	"automation/sdk"
)

type ListHostDatastoresParam struct {
	Datacenter  string       `json:"datacenter"`
	HostIp      string       `json:"host_ip"`
}

type HostDatastore struct {
    HostIp string            `json:"host_address"`
    Datastores []Datastore   `json:"datastores"`
}

type Datastore struct {
    Name string              `json:"name"`
    Type string              `json:"type"`
    RemoteHost string        `json:"remote_host"`
    RemotePath string        `json:"remote_path"`
}

func ListHostDatastores(skunkUrl string, p *ListHostDatastoresParam) (*HostDatastore, error) {
	client := sdk.NewClient(SkunkTimeout).SetRetryCount(1)

    output := &HostDatastore{}
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetBody(p).SetResult(output).Post(skunkUrl)

	if err != nil {
		logger.Errorf("listHostDatastores addr %s failed, err is %+v, params is %+v", skunkUrl, err, p)
		return nil, err
	}

  statusCode := resp.StatusCode()
  switch statusCode {
  case 200:
    return output, nil
  default:
		return nil, fmt.Errorf("listHostDatastores %s failed, status code %d, param is %+v, response body is %s", skunkUrl, statusCode, p, resp)
  }
}
