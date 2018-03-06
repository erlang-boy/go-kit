package swfclient

type HistoryEvent struct {

	// If the event is of type ActivityTaskcancelRequested then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	ActivityTaskCancelRequestedEventAttributes *ActivityTaskCancelRequestedEventAttributes `json:"activityTaskCancelRequestedEventAttributes,omitempty" type:"structure"`

	// If the event is of type ActivityTaskCanceled then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	ActivityTaskCanceledEventAttributes *ActivityTaskCanceledEventAttributes `json:"activityTaskCanceledEventAttributes,omitempty" type:"structure"`

	// If the event is of type ActivityTaskCompleted then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	ActivityTaskCompletedEventAttributes *ActivityTaskCompletedEventAttributes `json:"activityTaskCompletedEventAttributes,omitempty" type:"structure"`

	// If the event is of type ActivityTaskFailed then this member is set and provides
	// detailed information about the event. It isn't set for other event types.
	ActivityTaskFailedEventAttributes *ActivityTaskFailedEventAttributes `json:"activityTaskFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type ActivityTaskScheduled then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	ActivityTaskScheduledEventAttributes *ActivityTaskScheduledEventAttributes `json:"activityTaskScheduledEventAttributes,omitempty" type:"structure"`

	// If the event is of type ActivityTaskStarted then this member is set and provides
	// detailed information about the event. It isn't set for other event types.
	ActivityTaskStartedEventAttributes *ActivityTaskStartedEventAttributes `json:"activityTaskStartedEventAttributes,omitempty" type:"structure"`

	// If the event is of type ActivityTaskTimedOut then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	ActivityTaskTimedOutEventAttributes *ActivityTaskTimedOutEventAttributes `json:"activityTaskTimedOutEventAttributes,omitempty" type:"structure"`

	// If the event is of type CancelWorkflowExecutionFailed then this member is
	// set and provides detailed information about the event. It isn't set for other
	// event types.
	CancelWorkflowExecutionFailedEventAttributes *CancelWorkflowExecutionFailedEventAttributes `json:"cancelWorkflowExecutionFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type CompleteWorkflowExecutionFailed then this member
	// is set and provides detailed information about the event. It isn't set for
	// other event types.
	CompleteWorkflowExecutionFailedEventAttributes *CompleteWorkflowExecutionFailedEventAttributes `json:"completeWorkflowExecutionFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type DecisionTaskCompleted then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	DecisionTaskCompletedEventAttributes *DecisionTaskCompletedEventAttributes `json:"decisionTaskCompletedEventAttributes,omitempty" type:"structure"`

	// If the event is of type DecisionTaskScheduled then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	DecisionTaskScheduledEventAttributes *DecisionTaskScheduledEventAttributes `json:"decisionTaskScheduledEventAttributes,omitempty" type:"structure"`

	// If the event is of type DecisionTaskStarted then this member is set and provides
	// detailed information about the event. It isn't set for other event types.
	DecisionTaskStartedEventAttributes *DecisionTaskStartedEventAttributes `json:"decisionTaskStartedEventAttributes,omitempty" type:"structure"`

	// If the event is of type DecisionTaskTimedOut then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	DecisionTaskTimedOutEventAttributes *DecisionTaskTimedOutEventAttributes `json:"decisionTaskTimedOutEventAttributes,omitempty" type:"structure"`

	// The system generated ID of the event. This ID uniquely identifies the event
	// with in the workflow execution history.
	//
	// EventId is a required field
	EventId *int64 `json:"eventId,omitempty" type:"long" required:"true"`

	// The date and time when the event occurred.
	//
	// EventTimestamp is a required field
	EventTimestamp *int64 `json:"eventTimestamp,omitempty" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The type of the history event.
	//
	// EventType is a required field
	EventType *string `json:"eventType,omitempty" type:"string" required:"true" enum:"EventType"`

	// If the event is of type FailWorkflowExecutionFailed then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	FailWorkflowExecutionFailedEventAttributes *FailWorkflowExecutionFailedEventAttributes `json:"failWorkflowExecutionFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type RequestCancelActivityTaskFailed then this member
	// is set and provides detailed information about the event. It isn't set for
	// other event types.
	RequestCancelActivityTaskFailedEventAttributes *RequestCancelActivityTaskFailedEventAttributes `json:"requestCancelActivityTaskFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type ScheduleActivityTaskFailed then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	ScheduleActivityTaskFailedEventAttributes *ScheduleActivityTaskFailedEventAttributes `json:"scheduleActivityTaskFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionCancelRequested then this member
	// is set and provides detailed information about the event. It isn't set for
	// other event types.
	WorkflowExecutionCancelRequestedEventAttributes *WorkflowExecutionCancelRequestedEventAttributes `json:"workflowExecutionCancelRequestedEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionCanceled then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	WorkflowExecutionCanceledEventAttributes *WorkflowExecutionCanceledEventAttributes `json:"workflowExecutionCanceledEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionCompleted then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	WorkflowExecutionCompletedEventAttributes *WorkflowExecutionCompletedEventAttributes `json:"workflowExecutionCompletedEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionFailed then this member is set and
	// provides detailed information about the event. It isn't set for other event
	// types.
	WorkflowExecutionFailedEventAttributes *WorkflowExecutionFailedEventAttributes `json:"workflowExecutionFailedEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionStarted then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	WorkflowExecutionStartedEventAttributes *WorkflowExecutionStartedEventAttributes `json:"workflowExecutionStartedEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionTerminated then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	WorkflowExecutionTerminatedEventAttributes *WorkflowExecutionTerminatedEventAttributes `json:"workflowExecutionTerminatedEventAttributes,omitempty" type:"structure"`

	// If the event is of type WorkflowExecutionTimedOut then this member is set
	// and provides detailed information about the event. It isn't set for other
	// event types.
	WorkflowExecutionTimedOutEventAttributes *WorkflowExecutionTimedOutEventAttributes `json:"workflowExecutionTimedOutEventAttributes,omitempty" type:"structure"`
}

