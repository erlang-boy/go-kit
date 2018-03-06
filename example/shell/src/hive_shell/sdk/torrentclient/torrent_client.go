package torrentclient

import (
	"fmt"
	"time"

	"github.com/automation/loggo"

	"automation/sdk"
)

var logger = loggo.GetLogger("automation.sdk.torrentclient")
var TorrentTimeout = 5 * time.Second

type NfsExportParam struct {
	Path       string   `json:"path"`
	Groups     []string `json:"groups"`
	AccessType string   `json:"access"`
}

func AddNfsExport(p *NfsExportParam, nfsUrl string) error {
	logger.Debugf("add nfs export, url:%s, params:%+v", nfsUrl, p)
	client := sdk.NewClient(TorrentTimeout)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).SetBody(p).Post(nfsUrl)
	if err != nil {
		logger.Errorf("addNfsExport addr %s failed, err is %+v, params is %+v", nfsUrl, err, p)
		return err
	}

	statusCode := resp.StatusCode()
	body := resp.String()

  // 409 means already has the same export, we treat it as success
	if statusCode == 200 || statusCode == 409 {
    return nil
  }

	if statusCode < 200 || statusCode >= 300 {
		logger.Errorf("addNfsExport %s failed, http status code is %d, params is %+v, response is %+v", nfsUrl, statusCode, p, body)
		return fmt.Errorf("addNfsExport %s failed, http status code is %d, params is %+v, response is %+v", nfsUrl, statusCode, p, body)
	}
	return nil
}

func RemoveNfsExport(nfsUrl string) error {
	logger.Debugf("remove nfs export, url:%s", nfsUrl)
	client := sdk.NewClient(TorrentTimeout)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).Post(nfsUrl)

	statusCode := resp.StatusCode()
	body := resp.String()
	if err != nil {
		logger.Errorf("RemoveNfsExport addr %s failed, err is %+v, status code is %d, response body is %s", nfsUrl, err, statusCode, body)
		return err
	}
	if statusCode == 404 || statusCode == 400 {
		logger.Infof("ignore %d, return success", statusCode)
		return nil
	}

	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("RemoveNfsExport %s failed, http status code is %d, response body is %s", nfsUrl, statusCode, body)
	}
	return nil
}

type NfsExports struct {
	Exports []*NfsExportParam   `json:"exports"`
}

func GetNfsExport(nfsUrl string) (*NfsExports, error) {
	logger.Debugf("Get nfs export, url:%s", nfsUrl)
	client := sdk.NewClient(TorrentTimeout)

    output := &NfsExports{}
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
		}).
        SetResult(output).
        Get(nfsUrl)

	statusCode := resp.StatusCode()
	body := resp.String()
	if err != nil {
		logger.Errorf("GetNfsExport addr %s failed, err is %+v, status code is %d, response body is %s", nfsUrl, err, statusCode, body)
		return nil, err
	}

	if statusCode < 200 || statusCode >= 300 {
		return nil, fmt.Errorf("GetNfsExport %s failed, http status code is %d, response body is %s", nfsUrl, statusCode, body)
	}
	return output, nil
}
