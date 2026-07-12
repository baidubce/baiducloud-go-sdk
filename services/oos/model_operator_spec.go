package oos

type OperatorSpec struct {
	Name                *string         `json:"name,omitempty"`
	Description         *string         `json:"description,omitempty"`
	Tags                []*KeyValuePair `json:"tags,omitempty"`
	Operator            *string         `json:"operator,omitempty"`
	Label               *string         `json:"label,omitempty"`
	Template            *Template       `json:"template,omitempty"`
	Retries             *int32          `json:"retries,omitempty"`
	RetryInterval       *int32          `json:"retryInterval,omitempty"`
	Timeout             *int32          `json:"timeout,omitempty"`
	ParallelismRatio    *float64        `json:"parallelismRatio,omitempty"`
	ParallelismCount    *int32          `json:"parallelismCount,omitempty"`
	AllowedFailureRatio *float64        `json:"allowedFailureRatio,omitempty"`
	AllowedFailureCount *int32          `json:"allowedFailureCount,omitempty"`
	Manually            *bool           `json:"manually,omitempty"`
	ScheduleDelayMilli  *int32          `json:"scheduleDelayMilli,omitempty"`
	WaitOnAgentMilli    *int32          `json:"waitOnAgentMilli,omitempty"`
	AutoRollback        *bool           `json:"autoRollback,omitempty"`
	PauseOnFailure      *bool           `json:"pauseOnFailure,omitempty"`
	Properties          []*Property     `json:"properties,omitempty"`
	Output              []*Property     `json:"output,omitempty"`
	InitContext         *interface{}    `json:"initContext,omitempty"`
	Loops               []*LoopModel    `json:"loops,omitempty"`
	Parallel            *bool           `json:"parallel,omitempty"`
}
