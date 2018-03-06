package swfclient

import (
	"automation/sdk/sdkerr"
)

type WorkflowExecutionConfiguration struct {
	ChildPolicy                  *string   `json:"childPolicy,omitempty" type:"string" required:"true" enum:"ChildPolicy"`
	ExecutionStartToCloseTimeout *string   `json:"executionStartToCloseTimeout,omitempty" min:"1" type:"string" required:"true"`
	LambdaRole                   *string   `json:"lambdaRole,omitempty" min:"1" type:"string"`
	TaskList                     *TaskList `json:"taskList,omitempty" type:"structure" required:"true"`
	TaskPriority                 *string   `json:"taskPriority,omitempty" type:"string"`
	TaskStartToCloseTimeout      *string   `json:"taskStartToCloseTimeout,omitempty" min:"1" type:"string" required:"true"`
}

type WorkflowExecution struct {
	RunId      *string `json:"runId" min:"1" type:"string" required:"true"`
	WorkflowId *string `json:"workflowId" min:"1" type:"string" required:"true"`
}

func (s *WorkflowExecution) SetRunId(v string) *WorkflowExecution {
	s.RunId = &v
	return s
}

func (s *WorkflowExecution) SetWorkflowId(v string) *WorkflowExecution {
	s.WorkflowId = &v
	return s
}

func (s *WorkflowExecution) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "WorkflowExecution"}
	if s.RunId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("RunId"))
	}
	if s.RunId != nil && len(*s.RunId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("RunId", 1))
	}
	if s.WorkflowId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("WorkflowId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type WorkflowExecutionOpenCounts struct {
	OpenActivityTasks           *int64 `json:"openActivityTasks,omitempty" type:"integer" required:"true"`
	OpenChildWorkflowExecutions *int64 `json:"openChildWorkflowExecutions,omitempty" type:"integer" required:"true"`
	OpenDecisionTasks           *int64 `json:"openDecisionTasks,omitempty" type:"integer" required:"true"`
	OpenLambdaFunctions         *int64 `json:"openLambdaFunctions,omitempty" type:"integer"`
	OpenTimers                  *int64 `json:"openTimers,omitempty" type:"integer" required:"true"`
}

/*****************************
*  describe workflow execution
**/

type DescribeWorkflowExecutionInput struct {
	Domain    *string            `json:"domain" min:"1" type:"string" required:"true"`
	Execution *WorkflowExecution `json:"execution" type:"structure" required:"true"`
}

type WorkflowExecutionInfo struct {
	CancelRequested        *bool              `json:"cancelRequested,omitempty" type:"boolean"`
	CloseStatus            *string            `json:"closeStatus,omitempty" type:"string" enum:"CloseStatus"`
	CloseTimestamp         *int64             `json:"closeTimestamp,omitempty" type:"timestamp" timestampFormat:"unix"`
	Execution              *WorkflowExecution `json:"execution,omitempty" type:"structure" required:"true"`
	ExecutionStatus        *string            `json:"executionStatus,omitempty" type:"string" required:"true" enum:"ExecutionStatus"`
	Parent                 *WorkflowExecution `json:"parent,omitempty" type:"structure"`
	StartTimestamp         *int64             `json:"startTimestamp,omitempty" type:"timestamp" timestampFormat:"unix" required:"true"`
	TagList                []*string          `json:"tagList,omitempty" type:"list"`
	WorkflowType           *WorkflowType      `json:"workflowType,omitempty" type:"structure" required:"true"`
	LatestExecutionContext *string            `json:"latestExecutionContext,omitempty" type:"string"`
}

func (s *DescribeWorkflowExecutionInput) SetDomain(v string) *DescribeWorkflowExecutionInput {
	s.Domain = &v
	return s
}

func (s *DescribeWorkflowExecutionInput) SetExecution(v *WorkflowExecution) *DescribeWorkflowExecutionInput {
	s.Execution = v
	return s
}

func (s *DescribeWorkflowExecutionInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "DescribeWorkflowExecutionInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.Execution == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Execution"))
	}
	if s.Execution != nil {
		if err := s.Execution.Validate(); err != nil {
			invalidParams.AddNested("Execution", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkflowExecutionOutput struct {
	ExecutionConfiguration      *WorkflowExecutionConfiguration `json:"executionConfiguration,omitempty" type:"structure" required:"true"`
	ExecutionInfo               *WorkflowExecutionInfo          `json:"executionInfo,omitempty" type:"structure" required:"true"`
	LatestActivityTaskTimestamp *int64                          `json:"latestActivityTaskTimestamp,omitempty" type:"timestamp" timestampFormat:"unix"`
	LatestExecutionContext      *string                         `json:"latestExecutionContext,omitempty" type:"string"`
	OpenCounts                  *WorkflowExecutionOpenCounts    `json:"openCounts,omitempty" type:"structure" required:"true"`
}

/*****************************
*  list open workflow execution
**/

type ExecutionTimeFilter struct {
	LatestDate *int64 `json:"latestDate,omitempty" type:"timestamp" timestampFormat:"unix"`
	OldestDate *int64 `json:"oldestDate,omitempty" type:"timestamp" timestampFormat:"unix" required:"true"`
}

func (s *ExecutionTimeFilter) SetLatestDate(v int64) *ExecutionTimeFilter {
	s.LatestDate = &v
	return s
}

func (s *ExecutionTimeFilter) SetOldestDate(v int64) *ExecutionTimeFilter {
	s.OldestDate = &v
	return s
}

func (s *ExecutionTimeFilter) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ExecutionTimeFilter"}
	if s.OldestDate == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("OldestDate"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type WorkflowTypeFilter struct {
	Name    *string `json:"name,omitempty" min:"1" type:"string" required:"true"`
	Version *string `json:"version,omitempty" type:"string"`
}

func (s *WorkflowTypeFilter) SetName(v string) *WorkflowTypeFilter {
	s.Name = &v
	return s
}

func (s *WorkflowTypeFilter) SetVersion(v string) *WorkflowTypeFilter {
	s.Version = &v
	return s
}

func (s *WorkflowTypeFilter) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "WorkflowTypeFilter"}
	if s.Name == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type WorkflowExecutionFilter struct {
	WorkflowId *string `json:"workflowId,omitempty" min:"1" type:"string" required:"true"`
}

func (s *WorkflowExecutionFilter) SetWorkflowId(v string) *WorkflowExecutionFilter {
	s.WorkflowId = &v
	return s
}

func (s *WorkflowExecutionFilter) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "WorkflowExecutionFilter"}
	if s.WorkflowId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("WorkflowId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagFilter struct {
	Tag *string `json:"tag,omitempty" type:"string" required:"true"`
}

func (s *TagFilter) SetTag(v string) *TagFilter {
	s.Tag = &v
	return s
}

func (s *TagFilter) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "TagFilter"}
	if s.Tag == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Tag"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TaskList struct {
	Name *string `json:"name,omitempty" min:"1" type:"string" required:"true"`
}

func (s *TaskList) SetName(v string) *TaskList {
	s.Name = &v
	return s
}

func (s *TaskList) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "TaskList"}
	if s.Name == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListOpenWorkflowExecutionsInput struct {
	Domain          *string                  `json:"domain,omitempty" min:"1" type:"string" required:"true"`
	ExecutionFilter *WorkflowExecutionFilter `json:"executionFilter,omitempty" type:"structure"`
	MaximumPageSize *int64                   `json:"maximumPageSize,omitempty" type:"integer"`
	NextPageToken   *string                  `json:"nextPageToken,omitempty" type:"string"`
	ReverseOrder    *bool                    `json:"reverseOrder,omitempty" type:"boolean"`
	StartTimeFilter *ExecutionTimeFilter     `json:"startTimeFilter,omitempty" type:"structure" required:"true"`
	TagFilter       *TagFilter               `json:"tagFilter,omitempty" type:"structure"`
	TypeFilter      *WorkflowTypeFilter      `json:"typeFilter,omitempty" type:"structure"`
}

