package vdb

type CreateInstanceUsingPOSTRequest struct {
	EngineType        *string            `json:"-"`
	AutoRenew         *bool              `json:"autoRenew,omitempty"`
	AutoRenewTime     *int32             `json:"autoRenewTime,omitempty"`
	AutoRenewTimeUnit *string            `json:"autoRenewTimeUnit,omitempty"`
	Components        []*MilvusComponent `json:"components,omitempty"`
	Duration          *int32             `json:"duration,omitempty"`
	Env               *string            `json:"env,omitempty"`
	InstanceParam     *InstanceParam     `json:"instanceParam,omitempty"`
	ProductType       *string            `json:"productType,omitempty"`
	TimeUnit          *string            `json:"timeUnit,omitempty"`
}
