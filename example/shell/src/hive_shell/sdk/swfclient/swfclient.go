package swfclient

import (
	"fmt"
	"time"

	"github.com/go-resty/resty"

	"automation/sdk"
)

const (
	POLLTIMEOUT = 70
)

type SWF struct {
	swfAddr string
	client  *resty.Client
}

func NewSWF(swfAddr string) *SWF {
	return &SWF{
		swfAddr: swfAddr,
		client:  sdk.NewClient(POLLTIMEOUT * time.Second),
	}
}

/*******************domain************************/
func (c *SWF) ListDomains(input *ListDomainsInput, output *ListDomainsOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(LIST_DOMAINS, input, output)
}

func (c *SWF) DescribeDomain(input *DescribeDomainInput, output *DescribeDomainOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(DESCRIBE_DOMAIN, input, output)
}

func (c *SWF) RegisterDomain(input *RegisterDomainInput, output *RegisterDomainOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	resp, err := c.request(REGISTER_DOMAIN, input, output)
	if resp.StatusCode() == 400 {
		err = nil
	}
	return resp, err
}

/*******************workflow_type************************/
func (c *SWF) ListWorkflowTypes(input *ListWorkflowTypesInput, output *ListWorkflowTypesOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(LIST_WORKFLOW_TYPES, input, output)
}

func (c *SWF) DescribeWorkflowType(input *DescribeWorkflowTypeInput, output *DescribeWorkflowTypeOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(DESCRIBE_WORKFLOW_TYPE, input, output)
}

func (c *SWF) RegisterWorkflowType(input *RegisterWorkflowTypeInput, output *RegisterWorkflowTypeOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	resp, err := c.request(REGISTER_WORKFLOW_TYPE, input, output)
	if resp.StatusCode() == 400 {
		err = nil
	}
	return resp, err
}

/*******************actvity_type************************/
func (c *SWF) ListActivityTypes(input *ListActivityTypesInput, output *ListActivityTypesOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(LIST_ACTIVITY_TYPES, input, output)
}

func (c *SWF) RegisterActivityType(input *RegisterActivityTypeInput, output *RegisterActivityTypeOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	resp, err := c.request(REGISTER_ACTIVITY_TYPE, input, output)
	if resp.StatusCode() == 400 {
		err = nil
	}
	return resp, err
}

func (c *SWF) DescribeActivityType(input *DescribeActivityTypeInput, output *DescribeActivityTypeOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(DESCRIBE_ACTIVITY_TYPE, input, output)
}

/*******************workflow execution************************/
func (c *SWF) ListOpenWorkflowExecutions(input *ListOpenWorkflowExecutionsInput, output *WorkflowExecutionInfos) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(LIST_OPEN_WORKFLOW_EXECUTION, input, output)
}

func (c *SWF) ListCloseWorkflowExecutions(input *ListCloseWorkflowExecutionsInput, output *WorkflowExecutionInfos) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(LIST_CLOSE_WORKFLOW_EXECUTION, input, output)
}

func (c *SWF) DescribeWorkflowExecution(input *DescribeWorkflowExecutionInput, output *DescribeWorkflowExecutionOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(DESCRIBE_WORKFLOW_EXECUTION, input, output)
}

func (c *SWF) StartWorkflowExecution(input *StartWorkflowExecutionInput, output *StartWorkflowExecutionOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(START_WORKFLOW_EXECUTION, input, output)
}

func (c *SWF) TerminateWorkflowExecution(input *TerminateWorkflowExecutionInput, output *TerminateWorkflowExecutionOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(TERMINATE_WORKFLOW_EXECUTION, input, output)
}

func (c *SWF) PollForActivityTask(input *PollForActivityTaskInput, output *PollForActivityTaskOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(POLL_FOR_ACTIVITY_TASK, input, output)
}

func (c *SWF) RespondActivityTaskCanceled(input *RespondActivityTaskCanceledInput, output *RespondActivityTaskCanceledOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(RESPOND_ACTIVITY_TASK_CANCELED, input, output)
}

func (c *SWF) RespondActivityTaskFailed(input *RespondActivityTaskFailedInput, output *RespondActivityTaskFailedOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(RESPOND_ACTIVITY_TASK_FAILED, input, output)
}

func (c *SWF) RespondActivityTaskCompleted(input *RespondActivityTaskCompletedInput, output *RespondActivityTaskCompletedOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(RESPOND_ACTIVITY_TASK_COMPLETED, input, output)
}

func (c *SWF) PollForDecisionTask(input *PollForDecisionTaskInput, output *PollForDecisionTaskOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(POLL_FOR_DECISION_TASK, input, output)
}

func (c *SWF) RespondDecisionTaskCompleted(input *RespondDecisionTaskCompletedInput, output *RespondDecisionTaskCompletedOutput) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	return c.request(RESPOND_DECISION_TASK_COMPLETED, input, output)
}

func (c *SWF) request(path string, input interface{}, output interface{}) (*resty.Response, error) {
	url := c.swfAddr + path
	resp, err := c.client.R().SetHeaders(map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
	}).SetBody(input).SetResult(output).Post(url)

	if err != nil {
		return resp, err
	}
	statusCode := resp.StatusCode()
	if statusCode/100 != 2 {
		return resp, fmt.Errorf("status code is %d, should be 2xx", statusCode)
	}
	return resp, err
}