func (s *ListOpenWorkflowExecutionsInput) SetDomain(v string) *ListOpenWorkflowExecutionsInput {
	s.Domain = &v
	return s
}

func (s *ListOpenWorkflowExecutionsInput) SetExecutionFilter(v *WorkflowExecutionFilter) *ListOpenWorkflowExecutionsInput {
	s.ExecutionFilter = v
	return s
}

func (s *ListOpenWorkflowExecutionsInput) SetMaximumPageSize(v int64) *ListOpenWorkflowExecutionsInput {
	s.MaximumPageSize = &v
	return s
}

func (s *ListOpenWorkflowExecutionsInput) SetNextPageToken(v string) *ListOpenWorkflowExecutionsInput {
	s.NextPageToken = &v
	return s
}

func (s *ListOpenWorkflowExecutionsInput) SetReverseOrder(v bool) *ListOpenWorkflowExecutionsInput {
	s.ReverseOrder = &v
	return s
}

func (s *ListOpenWorkflowExecutionsInput) SetStartTimeFilter(v *ExecutionTimeFilter) *ListOpenWorkflowExecutionsInput {
	s.StartTimeFilter = v
	return s
}

func (s *ListOpenWorkflowExecutionsInput) SetTagFilter(v *TagFilter) *ListOpenWorkflowExecutionsInput {
	s.TagFilter = v
	return s
}

