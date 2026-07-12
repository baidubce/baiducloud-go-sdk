package oos

type ExecutionTaskSummary struct {
	Id                *string                 `json:"id,omitempty"`
	LoopIndex         *int32                  `json:"loopIndex,omitempty"`
	Namespace         *string                 `json:"namespace,omitempty"`
	Dag               *DagInstance            `json:"dag,omitempty"`
	Revision          *int64                  `json:"revision,omitempty"`
	CreatedTimestamp  *int64                  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  *int64                  `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp *int64                  `json:"finishedTimestamp,omitempty"`
	State             *string                 `json:"state,omitempty"`
	Operator          *TaskOperatorSummary    `json:"operator,omitempty"`
	Reason            *string                 `json:"reason,omitempty"`
	ErrorCode         *string                 `json:"errorCode,omitempty"`
	InitContext       *interface{}            `json:"initContext,omitempty"`
	Context           *interface{}            `json:"context,omitempty"`
	OutputContext     *interface{}            `json:"outputContext,omitempty"`
	Tries             *int32                  `json:"tries,omitempty"`
	Children          []*ExecutionTaskSummary `json:"children,omitempty"`
	Log               []*Log                  `json:"log,omitempty"`
}
