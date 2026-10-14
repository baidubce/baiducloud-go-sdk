package rds

type Progress struct {
	Step         *string `json:"step,omitempty"`
	Status       *string `json:"status,omitempty"`
	Description  *string `json:"description,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	FinishTime   *string `json:"finishTime,omitempty"`
	Percentage   *int32  `json:"percentage,omitempty"`
	EstimateTime *int32  `json:"estimateTime,omitempty"`
}
