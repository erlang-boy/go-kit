package guruclient

import (
	"fmt"
	"time"

	"github.com/go-resty/resty"

	"automation/sdk"
)

var GuruTimeout = 15 * time.Second

type GURU struct {
	guruAddr string
	client   *resty.Client
}

func NewGURU(addr string) *GURU {
	return &GURU{
		guruAddr: addr,
		client:   sdk.NewClient(GuruTimeout),
	}
}

func (c *GURU) ListCoolingWorkSet(input *ListCoolingWorkSetInput, output *ListCoolingWorkSetOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := make(map[string]string)
	params["unit_id"] = input.UnitId
	params["limit"] = input.Limit
	params["token"] = input.Token
	return c.request(LIST_COOLING_WORKSET, params, output)
}

func (c *GURU) ListHeatingWorkSet(input *ListHeatingWorkSetInput, output *ListHeatingWorkSetOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := make(map[string]string)
	params["node_id"] = input.NodeId
	params["generation_num"] = input.GenerationNum
	params["limit"] = input.Limit
	params["token"] = input.Token
	return c.request(LIST_HEATING_WORKSET, params, output)
}

func (c *GURU) ListGcWorkSet(input *ListGcWorkSetInput, output *ListGcWorkSetOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := make(map[string]string)
	params["owner_id"] = input.OwnerId
	params["owner_type"] = input.OwnerType
	params["limit"] = input.Limit
	params["token"] = input.Token
	return c.request(LIST_GC_WORKSET, params, output)
}

func (c *GURU) ListDataRepairWorkSet(input *ListDataRepairWorkSetInput, output *ListDataRepairWorkSetOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	params := make(map[string]string)
	params["repair_disk_id"] = input.RepairDiskId
	params["replicate_disk_id"] = input.ReplicateDiskId
	params["limit"] = input.Limit
	params["next_token"] = input.Token
	return c.request(LIST_DATA_REPAIR_WORKSET, params, output)
}

func (c *GURU) GetDestoryVolumeDriverPlan(input *GetDestoryVolumeDriverPlanInput, output *GetDestoryVolumeDriverPlanOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.requestPost(GET_DESTORY_VOLUME_DRIVER_PLAN, input, output)
}

func (c *GURU) ChooseDestCvm(input *ChooseDestCvmInput, output *ChooseDestCvmOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.requestPost(CHOOSE_DEST_CVM, input, output)
}

func (c *GURU) request(path string, input map[string]string, output interface{}) (*resty.Response, error) {
	url := c.guruAddr + path
	resp, err := c.client.R().SetHeaders(map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
	}).SetQueryParams(input).SetResult(output).Get(url)

	if err != nil {
		return resp, err
	}
	statusCode := resp.StatusCode()
	if statusCode/100 != 2 {
		return resp, fmt.Errorf("status code is %d, should be 2xx, input is %+v, resp is %+v", statusCode, input, resp)
	}
	return resp, err
}

func (c *GURU) requestPost(path string, input interface{}, output interface{}) (*resty.Response, error) {
	url := c.guruAddr + path
	resp, err := c.client.R().SetHeaders(map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
	}).SetBody(input).SetResult(output).Post(url)

	if err != nil {
		return resp, err
	}
	statusCode := resp.StatusCode()
	if statusCode/100 != 2 {
		return resp, fmt.Errorf("status code is %d, should be 2xx, input is %+v, resp is %+v", statusCode, input, resp)
	}
	return resp, err
}
