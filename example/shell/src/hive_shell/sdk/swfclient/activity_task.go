package swfclient

import (
	"automation/sdk/sdkerr"
)

type PollForActivityTaskInput struct {
	Domain   *string   `json:"domain" min:"1" type:"string" required:"true"`
	Identity *string   `json:"identity" type:"string"`
	TaskList *TaskList `json:"taskList" type:"structure" required:"true"`
}

// SetDomain sets the Domain field's value.
func (s *PollForActivityTaskInput) SetDomain(v string) *PollForActivityTaskInput {
	s.Domain = &v
	return s
}

// SetIdentity sets the Identity field's value.
func (s *PollForActivityTaskInput) SetIdentity(v string) *PollForActivityTaskInput {
	s.Identity = &v
	return s
}

// SetTaskList sets the TaskList field's value.
func (s *PollForActivityTaskInput) SetTaskList(v *TaskList) *PollForActivityTaskInput {
	s.TaskList = v
	return s
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PollForActivityTaskInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "PollForActivityTaskInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.TaskList == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("TaskList"))
	}
	if s.TaskList != nil {
		if err := s.TaskList.Validate(); err != nil {
			invalidParams.AddNested("TaskList", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PollForActivityTaskOutput struct {
	ActivityId        *string            `json:"activityId" min:"1" type:"string" required:"true"`
	ActivityType      *ActivityType      `json:"activityType" type:"structure" required:"true"`
	Input             *string            `json:"input" type:"string"`
	StartedEventId    *int64             `json:"startedEventId" type:"long" required:"true"`
	TaskToken         *string            `json:"taskToken" min:"1" type:"string" required:"true"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution" type:"structure" required:"true"`
}

type RespondActivityTaskCanceledInput struct {
	Details   *string `json:"details" type:"string"`
	TaskToken *string `json:"taskToken" min:"1" type:"string" required:"true"`
}

// SetDetails sets the Details field's value.
func (s *RespondActivityTaskCanceledInput) SetDetails(v string) *RespondActivityTaskCanceledInput {
	s.Details = &v
	return s
}

// SetTaskToken sets the TaskToken field's value.
func (s *RespondActivityTaskCanceledInput) SetTaskToken(v string) *RespondActivityTaskCanceledInput {
	s.TaskToken = &v
	return s
}

func (s *RespondActivityTaskCanceledInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RespondActivityTaskCanceledInput"}
	if s.TaskToken == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RespondActivityTaskCanceledOutput struct {
}

type RespondActivityTaskCompletedInput struct {
	Result    *string `json:"result" type:"string"`
	TaskToken *string `json:"taskToken" min:"1" type:"string" required:"true"`
}

// SetResult sets the Result field's value.
func (s *RespondActivityTaskCompletedInput) SetResult(v string) *RespondActivityTaskCompletedInput {
	s.Result = &v
	return s
}

// SetTaskToken sets the TaskToken field's value.
func (s *RespondActivityTaskCompletedInput) SetTaskToken(v string) *RespondActivityTaskCompletedInput {
	s.TaskToken = &v
	return s
}
func (s *RespondActivityTaskCompletedInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RespondActivityTaskCompletedInput"}
	if s.TaskToken == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RespondActivityTaskCompletedOutput struct {
}

type RespondActivityTaskFailedInput struct {
	Details   *string `json:"details" type:"string"`
	Reason    *string `json:"reason" type:"string"`
	TaskToken *string `json:"taskToken" min:"1" type:"string" required:"true"`
}

// SetDetails sets the Details field's value.
func (s *RespondActivityTaskFailedInput) SetDetails(v string) *RespondActivityTaskFailedInput {
	s.Details = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *RespondActivityTaskFailedInput) SetReason(v string) *RespondActivityTaskFailedInput {
	s.Reason = &v
	return s
}

// SetTaskToken sets the TaskToken field's value.
func (s *RespondActivityTaskFailedInput) SetTaskToken(v string) *RespondActivityTaskFailedInput {
	s.TaskToken = &v
	return s
}
func (s *RespondActivityTaskFailedInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RespondActivityTaskFailedInput"}
	if s.TaskToken == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RespondActivityTaskFailedOutput struct {
}
