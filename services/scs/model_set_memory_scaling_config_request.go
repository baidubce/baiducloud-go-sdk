package scs

type SetMemoryScalingConfigRequest struct {
	InstanceId *string  `json:"-"`
	MemSpec    *MemSpec `json:"memSpec,omitempty"`
}
