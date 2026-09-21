package vdb

type ModifyInstanceConfigRequest struct {
	InstanceId  *string                     `json:"instanceId,omitempty"`
	Reason      *string                     `json:"reason,omitempty"`
	UserConfigs []*InstanceConfigUserConfig `json:"userConfigs,omitempty"`
}
