package hiveclient

import (
	"fmt"
	"time"

	"github.com/automation/loggo"

	"automation/sdk"
)

var logger = loggo.GetLogger("automation.sdk.hiveclient")

var (
	HiveTimeout = 5 * time.Second

	DESTROY_VOLUME_DRIVER       = "/hive_service/destroy_volume_driver"
	DESTROY_VOLUME_DIRVER_STATS = "/hive_service/destroy_volume_driver_stats"
)

//address: hive ip address
//key is string "X-Volume-Ids" when do iopath
//value is a string
func DestroyVolumeDriver(address, key, value string) error {
	endpoint := fmt.Sprintf("http://%s:9041%s", address, DESTROY_VOLUME_DRIVER)
	logger.Debugf("request to hive: endpoint: %s, key: %s, value: %s", endpoint, key, value)

	client := sdk.NewClient(HiveTimeout)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
			key:            value,
		}).Post(endpoint)
	if err != nil {
		logger.Errorf("destroy volume driver failed, error: %s", err)
		return err
	}
	if resp.StatusCode() != 200 {
		logger.Errorf("destroy volume driver failed: %s, params header is %s:%s, status code is %d, response body: %s", err, key, value, resp.StatusCode(), resp.String())
		return fmt.Errorf("destroy volume driver failed, resp code is %d", resp.StatusCode())
	}
	return nil
}

type ResInfo struct {
	Destroyed bool `json:"destroyed"`
}

func DestroyVolumeDriverStatus(address, key, value string) (bool, error) {
	endpoint := fmt.Sprintf("http://%s:9041%s", address, DESTROY_VOLUME_DIRVER_STATS)
	logger.Debugf("request to hive: endpoint: %s, key: %s, value: %s", endpoint, key, value)

	res := &ResInfo{}
	client := sdk.NewClient(HiveTimeout)
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Accept":       "*/*",
			"Content-Type": "application/json",
			key:            value,
		}).SetResult(res).Post(endpoint)

	if err != nil {
		logger.Errorf("destroy volume driver status failed: %s, params header is %s:%s, status code is %d, response body: %s", err, key, value, resp.StatusCode(), resp.String())
		return false, err
	}

	if res.Destroyed != true {
		return false, nil
	}

	return true, nil
}
