package swfclient

import (
	"automation/sdk/sdkerr"
)

type FailWorkflowExecutionDecisionAttributes struct {
	Details *string `json:"details,omitempty" type:"string"`
	Reason  *string `json:"reason,omitempty" type:"string"`
}

func (s *FailWorkflowExecutionDecisionAttributes) SetDetails(v string) *FailWorkflowExecutionDecisionAttributes {
	s.Details = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *FailWorkflowExecutionDecisionAttributes) SetReason(v string) *FailWorkflowExecutionDecisionAttributes {
	s.Reason = &v
	return s
}

type CompleteWorkflowExecutionDecisionAttributes struct {
	Result *string `json:"result,omitempty" type:"string"`
}

func (s *CompleteWorkflowExecutionDecisionAttributes) SetResult(v string) *CompleteWorkflowExecutionDecisionAttributes {
	s.Result = &v
	return s
}

type RequestCancelActivityTaskDecisionAttributes struct {
	ActivityId *string `json:"activityId,omitempty" min:"1" type:"string" required:"true"`
}

// SetActivityId sets the ActivityId field's value.
func (s *RequestCancelActivityTaskDecisionAttributes) SetActivityId(v string) *RequestCancelActivityTaskDecisionAttributes {
	s.ActivityId = &v
	return s
}
func (s *RequestCancelActivityTaskDecisionAttributes) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RequestCancelActivityTaskDecisionAttributes"}
	if s.ActivityId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("ActivityId"))
	}
	if s.ActivityId != nil && len(*s.ActivityId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("ActivityId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ScheduleActivityTaskDecisionAttributes struct {
	ActivityId             *string       `json:"activityId,omitempty" min:"1" type:"string" required:"true"`
	ActivityType           *ActivityType `json:"activityType,omitempty" type:"structure" required:"true"`
	Control                *string       `json:"control,omitempty" type:"string"`
	HeartbeatTimeout       *string       `json:"heartbeatTimeout,omitempty" type:"string"`
	Input                  *string       `json:"input,omitempty" type:"string"`
	ScheduleToCloseTimeout *string       `json:"scheduleToCloseTimeout,omitempty" type:"string"`
	ScheduleToStartTimeout *string       `json:"scheduleToStartTimeout,omitempty" type:"string"`
	StartToCloseTimeout    *string       `json:"startToCloseTimeout,omitempty" type:"string"`
	TaskList               *TaskList     `json:"taskList,omitempty" type:"structure"`
	TaskPriority           *string       `json:"taskPriority,omitempty" type:"string"`
}

// SetActivityId sets the ActivityId field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetActivityId(v string) *ScheduleActivityTaskDecisionAttributes {
	s.ActivityId = &v
	return s
}

// SetActivityType sets the ActivityType field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetActivityType(v *ActivityType) *ScheduleActivityTaskDecisionAttributes {
	s.ActivityType = v
	return s
}

// SetControl sets the Control field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetControl(v string) *ScheduleActivityTaskDecisionAttributes {
	s.Control = &v
	return s
}

// SetHeartbeatTimeout sets the HeartbeatTimeout field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetHeartbeatTimeout(v string) *ScheduleActivityTaskDecisionAttributes {
	s.HeartbeatTimeout = &v
	return s
}

// SetInput sets the Input field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetInput(v string) *ScheduleActivityTaskDecisionAttributes {
	s.Input = &v
	return s
}

// SetScheduleToCloseTimeout sets the ScheduleToCloseTimeout field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetScheduleToCloseTimeout(v string) *ScheduleActivityTaskDecisionAttributes {
	s.ScheduleToCloseTimeout = &v
	return s
}

// SetScheduleToStartTimeout sets the ScheduleToStartTimeout field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetScheduleToStartTimeout(v string) *ScheduleActivityTaskDecisionAttributes {
	s.ScheduleToStartTimeout = &v
	return s
}

// SetStartToCloseTimeout sets the StartToCloseTimeout field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetStartToCloseTimeout(v string) *ScheduleActivityTaskDecisionAttributes {
	s.StartToCloseTimeout = &v
	return s
}

// SetTaskList sets the TaskList field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetTaskList(v *TaskList) *ScheduleActivityTaskDecisionAttributes {
	s.TaskList = v
	return s
}

// SetTaskPriority sets the TaskPriority field's value.
func (s *ScheduleActivityTaskDecisionAttributes) SetTaskPriority(v string) *ScheduleActivityTaskDecisionAttributes {
	s.TaskPriority = &v
	return s
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ScheduleActivityTaskDecisionAttributes) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ScheduleActivityTaskDecisionAttributes"}
	if s.ActivityId == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("ActivityId"))
	}
	if s.ActivityId != nil && len(*s.ActivityId) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("ActivityId", 1))
	}
	if s.ActivityType == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("ActivityType"))
	}
	if s.ActivityType != nil {
		if err := s.ActivityType.Validate(); err != nil {
			invalidParams.AddNested("ActivityType", err.(sdkerr.ErrInvalidParams))
		}
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
