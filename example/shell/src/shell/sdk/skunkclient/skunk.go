package skunkclient

import (
  "errors"
	"fmt"
	"strings"
	"time"

	"github.com/automation/loggo"

	"automation/sdk"
)

var logger = loggo.GetLogger("automation.sdk.skunkclient")

const (
	CHANGE_HOST_ROUTE = "/vsphere/host/route"
	PAUSE_VMS         = "/vsphere/host/pausevm"
	RESUME_VM         = "/vsphere/host/resumevm"
)

var SkunkTimeout = 180 * time.Second

var (
  ErrDatastoreAlreadyExist error = errors.New("DatastoreAlreadyExisted")
)

type ChangeHostRouteParams struct {
	DataCenter string `json:"datacenter"`
	Host       string `json:"host"`
	Network    string `json:"network"`
	Gateway    string `json:"gateway"`
	Netmask    string `json:"netmask"`
}

func (c *ChangeHostRouteParams) SetDataCenter(dataCenter string) *ChangeHostRouteParams {
	c.DataCenter = dataCenter
	return c
}

func (c *ChangeHostRouteParams) SetHost(host string) *ChangeHostRouteParams {
	c.Host = host
	return c
}

func (c *ChangeHostRouteParams) SetNetwork(network string) *ChangeHostRouteParams {
	c.Network = network
	return c
}

func (c *ChangeHostRouteParams) SetGateway(gateway string) *ChangeHostRouteParams {
	c.Gateway = gateway
	return c
}

func (c *ChangeHostRouteParams) SetNetmask(netmask string) *ChangeHostRouteParams {
	c.Netmask = netmask
	return c
}

func ChangeHostRoute(endpoint string, params *ChangeHostRouteParams) (*ChangeHostRouteRes, error) {
	resInfo := &ChangeHostRouteRes{}
	client := sdk.NewClient(SkunkTimeout)
	endpoint = strings.TrimSuffix(endpoint, "/")
	url := endpoint + CHANGE_HOST_ROUTE
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/x-www-form-urlencoded",
		}).SetBody(params).SetResult(resInfo).Post(url)
	if err != nil {
		logger.Errorf("ChangeHostRoute => httpReq.Response failed, err is %s, params is %+v, status code is %d, response body is %s",
			err, params, resp.StatusCode(), resp)
	}
	return resInfo, err
}

type router struct {
	Network string `json:"network"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

type ChangeHostRouteRes struct {
	DataCenter string `json:"datacenter"`
	Host       string `json:"host"`
	Route      router `json:"route"`
}

type OptVmsParams struct {
	DataCenter string `json:"datacenter"`
	Host       string `json:"host"`
	NfsAddress string `json:"nfs_address"`
}

func OptVMs(endpoint, action string, params *OptVmsParams) error {
	var resInfo struct {
		Ok string `json:"success"`
	}
	endpoint = strings.TrimSuffix(endpoint, "/")
	url := endpoint + action

	client := sdk.NewClient(SkunkTimeout)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetBody(params).SetResult(resInfo).Post(url)

	if err != nil {
		logger.Errorf("OptVMs => httpReq.Response failed, err is %s, params is %+v, status code is %d, response body is %s",
			err, params, resp.StatusCode(), resp)
	}
	if resInfo.Ok != "true" {
		return fmt.Errorf("OptVms %s => parse result get false, res:%+v", action, resInfo)
	}
	return err
}

type CreateDatastoreParam struct {
	ProtocolType    string `json:"type"`
	Datastore       string `json:"datastore"`
	TargetHost      string `json:"target_host"`
	TargetPath      string `json:"target_path"`
	AccessMode      string `json:"access_mode"`
	EsxiHostAddress string `json:"esxi_host_address"`
}

func CreateDatastore(skunkUrl string, p *CreateDatastoreParam) error {
	client := sdk.NewClient(SkunkTimeout).SetRetryCount(1)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetBody(p).Post(skunkUrl)
	if err != nil {
		logger.Errorf("create datastore addr %s failed, err is %+v, params is %+v", skunkUrl, err, p)
		return err
	}

	statusCode := resp.StatusCode()
  switch statusCode {
  case 200:
    return nil
  case 409:
    return ErrDatastoreAlreadyExist
  default:
		return fmt.Errorf("create datastore %s failed, status code %d, param is %+v, response body is %s", skunkUrl, statusCode, p, resp)
  }
}

type RemoveDatastoreParam struct {
	Datastore       string `json:"datastore"`
	EsxiHostAddress string `json:"esxi_host_address"`
}

//notes: if hosts or datastore does not exist, return success
func RemoveDatastore(skunkUrl string, p *RemoveDatastoreParam) error {
	client := sdk.NewClient(SkunkTimeout).SetRetryCount(1)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetBody(p).Post(skunkUrl)
	if err != nil {
		logger.Errorf("RemoveDatastore addr %s failed, err is %+v, params is %+v, response body is %s", skunkUrl, err, p, resp.String())
		return err
	}

	statusCode := resp.StatusCode()
	if statusCode == 200 || statusCode == 404 {
		return nil
	} else {
		return fmt.Errorf("remove datastore %s failed, status code %d, param is %+v, response body is %s", skunkUrl, statusCode, p, resp)
	}

}
