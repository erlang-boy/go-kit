package swfclient

import (
	"time"
)

var PollTimeout time.Duration = 70 * time.Second //must be lager than 60 by swf, if swf response token empty, it take 60s

const (
	REGISTER_DOMAIN = "/v1/action/RegisterDomain"
	LIST_DOMAINS    = "/v1/action/ListDomains"
	DESCRIBE_DOMAIN = "/v1/action/DescribeDomain"

	REGISTER_WORKFLOW_TYPE = "/v1/action/RegisterWorkflowType"
	LIST_WORKFLOW_TYPES    = "/v1/action/ListWorkflowTypes"
	DESCRIBE_WORKFLOW_TYPE = "/v1/action/DescribeWorkflowType"

	REGISTER_ACTIVITY_TYPE = "/v1/action/RegisterActivityType"
	LIST_ACTIVITY_TYPES    = "/v1/action/ListActivityTypes"
	DESCRIBE_ACTIVITY_TYPE = "/v1/action/DescribeActivityType"

	LIST_OPEN_WORKFLOW_EXECUTION  = "/v1/action/ListOpenWorkflowExecutions"
	LIST_CLOSE_WORKFLOW_EXECUTION = "/v1/action/ListCloseWorkflowExecutions"
	DESCRIBE_WORKFLOW_EXECUTION   = "/v1/action/DescribeWorkflowExecution"
	START_WORKFLOW_EXECUTION      = "/v1/action/StartWorkflowExecution"
	TERMINATE_WORKFLOW_EXECUTION  = "/v1/action/TerminateWorkflowExecution"

	POLL_FOR_DECISION_TASK          = "/v1/action/PollForDecisionTask"
	RESPOND_DECISION_TASK_COMPLETED = "/v1/action/RespondDecisionTaskCompleted"

	POLL_FOR_ACTIVITY_TASK          = "/v1/action/PollForActivityTask"
	RESPOND_ACTIVITY_TASK_COMPLETED = "/v1/action/RespondActivityTaskCompleted"
	RESPOND_ACTIVITY_TASK_FAILED    = "/v1/action/RespondActivityTaskFailed"
	RESPOND_ACTIVITY_TASK_CANCELED  = "/v1/action/RespondActivityTaskCanceled"
)
