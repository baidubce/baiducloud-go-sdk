package cce

type Step struct {
	StepName     *string      `json:"stepName,omitempty"`
	StepStatus   *string      `json:"stepStatus,omitempty"`
	Ready        *bool        `json:"ready,omitempty"`
	StartTime    *string      `json:"startTime,omitempty"`
	FinishedTime *string      `json:"finishedTime,omitempty"`
	CostSeconds  *int32       `json:"costSeconds,omitempty"`
	RetryCount   *int32       `json:"retryCount,omitempty"`
	ErrInfo      *interface{} `json:"errInfo,omitempty"`
}
