package swfclient

type DecisionOps interface {
	PollForDecisionTaskPages(*PollForDecisionTaskInput, func(*PollForDecisionTaskOutput, bool) bool) error
}

type ActivityOps interface {
	PollForActivityTask(req *PollForActivityTaskInput) (resp *PollForActivityTaskOutput, err error)
}