type ActivityTaskCancelRequestedEventAttributes struct {
	// The unique ID of the task.
	//
	// ActivityId is a required field
	ActivityId *string `json:"activityId,omitempty" min:"1" type:"string" required:"true"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the RequestCancelActivityTask decision for this cancellation
	// request. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`
}

// Provides the details of the ActivityTaskCanceled event.
type ActivityTaskCanceledEventAttributes struct {

	// Details of the cancellation.
	Details *string `json:"details,omitempty" type:"string"`

	// If set, contains the ID of the last ActivityTaskCancelRequested event recorded
	// for this activity task. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	LatestCancelRequestedEventId *int64 `json:"latestCancelRequestedEventId,omitempty" type:"long"`

	// The ID of the ActivityTaskScheduled event that was recorded when this activity
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`

	// The ID of the ActivityTaskStarted event recorded when this activity task
	// was started. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"startedEventId,omitempty" type:"long" required:"true"`
}

type ActivityTaskCompletedEventAttributes struct {

	// The results of the activity task.
	Result *string `json:"result,omitempty" type:"string"`

	// The ID of the ActivityTaskScheduled event that was recorded when this activity
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`

	// The ID of the ActivityTaskStarted event recorded when this activity task
	// was started. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"startedEventId,omitempty" type:"long" required:"true"`
}

type ActivityTaskFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The details of the failure.
	Details *string `json:"details,omitempty" type:"string"`

	// The reason provided for the failure.
	Reason *string `json:"reason,omitempty" type:"string"`

	// The ID of the ActivityTaskScheduled event that was recorded when this activity
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`

	// The ID of the ActivityTaskStarted event recorded when this activity task
	// was started. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"startedEventId,omitempty" type:"long" required:"true"`
}

