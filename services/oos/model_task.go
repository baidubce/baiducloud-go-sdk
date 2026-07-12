package oos

type Task struct {
	Namespace         *string      `json:"namespace,omitempty"`
	UserId            *string      `json:"userId,omitempty"`
	Id                *string      `json:"id,omitempty"`
	Revision          *int64       `json:"revision,omitempty"`
	LoopIndex         *int32       `json:"loopIndex,omitempty"`
	Reason            *string      `json:"reason,omitempty"`
	ErrorCode         *string      `json:"errorCode,omitempty"`
	Execution         *Execution   `json:"execution,omitempty"`
	Operator          *Operator    `json:"operator,omitempty"`
	CreatedTimestamp  *int64       `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  *int64       `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp *int64       `json:"finishedTimestamp,omitempty"`
	State             *string      `json:"state,omitempty"`
	Properties        *interface{} `json:"properties,omitempty"`
	Tries             *int32       `json:"tries,omitempty"`
	Children          []*Task      `json:"children,omitempty"`
	Executions        []*Execution `json:"executions,omitempty"`
	OutputContext     *interface{} `json:"outputContext,omitempty"`
}
