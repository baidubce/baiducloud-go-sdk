package scs

type InstanceVersionUpgradeRequest struct {
	InstanceId    *string `json:"-"`
	KernelVersion *string `json:"kernelVersion,omitempty"`
	IsDefer       *bool   `json:"isDefer,omitempty"`
}
