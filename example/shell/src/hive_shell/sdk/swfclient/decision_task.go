package swfclient

import (
	"automation/sdk/sdkerr"
	"fmt"
)

type Decision struct {
	CompleteWorkflowExecutionDecisionAttributes *CompleteWorkflowExecutionDecisionAttributes `json:"completeWorkflowExecutionDecisionAttributes,omitempty" type:"structure"`
	DecisionType                                *string                                      `json:"decisionType,omitempty" type:"string" required:"true" enum:"DecisionType"`
	FailWorkflowExecutionDecisionAttributes     *FailWorkflowExecutionDecisionAttributes     `json:"failWorkflowExecutionDecisionAttributes,omitempty" type:"structure"`

	RequestCancelActivityTaskDecisionAttributes *RequestCancelActivityTaskDecisionAttributes `json:"ancelActivityTaskDecisionAttributes,omitempty" type:"structure"`

	ScheduleActivityTaskDecisionAttributes *ScheduleActivityTaskDecisionAttributes `json:"scheduleActivityTaskDecisionAttributes,omitempty" type:"structure"`
}

func (s *Decision) SetDecisionType(v string) *Decision {
	s.DecisionType = &v
	return s
}
func (s *Decision) SetScheduleActivityTaskDecisionAttributes(v *ScheduleActivityTaskDecisionAttributes) *Decision {
	s.ScheduleActivityTaskDecisionAttributes = v
	return s
}
func (s *Decision) SetRequestCancelActivityTaskDecisionAttributes(v *RequestCancelActivityTaskDecisionAttributes) *Decision {
	s.RequestCancelActivityTaskDecisionAttributes = v
	return s
}
func (s *Decision) SetCompleteWorkflowExecutionDecisionAttributes(v *CompleteWorkflowExecutionDecisionAttributes) *Decision {
	s.CompleteWorkflowExecutionDecisionAttributes = v
	return s
}

func (s *Decision) SetFailWorkflowExecutionDecisionAttributes(v *FailWorkflowExecutionDecisionAttributes) *Decision {
	s.FailWorkflowExecutionDecisionAttributes = v
	return s
}