type ActivityTaskScheduledEventAttributes struct {
	_ struct{} `type:"structure"`

	// The unique ID of the activity task.
	//
	// ActivityId is a required field
	ActivityId *string `json:"activityId,omitempty" min:"1" type:"string" required:"true"`

	// The type of the activity task.
	//
	// ActivityType is a required field
	ActivityType *ActivityType `json:"activityType,omitempty" type:"structure" required:"true"`

	// Data attached to the event that can be used by the decider in subsequent
	// workflow tasks. This data isn't sent to the activity.
	Control *string `json:"control,omitempty" type:"string"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision that
	// resulted in the scheduling of this activity task. This information can be
	// useful for diagnosing problems by tracing back the chain of events leading
	// up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`

	HeartbeatTimeout *string `json:"heartbeatTimeout,omitempty" type:"string"`

	// The input provided to the activity task.
	Input *string `json:"input,omitempty" type:"string"`

	// The maximum amount of time for this activity task.
	ScheduleToCloseTimeout *string `json:"scheduleToCloseTimeout,omitempty" type:"string"`

	ScheduleToStartTimeout *string `json:"scheduleToStartTimeout,omitempty" type:"string"`

	StartToCloseTimeout *string `json:"startToCloseTimeout,omitempty" type:"string"`

	// The task list in which the activity task has been scheduled.
	//
	// TaskList is a required field
	TaskList *TaskList `json:"taskList,omitempty" type:"structure" required:"true"`

	// The priority to assign to the scheduled activity task. If set, this overrides
	// any default priority value that was assigned when the activity type was registered.
	//
	// Valid values are integers that range from Java's Integer.MIN_VALUE (-2147483648)
	// to Integer.MAX_VALUE (2147483647). Higher numbers indicate higher priority.
	//
	// For more information about setting task priority, see Setting Task Priority
	// (http://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	TaskPriority *string `json:"taskPriority,omitempty" type:"string"`
}

// Provides the details of the ActivityTaskStarted event.
type ActivityTaskStartedEventAttributes struct {
	_ struct{} `type:"structure"`

	Identity *string `json:"identity,omitempty" type:"string"`

	// The ID of the ActivityTaskScheduled event that was recorded when this activity
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`
}

type ActivityTaskTimedOutEventAttributes struct {
	_ struct{} `type:"structure"`

	// Contains the content of the details parameter for the last call made by the
	// activity to RecordActivityTaskHeartbeat.
	Details *string `json:"details,omitempty" type:"string"`

	// The ID of the ActivityTaskScheduled event that was recorded when this activity
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`

	// The ID of the ActivityTaskStarted event recorded when this activity task
	// was started. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"startedEventId,omitempty" type:"long" required:"true"`

	// The type of the timeout that caused this event.
	//
	// TimeoutType is a required field
	TimeoutType *string `json:"timeoutType,omitempty" type:"string" required:"true" enum:"ActivityTaskTimeoutType"`
}

type CancelWorkflowExecutionFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The cause of the failure. This information is generated by the system and
	// can be useful for diagnostic purposes.
	//
	// If cause is set to OPERATION_NOT_PERMITTED, the decision failed because it
	// lacked sufficient permissions. For details and example IAM policies, see
	// Using IAM to Manage Access to Amazon SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
	// in the Amazon SWF Developer Guide.
	//
	// Cause is a required field
	Cause *string `json:"cause,omitempty" type:"string" required:"true" enum:"CancelWorkflowExecutionFailedCause"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the CancelWorkflowExecution decision for this cancellation
	// request. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`
}

type CompleteWorkflowExecutionFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The cause of the failure. This information is generated by the system and
	// can be useful for diagnostic purposes.
	//
	// If cause is set to OPERATION_NOT_PERMITTED, the decision failed because it
	// lacked sufficient permissions. For details and example IAM policies, see
	// Using IAM to Manage Access to Amazon SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
	// in the Amazon SWF Developer Guide.
	//
	// Cause is a required field
	Cause *string `json:"cause,omitempty" type:"string" required:"true" enum:"CompleteWorkflowExecutionFailedCause"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the CompleteWorkflowExecution decision to complete this
	// execution. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`
}