// SetTypeFilter sets the TypeFilter field's value.
func (s *ListOpenWorkflowExecutionsInput) SetTypeFilter(v *WorkflowTypeFilter) *ListOpenWorkflowExecutionsInput {
	s.TypeFilter = v
	return s
}
func (s *ListOpenWorkflowExecutionsInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListOpenWorkflowExecutionsInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.StartTimeFilter == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("StartTimeFilter"))
	}
	if s.ExecutionFilter != nil {
		if err := s.ExecutionFilter.Validate(); err != nil {
			invalidParams.AddNested("ExecutionFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.StartTimeFilter != nil {
		if err := s.StartTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("StartTimeFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.TagFilter != nil {
		if err := s.TagFilter.Validate(); err != nil {
			invalidParams.AddNested("TagFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.TypeFilter != nil {
		if err := s.TypeFilter.Validate(); err != nil {
			invalidParams.AddNested("TypeFilter", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCloseWorkflowExecutionsInput struct {
	Domain          *string                  `json:"domain,omitempty" min:"1" type:"string" required:"true"`
	ExecutionFilter *WorkflowExecutionFilter `json:"executionFilter,omitempty" type:"structure"`
	MaximumPageSize *int64                   `json:"maximumPageSize,omitempty" type:"integer"`
	NextPageToken   *string                  `json:"nextPageToken,omitempty" type:"string"`
	ReverseOrder    *bool                    `json:"reverseOrder,omitempty" type:"boolean"`
	StartTimeFilter *ExecutionTimeFilter     `json:"startTimeFilter,omitempty" type:"structure" required:"true"`
	CloseTimeFilter *ExecutionTimeFilter     `json:"closeTimeFilter,omitempty" type:"structure" required:"true"`
	TagFilter       *TagFilter               `json:"tagFilter,omitempty" type:"structure"`
	TypeFilter      *WorkflowTypeFilter      `json:"typeFilter,omitempty" type:"structure"`
}

func (s *ListCloseWorkflowExecutionsInput) SetDomain(v string) *ListCloseWorkflowExecutionsInput {
	s.Domain = &v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetExecutionFilter(v *WorkflowExecutionFilter) *ListCloseWorkflowExecutionsInput {
	s.ExecutionFilter = v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetMaximumPageSize(v int64) *ListCloseWorkflowExecutionsInput {
	s.MaximumPageSize = &v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetNextPageToken(v string) *ListCloseWorkflowExecutionsInput {
	s.NextPageToken = &v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetReverseOrder(v bool) *ListCloseWorkflowExecutionsInput {
	s.ReverseOrder = &v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetStartTimeFilter(v *ExecutionTimeFilter) *ListCloseWorkflowExecutionsInput {
	s.StartTimeFilter = v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetCloseTimeFilter(v *ExecutionTimeFilter) *ListCloseWorkflowExecutionsInput {
	s.CloseTimeFilter = v
	return s
}

func (s *ListCloseWorkflowExecutionsInput) SetTagFilter(v *TagFilter) *ListCloseWorkflowExecutionsInput {
	s.TagFilter = v
	return s
}

// SetTypeFilter sets the TypeFilter field's value.
func (s *ListCloseWorkflowExecutionsInput) SetTypeFilter(v *WorkflowTypeFilter) *ListCloseWorkflowExecutionsInput {
	s.TypeFilter = v
	return s
}
func (s *ListCloseWorkflowExecutionsInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListCloseWorkflowExecutionsInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.StartTimeFilter == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("StartTimeFilter"))
	}
	if s.ExecutionFilter != nil {
		if err := s.ExecutionFilter.Validate(); err != nil {
			invalidParams.AddNested("ExecutionFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.StartTimeFilter != nil {
		if err := s.StartTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("StartTimeFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.CloseTimeFilter != nil {
		if err := s.CloseTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("CloseTimeFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.TagFilter != nil {
		if err := s.TagFilter.Validate(); err != nil {
			invalidParams.AddNested("TagFilter", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.TypeFilter != nil {
		if err := s.TypeFilter.Validate(); err != nil {
			invalidParams.AddNested("TypeFilter", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type WorkflowExecutionInfos struct {
	ExecutionInfos []*WorkflowExecutionInfo `json:"executionInfos,omitempty" type:"list" required:"true"`
	NextPageToken  *string                  `json:"nextPageToken,omitempty" type:"string"`
}

/*****************************
*  start workflow execution
**/

type StartWorkflowExecutionInput struct {
	ChildPolicy                  *string       `json:"childPolicy,omitempty" type:"string" enum:"ChildPolicy"`
	Domain                       *string       `json:"domain,omitempty" min:"1" type:"string" required:"true"`
	ExecutionStartToCloseTimeout *string       `json:"executionStartToCloseTimeout,omitempty" type:"string"`
	Input                        *string       `json:"input,omitempty" type:"string"`
	LambdaRole                   *string       `json:"lambdaRole,omitempty" min:"1" type:"string"`
	TagList                      []*string     `json:"tagList,omitempty" type:"list"`
	TaskList                     *TaskList     `json:"taskList,omitempty" type:"structure"`
	TaskPriority                 *string       `json:"taskPriority,omitempty" type:"string"`
	TaskStartToCloseTimeout      *string       `json:"taskStartToCloseTimeout,omitempty" type:"string"`
	WorkflowId                   *string       `json:"workflowId,omitempty" min:"1" type:"string" required:"true"`
	WorkflowType                 *WorkflowType `json:"workflowType,omitempty" type:"structure" required:"true"`
}

func (s *StartWorkflowExecutionInput) SetChildPolicy(v string) *StartWorkflowExecutionInput {
	s.ChildPolicy = &v
	return s
}

func (s *StartWorkflowExecutionInput) SetDomain(v string) *StartWorkflowExecutionInput {
	s.Domain = &v
	return s
}

func (s *StartWorkflowExecutionInput) SetExecutionStartToCloseTimeout(v string) *StartWorkflowExecutionInput {
	s.ExecutionStartToCloseTimeout = &v
	return s
}

func (s *StartWorkflowExecutionInput) SetInput(v string) *StartWorkflowExecutionInput {
	s.Input = &v
	return s
}

// SetLambdaRole sets the LambdaRole field's value.
func (s *StartWorkflowExecutionInput) SetLambdaRole(v string) *StartWorkflowExecutionInput {
	s.LambdaRole = &v
	return s
}

// SetTagList sets the TagList field's value.
func (s *StartWorkflowExecutionInput) SetTagList(v []*string) *StartWorkflowExecutionInput {
	s.TagList = v
	return s
}

// SetTaskList sets the TaskList field's value.
func (s *StartWorkflowExecutionInput) SetTaskList(v *TaskList) *StartWorkflowExecutionInput {
	s.TaskList = v
	return s
}

// SetTaskPriority sets the TaskPriority field's value.
func (s *StartWorkflowExecutionInput) SetTaskPriority(v string) *StartWorkflowExecutionInput {
	s.TaskPriority = &v
	return s
}

// SetTaskStartToCloseTimeout sets the TaskStartToCloseTimeout field's value.
func (s *StartWorkflowExecutionInput) SetTaskStartToCloseTimeout(v string) *StartWorkflowExecutionInput {
	s.TaskStartToCloseTimeout = &v
	return s
}

// SetWorkflowId sets the WorkflowId field's value.
func (s *StartWorkflowExecutionInput) SetWorkflowId(v string) *StartWorkflowExecutionInput {
	s.WorkflowId = &v
	return s
}

// SetWorkflowType sets the WorkflowType field's value.
func (s *StartWorkflowExecutionInput) SetWorkflowType(v *WorkflowType) *StartWorkflowExecutionInput {
	s.WorkflowType = v
	return s
}
func (s *StartWorkflowExecutionInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "StartWorkflowExecutionInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.LambdaRole != nil && len(*s.LambdaRole) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("LambdaRole", 1))
	}
	if s.WorkflowId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("WorkflowId", 1))
	}
	if s.WorkflowType == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowType"))
	}
	if s.TaskList != nil {
		if err := s.TaskList.Validate(); err != nil {
			invalidParams.AddNested("TaskList", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.WorkflowType != nil {
		if err := s.WorkflowType.Validate(); err != nil {
			invalidParams.AddNested("WorkflowType", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartWorkflowExecutionOutput struct {
	RunId *string `json:"runId" min:"1" type:"string"`
}

/*****************************
*  terminate workflow execution
**/

type TerminateWorkflowExecutionInput struct {
	ChildPolicy *string `json:"childPolicy,omitempty" type:"string" enum:"ChildPolicy"`
	Details     *string `json:"details,omitempty" type:"string"`
	Domain      *string `json:"domain,omitempty" min:"1" type:"string" required:"true"`
	Reason      *string `json:"reason,omitempty" type:"string"`
	RunId       *string `json:"runId,omitempty" type:"string"`
	WorkflowId  *string `json:"workflowId,omitempty" min:"1" type:"string" required:"true"`
}

func (s *TerminateWorkflowExecutionInput) SetChildPolicy(v string) *TerminateWorkflowExecutionInput {
	s.ChildPolicy = &v
	return s
}

// SetDetails sets the Details field's value.
func (s *TerminateWorkflowExecutionInput) SetDetails(v string) *TerminateWorkflowExecutionInput {
	s.Details = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *TerminateWorkflowExecutionInput) SetDomain(v string) *TerminateWorkflowExecutionInput {
	s.Domain = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *TerminateWorkflowExecutionInput) SetReason(v string) *TerminateWorkflowExecutionInput {
	s.Reason = &v
	return s
}

// SetRunId sets the RunId field's value.
func (s *TerminateWorkflowExecutionInput) SetRunId(v string) *TerminateWorkflowExecutionInput {
	s.RunId = &v
	return s
}

// SetWorkflowId sets the WorkflowId field's value.
func (s *TerminateWorkflowExecutionInput) SetWorkflowId(v string) *TerminateWorkflowExecutionInput {
	s.WorkflowId = &v
	return s
}

func (s *TerminateWorkflowExecutionInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "TerminateWorkflowExecutionInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.WorkflowId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("WorkflowId", 1))
	}
	return nil
}

type TerminateWorkflowExecutionOutput struct {
	_ struct{} `type:"structure"`
}
