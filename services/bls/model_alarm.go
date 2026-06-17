package bls

type Alarm struct {
	Id               *string            `json:"id,omitempty"`
	GroupId          *string            `json:"groupId,omitempty"`
	StartTime        *string            `json:"startTime,omitempty"`
	EndTime          *string            `json:"endTime,omitempty"`
	State            *string            `json:"state,omitempty"`
	CloseReason      *string            `json:"closeReason,omitempty"`
	Policy           *Policy            `json:"policy,omitempty"`
	Object           *LogStore          `json:"object,omitempty"`
	TriggerCondition *TriggerCondition  `json:"triggerCondition,omitempty"`
	Groups           *map[string]string `json:"groups,omitempty"`
	Executions       []*Execution       `json:"executions,omitempty"`
}