type DecisionTaskCompletedEventAttributes struct {
	_ struct{} `type:"structure"`

	// User defined context for the workflow execution.
	ExecutionContext *string `json:"executionContext,omitempty" type:"string"`

	// The ID of the DecisionTaskScheduled event that was recorded when this decision
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`

	// The ID of the DecisionTaskStarted event recorded when this decision task
	// was started. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"startedEventId,omitempty" type:"long" required:"true"`
}

type DecisionTaskScheduledEventAttributes struct {
	_ struct{} `type:"structure"`

	// The maximum duration for this decision task. The task is considered timed
	// out if it doesn't completed within this duration.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	StartToCloseTimeout *string `json:"startToCloseTimeout,omitempty" type:"string"`

	// The name of the task list in which the decision task was scheduled.
	//
	// TaskList is a required field
	TaskList *TaskList `json:"taskList,omitempty" type:"structure" required:"true"`

	// A task priority that, if set, specifies the priority for this decision task.
	// Valid values are integers that range from Java's Integer.MIN_VALUE (-2147483648)
	// to Integer.MAX_VALUE (2147483647). Higher numbers indicate higher priority.
	//
	// For more information about setting task priority, see Setting Task Priority
	// (http://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	TaskPriority *string `json:"taskPriority,omitempty" type:"string"`
}

type DecisionTaskStartedEventAttributes struct {
	_ struct{} `type:"structure"`

	// Identity of the decider making the request. This enables diagnostic tracing
	// when problems arise. The form of this identity is user defined.
	Identity *string `json:"identity,omitempty" type:"string"`

	// The ID of the DecisionTaskScheduled event that was recorded when this decision
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`
}

type DecisionTaskTimedOutEventAttributes struct {
	_ struct{} `type:"structure"`

	// The ID of the DecisionTaskScheduled event that was recorded when this decision
	// task was scheduled. This information can be useful for diagnosing problems
	// by tracing back the chain of events leading up to this event.
	//
	// ScheduledEventId is a required field
	ScheduledEventId *int64 `json:"scheduledEventId,omitempty" type:"long" required:"true"`

	// The ID of the DecisionTaskStarted event recorded when this decision task
	// was started. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"startedEventId,omitempty" type:"long" required:"true"`

	// The type of timeout that expired before the decision task could be completed.
	//
	// TimeoutType is a required field
	TimeoutType *string `json:"timeoutType,omitempty" type:"string" required:"true" enum:"DecisionTaskTimeoutType"`
}

type FailWorkflowExecutionFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The cause of the failure. This information is generated by the system and
	// can be useful for diagnostic purposes.
	//
	// If cause is set to OPERATION_NOT_PERMITTED, the decision failed because it
	// lacked sufficient permissions. For details and example IAM policies, see
	// Using IAM to Manage Access to Amazon SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
	// in the Amazon SWF Developer Guide.
	//
	// Cause is a required field
	Cause *string `json:"cause,omitempty" type:"string" required:"true" enum:"FailWorkflowExecutionFailedCause"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the FailWorkflowExecution decision to fail this execution.
	// This information can be useful for diagnosing problems by tracing back the
	// chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`
}

type RequestCancelActivityTaskFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The activityId provided in the RequestCancelActivityTask decision that failed.
	//
	// ActivityId is a required field
	ActivityId *string `json:"activityId,omitempty" min:"1" type:"string" required:"true"`

	// The cause of the failure. This information is generated by the system and
	// can be useful for diagnostic purposes.
	//
	// If cause is set to OPERATION_NOT_PERMITTED, the decision failed because it
	// lacked sufficient permissions. For details and example IAM policies, see
	// Using IAM to Manage Access to Amazon SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
	// in the Amazon SWF Developer Guide.
	//
	// Cause is a required field
	Cause *string `json:"cause,omitempty" type:"string" required:"true" enum:"RequestCancelActivityTaskFailedCause"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the RequestCancelActivityTask decision for this cancellation
	// request. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`
}

type ScheduleActivityTaskFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The activityId provided in the ScheduleActivityTask decision that failed.
	//
	// ActivityId is a required field
	ActivityId *string `json:"activityId,omitempty" min:"1" type:"string" required:"true"`

	// The activity type provided in the ScheduleActivityTask decision that failed.
	//
	// ActivityType is a required field
	ActivityType *ActivityType `json:"activityType,omitempty" type:"structure" required:"true"`

	// The cause of the failure. This information is generated by the system and
	// can be useful for diagnostic purposes.
	//
	// If cause is set to OPERATION_NOT_PERMITTED, the decision failed because it
	// lacked sufficient permissions. For details and example IAM policies, see
	// Using IAM to Manage Access to Amazon SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
	// in the Amazon SWF Developer Guide.
	//
	// Cause is a required field
	Cause *string `json:"cause,omitempty" type:"string" required:"true" enum:"ScheduleActivityTaskFailedCause"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision that
	// resulted in the scheduling of this activity task. This information can be
	// useful for diagnosing problems by tracing back the chain of events leading
	// up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`
}

type WorkflowExecutionCancelRequestedEventAttributes struct {
	_ struct{} `type:"structure"`

	// If set, indicates that the request to cancel the workflow execution was automatically
	// generated, and specifies the cause. This happens if the parent workflow execution
	// times out or is terminated, and the child policy is set to cancel child executions.
	Cause *string `json:"cause,omitempty" type:"string" enum:"WorkflowExecutionCancelRequestedCause"`

	// The ID of the RequestCancelExternalWorkflowExecutionInitiated event corresponding
	// to the RequestCancelExternalWorkflowExecution decision to cancel this workflow
	// execution.The source event with this ID can be found in the history of the
	// source workflow execution. This information can be useful for diagnosing
	// problems by tracing back the chain of events leading up to this event.
	ExternalInitiatedEventId *int64 `json:"externalInitiatedEventId,omitempty" type:"long"`

	// The external workflow execution for which the cancellation was requested.
	ExternalWorkflowExecution *WorkflowExecution `json:"externalWorkflowExecution,omitempty" type:"structure"`
}

type WorkflowExecutionCanceledEventAttributes struct {
	_ struct{} `type:"structure"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the CancelWorkflowExecution decision for this cancellation
	// request. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`

	// The details of the cancellation.
	Details *string `json:"details,omitempty" type:"string"`
}

type WorkflowExecutionCompletedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the CompleteWorkflowExecution decision to complete this
	// execution. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`

	// The result produced by the workflow execution upon successful completion.
	Result *string `json:"result,omitempty" type:"string"`
}

type WorkflowExecutionFailedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The ID of the DecisionTaskCompleted event corresponding to the decision task
	// that resulted in the FailWorkflowExecution decision to fail this execution.
	// This information can be useful for diagnosing problems by tracing back the
	// chain of events leading up to this event.
	//
	// DecisionTaskCompletedEventId is a required field
	DecisionTaskCompletedEventId *int64 `json:"decisionTaskCompletedEventId,omitempty" type:"long" required:"true"`

	// The details of the failure.
	Details *string `json:"details,omitempty" type:"string"`

	// The descriptive reason provided for the failure.
	Reason *string `json:"reason,omitempty" type:"string"`
}

