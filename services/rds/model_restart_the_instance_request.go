package rds

type RestartTheInstanceRequest struct {
	InstanceId    *string `json:"-"`
	EffectiveTime *string `json:"effectiveTime,omitempty"`
}
