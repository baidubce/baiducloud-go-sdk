package oos

type TaskOperatorSummary struct {
	Name                   *string                `json:"name,omitempty"`
	Description            *string                `json:"description,omitempty"`
	Tags                   *interface{}           `json:"tags,omitempty"`
	Operator               *string                `json:"operator,omitempty"`
	DagSpec                *DagSpec               `json:"dagSpec,omitempty"`
	Inline                 *bool                  `json:"inline,omitempty"`
	Retries                *int32                 `json:"retries,omitempty"`
	RetryInterval          *int32                 `json:"retryInterval,omitempty"`
	Timeout                *int32                 `json:"timeout,omitempty"`
	InitContext            *interface{}           `json:"initContext,omitempty"`
	Loops                  []*interface{}         `json:"loops,omitempty"`
	ParallelismRatio       *float64               `json:"parallelismRatio,omitempty"`
	ParallelismCount       *int32                 `json:"parallelismCount,omitempty"`
	AllowedFailureRatio    *float64               `json:"allowedFailureRatio,omitempty"`
	AllowedFailureCount    *int32                 `json:"allowedFailureCount,omitempty"`
	Manually               *bool                  `json:"manually,omitempty"`
	PauseOnFailure         *bool                  `json:"pauseOnFailure,omitempty"`
	ScheduleDelayMilli     *int32                 `json:"scheduleDelayMilli,omitempty"`
	WaitOnAgentMilli       *int32                 `json:"waitOnAgentMilli,omitempty"`
	Condition              *interface{}           `json:"condition,omitempty"`
	Breakpoints            []*int32               `json:"breakpoints,omitempty"`
	TriggerRule            *string                `json:"triggerRule,omitempty"`
	LoopWindowType         *string                `json:"loopWindowType,omitempty"`
	WorkerSelectors        []*TagSelector         `json:"workerSelectors,omitempty"`
	CollectChildrenContext *string                `json:"collectChildrenContext,omitempty"`
	RollbackOperator       *TaskOperatorSummary   `json:"rollbackOperator,omitempty"`
	Events                 []*EventModel          `json:"events,omitempty"`
	InitOperators          []*TaskOperatorSummary `json:"initOperators,omitempty"`
	ByBsmAgent             *bool                  `json:"byBsmAgent,omitempty"`
}