type WorkflowExecutionStartedEventAttributes struct {
	_ struct{} `type:"structure"`

	// The policy to use for the child workflow executions if this workflow execution
	// is terminated, by calling the TerminateWorkflowExecution action explicitly
	// or due to an expired timeout.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	//
	// ChildPolicy is a required field
	ChildPolicy *string `json:"childPolicy,omitempty" type:"string" required:"true" enum:"ChildPolicy"`

	// If this workflow execution was started due to a ContinueAsNewWorkflowExecution
	// decision, then it contains the runId of the previous workflow execution that
	// was closed and continued as this execution.
	ContinuedExecutionRunId *string `json:"continuedExecutionRunId,omitempty" type:"string"`

	// The maximum duration for this workflow execution.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	ExecutionStartToCloseTimeout *string `json:"executionStartToCloseTimeout,omitempty" type:"string"`

	// The input provided to the workflow execution.
	Input *string `json:"input,omitempty" type:"string"`

	// The IAM role attached to the workflow execution.
	LambdaRole *string `json:"lambdaRole,omitempty" min:"1" type:"string"`

	// The ID of the StartChildWorkflowExecutionInitiated event corresponding to
	// the StartChildWorkflowExecutionDecision to start this workflow execution.
	// The source event with this ID can be found in the history of the source workflow
	// execution. This information can be useful for diagnosing problems by tracing
	// back the chain of events leading up to this event.
	ParentInitiatedEventId *int64 `json:"parentInitiatedEventId,omitempty" type:"long"`

	// The source workflow execution that started this workflow execution. The member
	// isn't set if the workflow execution was not started by a workflow.
	ParentWorkflowExecution *WorkflowExecution `json:"parentWorkflowExecution,omitempty" type:"structure"`

	// The list of tags associated with this workflow execution. An execution can
	// have up to 5 tags.
	TagList []*string `json:"tagList,omitempty" type:"list"`

	// The name of the task list for scheduling the decision tasks for this workflow
	// execution.
	//
	// TaskList is a required field
	TaskList *TaskList `json:"taskList,omitempty" type:"structure" required:"true"`

	// The priority of the decision tasks in the workflow execution.
	TaskPriority *string `json:"taskPriority,omitempty" type:"string"`

	// The maximum duration of decision tasks for this workflow type.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	TaskStartToCloseTimeout *string `json:"taskStartToCloseTimeout,omitempty" type:"string"`

	// The workflow type of this execution.
	//
	// WorkflowType is a required field
	WorkflowType *WorkflowType `json:"workflowType,omitempty" type:"structure" required:"true"`
}

type WorkflowExecutionTerminatedEventAttributes struct {
	_ struct{} `type:"structure"`

	// If set, indicates that the workflow execution was automatically terminated,
	// and specifies the cause. This happens if the parent workflow execution times
	// out or is terminated and the child policy is set to terminate child executions.
	Cause *string `json:"cause,omitempty" type:"string" enum:"WorkflowExecutionTerminatedCause"`

	// The policy used for the child workflow executions of this workflow execution.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	//
	// ChildPolicy is a required field
	ChildPolicy *string `json:"childPolicy,omitempty" type:"string" required:"true" enum:"ChildPolicy"`

	// The details provided for the termination.
	Details *string `json:"details,omitempty" type:"string"`

	// The reason provided for the termination.
	Reason *string `json:"reason,omitempty" type:"string"`
}

type WorkflowExecutionTimedOutEventAttributes struct {
	_ struct{} `type:"structure"`

	// The policy used for the child workflow executions of this workflow execution.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	//
	// ChildPolicy is a required field
	ChildPolicy *string `json:"childPolicy,omitempty" type:"string" required:"true" enum:"ChildPolicy"`

	// The type of timeout that caused this event.
	//
	// TimeoutType is a required field
	TimeoutType *string `json:"timeoutType,omitempty" type:"string" required:"true" enum:"WorkflowExecutionTimeoutType"`
}

