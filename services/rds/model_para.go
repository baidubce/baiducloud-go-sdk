package rds

type Para struct {
	Key       *string `json:"key,omitempty"`
	Value     *string `json:"value,omitempty"`
	RdsRange  *string `json:"range,omitempty"`
	Dynamic   *int32  `json:"dynamic,omitempty"`
	Desc      *string `json:"desc,omitempty"`
	Attention *string `json:"attention,omitempty"`
	RdsType   *string `json:"type,omitempty"`
	BestValue *string `json:"bestValue,omitempty"`
	Precision *int32  `json:"precision,omitempty"`
	Required  *bool   `json:"required,omitempty"`
}