func (s *Decision) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "Decision"}
	if s.DecisionType == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("DecisionType"))
	}
	if s.RequestCancelActivityTaskDecisionAttributes != nil {
		if err := s.RequestCancelActivityTaskDecisionAttributes.Validate(); err != nil {
			invalidParams.AddNested("RequestCancelActivityTaskDecisionAttributes", err.(sdkerr.ErrInvalidParams))
		}
	}
	if s.ScheduleActivityTaskDecisionAttributes != nil {
		if err := s.ScheduleActivityTaskDecisionAttributes.Validate(); err != nil {
			invalidParams.AddNested("ScheduleActivityTaskDecisionAttributes", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PollForDecisionTaskInput struct {
	Domain          *string   `json:"domain" min:"1" type:"string" required:"true"`
	Identity        *string   `json:"identity,omitempty" type:"string"`
	MaximumPageSize *int64    `json:"maximumPageSize,omitempty" type:"integer"`
	NextPageToken   *string   `json:"nextPageToken,omitempty" type:"string"`
	ReverseOrder    *bool     `json:"reverseOrder,omitempty" type:"boolean"`
	TaskList        *TaskList `json:"taskList" type:"structure" required:"true"`
}

// SetDomain sets the Domain field's value.
func (s *PollForDecisionTaskInput) SetDomain(v string) *PollForDecisionTaskInput {
	s.Domain = &v
	return s
}

// SetIdentity sets the Identity field's value.
func (s *PollForDecisionTaskInput) SetIdentity(v string) *PollForDecisionTaskInput {
	s.Identity = &v
	return s
}

// SetMaximumPageSize sets the MaximumPageSize field's value.
func (s *PollForDecisionTaskInput) SetMaximumPageSize(v int64) *PollForDecisionTaskInput {
	s.MaximumPageSize = &v
	return s
}

// SetNextPageToken sets the NextPageToken field's value.
func (s *PollForDecisionTaskInput) SetNextPageToken(v string) *PollForDecisionTaskInput {
	s.NextPageToken = &v
	return s
}

// SetReverseOrder sets the ReverseOrder field's value.
func (s *PollForDecisionTaskInput) SetReverseOrder(v bool) *PollForDecisionTaskInput {
	s.ReverseOrder = &v
	return s
}

func (s *PollForDecisionTaskInput) SetTaskList(v *TaskList) *PollForDecisionTaskInput {
	s.TaskList = v
	return s
}

func (s *PollForDecisionTaskInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "PollForDecisionTaskInput"}
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

type PollForDecisionTaskOutput struct {
	Events                 []*HistoryEvent    `json:"events,omitempty" type:"list" required:"true"`
	NextPageToken          *string            `json:"nextPageToken,omitempty" type:"string"`
	PreviousStartedEventId *int64             `json:"previousStartedEventId,omitempty" type:"long"`
	StartedEventId         *int64             `json:"startedEventId,omitempty" type:"long" required:"true"`
	TaskToken              *string            `json:"taskToken,omitempty" min:"1" type:"string" required:"true"`
	WorkflowExecution      *WorkflowExecution `json:"workflowExecution,omitempty" type:"structure" required:"true"`

	WorkflowType *WorkflowType `json:"workflowType" type:"structure" required:"true"`
}

type RespondDecisionTaskCompletedInput struct {
	Decisions        []*Decision `json:"decisions" type:"list"`
	ExecutionContext *string     `json:"executionContext,omitempty" type:"string"`
	TaskToken        *string     `json:"taskToken,omitempty" min:"1" type:"string" required:"true"`
}

// SetDecisions sets the Decisions field's value.
func (s *RespondDecisionTaskCompletedInput) SetDecisions(v []*Decision) *RespondDecisionTaskCompletedInput {
	s.Decisions = v
	return s
}

// SetExecutionContext sets the ExecutionContext field's value.
func (s *RespondDecisionTaskCompletedInput) SetExecutionContext(v string) *RespondDecisionTaskCompletedInput {
	s.ExecutionContext = &v
	return s
}

// SetTaskToken sets the TaskToken field's value.
func (s *RespondDecisionTaskCompletedInput) SetTaskToken(v string) *RespondDecisionTaskCompletedInput {
	s.TaskToken = &v
	return s
}

func (s *RespondDecisionTaskCompletedInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RespondDecisionTaskCompletedInput"}
	if s.TaskToken == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("TaskToken", 1))
	}
	if s.Decisions != nil {
		for i, v := range s.Decisions {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Decisions", i), err.(sdkerr.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RespondDecisionTaskCompletedOutput struct {
	_ struct{} `type:"structure"`
}

const (
	// DecisionTypeScheduleActivityTask is a DecisionType enum value
	DecisionTypeScheduleActivityTask = "ScheduleActivityTask"

	// DecisionTypeRequestCancelActivityTask is a DecisionType enum value
	DecisionTypeRequestCancelActivityTask = "RequestCancelActivityTask"

	// DecisionTypeCompleteWorkflowExecution is a DecisionType enum value
	DecisionTypeCompleteWorkflowExecution = "CompleteWorkflowExecution"

	// DecisionTypeFailWorkflowExecution is a DecisionType enum value
	DecisionTypeFailWorkflowExecution = "FailWorkflowExecution"

	// DecisionTypeCancelWorkflowExecution is a DecisionType enum value
	DecisionTypeCancelWorkflowExecution = "CancelWorkflowExecution"

	// DecisionTypeContinueAsNewWorkflowExecution is a DecisionType enum value
	DecisionTypeContinueAsNewWorkflowExecution = "ContinueAsNewWorkflowExecution"

	// DecisionTypeRecordMarker is a DecisionType enum value
	DecisionTypeRecordMarker = "RecordMarker"

	// DecisionTypeStartTimer is a DecisionType enum value
	DecisionTypeStartTimer = "StartTimer"

	// DecisionTypeCancelTimer is a DecisionType enum value
	DecisionTypeCancelTimer = "CancelTimer"

	// DecisionTypeSignalExternalWorkflowExecution is a DecisionType enum value
	DecisionTypeSignalExternalWorkflowExecution = "SignalExternalWorkflowExecution"

	// DecisionTypeRequestCancelExternalWorkflowExecution is a DecisionType enum value
	DecisionTypeRequestCancelExternalWorkflowExecution = "RequestCancelExternalWorkflowExecution"

	// DecisionTypeStartChildWorkflowExecution is a DecisionType enum value
	DecisionTypeStartChildWorkflowExecution = "StartChildWorkflowExecution"

	// DecisionTypeScheduleLambdaFunction is a DecisionType enum value
	DecisionTypeScheduleLambdaFunction = "ScheduleLambdaFunction"
)