const (
	// EventTypeWorkflowExecutionStarted is a EventType enum value
	EventTypeWorkflowExecutionStarted = "WorkflowExecutionStarted"

	// EventTypeWorkflowExecutionCancelRequested is a EventType enum value
	EventTypeWorkflowExecutionCancelRequested = "WorkflowExecutionCancelRequested"

	// EventTypeWorkflowExecutionCompleted is a EventType enum value
	EventTypeWorkflowExecutionCompleted = "WorkflowExecutionCompleted"

	// EventTypeCompleteWorkflowExecutionFailed is a EventType enum value
	EventTypeCompleteWorkflowExecutionFailed = "CompleteWorkflowExecutionFailed"

	// EventTypeWorkflowExecutionFailed is a EventType enum value
	EventTypeWorkflowExecutionFailed = "WorkflowExecutionFailed"

	// EventTypeFailWorkflowExecutionFailed is a EventType enum value
	EventTypeFailWorkflowExecutionFailed = "FailWorkflowExecutionFailed"

	// EventTypeWorkflowExecutionTimedOut is a EventType enum value
	EventTypeWorkflowExecutionTimedOut = "WorkflowExecutionTimedOut"

	// EventTypeWorkflowExecutionCanceled is a EventType enum value
	EventTypeWorkflowExecutionCanceled = "WorkflowExecutionCanceled"

	// EventTypeCancelWorkflowExecutionFailed is a EventType enum value
	EventTypeCancelWorkflowExecutionFailed = "CancelWorkflowExecutionFailed"

	// EventTypeWorkflowExecutionContinuedAsNew is a EventType enum value
	EventTypeWorkflowExecutionContinuedAsNew = "WorkflowExecutionContinuedAsNew"

	// EventTypeContinueAsNewWorkflowExecutionFailed is a EventType enum value
	EventTypeContinueAsNewWorkflowExecutionFailed = "ContinueAsNewWorkflowExecutionFailed"

	// EventTypeWorkflowExecutionTerminated is a EventType enum value
	EventTypeWorkflowExecutionTerminated = "WorkflowExecutionTerminated"

	// EventTypeDecisionTaskScheduled is a EventType enum value
	EventTypeDecisionTaskScheduled = "DecisionTaskScheduled"

	// EventTypeDecisionTaskStarted is a EventType enum value
	EventTypeDecisionTaskStarted = "DecisionTaskStarted"

	// EventTypeDecisionTaskCompleted is a EventType enum value
	EventTypeDecisionTaskCompleted = "DecisionTaskCompleted"

	// EventTypeDecisionTaskTimedOut is a EventType enum value
	EventTypeDecisionTaskTimedOut = "DecisionTaskTimedOut"

	// EventTypeActivityTaskScheduled is a EventType enum value
	EventTypeActivityTaskScheduled = "ActivityTaskScheduled"

	// EventTypeScheduleActivityTaskFailed is a EventType enum value
	EventTypeScheduleActivityTaskFailed = "ScheduleActivityTaskFailed"

	// EventTypeActivityTaskStarted is a EventType enum value
	EventTypeActivityTaskStarted = "ActivityTaskStarted"

	// EventTypeActivityTaskCompleted is a EventType enum value
	EventTypeActivityTaskCompleted = "ActivityTaskCompleted"

	// EventTypeActivityTaskFailed is a EventType enum value
	EventTypeActivityTaskFailed = "ActivityTaskFailed"

	// EventTypeActivityTaskTimedOut is a EventType enum value
	EventTypeActivityTaskTimedOut = "ActivityTaskTimedOut"

	// EventTypeActivityTaskCanceled is a EventType enum value
	EventTypeActivityTaskCanceled = "ActivityTaskCanceled"

	// EventTypeActivityTaskCancelRequested is a EventType enum value
	EventTypeActivityTaskCancelRequested = "ActivityTaskCancelRequested"

	// EventTypeRequestCancelActivityTaskFailed is a EventType enum value
	EventTypeRequestCancelActivityTaskFailed = "RequestCancelActivityTaskFailed"

	// EventTypeWorkflowExecutionSignaled is a EventType enum value
	EventTypeWorkflowExecutionSignaled = "WorkflowExecutionSignaled"

	// EventTypeMarkerRecorded is a EventType enum value
	EventTypeMarkerRecorded = "MarkerRecorded"

	// EventTypeRecordMarkerFailed is a EventType enum value
	EventTypeRecordMarkerFailed = "RecordMarkerFailed"

	// EventTypeTimerStarted is a EventType enum value
	EventTypeTimerStarted = "TimerStarted"

	// EventTypeStartTimerFailed is a EventType enum value
	EventTypeStartTimerFailed = "StartTimerFailed"

	// EventTypeTimerFired is a EventType enum value
	EventTypeTimerFired = "TimerFired"

	// EventTypeTimerCanceled is a EventType enum value
	EventTypeTimerCanceled = "TimerCanceled"

	// EventTypeCancelTimerFailed is a EventType enum value
	EventTypeCancelTimerFailed = "CancelTimerFailed"

	// EventTypeStartChildWorkflowExecutionInitiated is a EventType enum value
	EventTypeStartChildWorkflowExecutionInitiated = "StartChildWorkflowExecutionInitiated"

	// EventTypeStartChildWorkflowExecutionFailed is a EventType enum value
	EventTypeStartChildWorkflowExecutionFailed = "StartChildWorkflowExecutionFailed"

	// EventTypeChildWorkflowExecutionStarted is a EventType enum value
	EventTypeChildWorkflowExecutionStarted = "ChildWorkflowExecutionStarted"

	// EventTypeChildWorkflowExecutionCompleted is a EventType enum value
	EventTypeChildWorkflowExecutionCompleted = "ChildWorkflowExecutionCompleted"

	// EventTypeChildWorkflowExecutionFailed is a EventType enum value
	EventTypeChildWorkflowExecutionFailed = "ChildWorkflowExecutionFailed"

	// EventTypeChildWorkflowExecutionTimedOut is a EventType enum value
	EventTypeChildWorkflowExecutionTimedOut = "ChildWorkflowExecutionTimedOut"

	// EventTypeChildWorkflowExecutionCanceled is a EventType enum value
	EventTypeChildWorkflowExecutionCanceled = "ChildWorkflowExecutionCanceled"

	// EventTypeChildWorkflowExecutionTerminated is a EventType enum value
	EventTypeChildWorkflowExecutionTerminated = "ChildWorkflowExecutionTerminated"

	// EventTypeSignalExternalWorkflowExecutionInitiated is a EventType enum value
	EventTypeSignalExternalWorkflowExecutionInitiated = "SignalExternalWorkflowExecutionInitiated"

	// EventTypeSignalExternalWorkflowExecutionFailed is a EventType enum value
	EventTypeSignalExternalWorkflowExecutionFailed = "SignalExternalWorkflowExecutionFailed"

	// EventTypeExternalWorkflowExecutionSignaled is a EventType enum value
	EventTypeExternalWorkflowExecutionSignaled = "ExternalWorkflowExecutionSignaled"

	// EventTypeRequestCancelExternalWorkflowExecutionInitiated is a EventType enum value
	EventTypeRequestCancelExternalWorkflowExecutionInitiated = "RequestCancelExternalWorkflowExecutionInitiated"

	// EventTypeRequestCancelExternalWorkflowExecutionFailed is a EventType enum value
	EventTypeRequestCancelExternalWorkflowExecutionFailed = "RequestCancelExternalWorkflowExecutionFailed"

	// EventTypeExternalWorkflowExecutionCancelRequested is a EventType enum value
	EventTypeExternalWorkflowExecutionCancelRequested = "ExternalWorkflowExecutionCancelRequested"

	// EventTypeLambdaFunctionScheduled is a EventType enum value
	EventTypeLambdaFunctionScheduled = "LambdaFunctionScheduled"

	// EventTypeLambdaFunctionStarted is a EventType enum value
	EventTypeLambdaFunctionStarted = "LambdaFunctionStarted"

	// EventTypeLambdaFunctionCompleted is a EventType enum value
	EventTypeLambdaFunctionCompleted = "LambdaFunctionCompleted"

	// EventTypeLambdaFunctionFailed is a EventType enum value
	EventTypeLambdaFunctionFailed = "LambdaFunctionFailed"

	// EventTypeLambdaFunctionTimedOut is a EventType enum value
	EventTypeLambdaFunctionTimedOut = "LambdaFunctionTimedOut"

	// EventTypeScheduleLambdaFunctionFailed is a EventType enum value
	EventTypeScheduleLambdaFunctionFailed = "ScheduleLambdaFunctionFailed"

	// EventTypeStartLambdaFunctionFailed is a EventType enum value
	EventTypeStartLambdaFunctionFailed = "StartLambdaFunctionFailed"
)

const (
	// ExecutionStatusOpen is a ExecutionStatus enum value
	ExecutionStatusOpen = "OPEN"

	// ExecutionStatusClosed is a ExecutionStatus enum value
	ExecutionStatusClosed = "CLOSED"
)
