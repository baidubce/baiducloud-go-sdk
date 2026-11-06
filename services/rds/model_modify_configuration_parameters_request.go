package rds

type ModifyConfigurationParametersRequest struct {
	InstanceId    *string      `json:"-"`
	EffectiveTime *string      `json:"effectiveTime,omitempty"`
	Parameters    []*Parameter `json:"parameters,omitempty"`
}
