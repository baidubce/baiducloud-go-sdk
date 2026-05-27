package csn

type Instance struct {
	AttachId          *string `json:"attachId,omitempty"`
	InstanceType      *string `json:"instanceType,omitempty"`
	InstanceId        *string `json:"instanceId,omitempty"`
	InstanceName      *string `json:"instanceName,omitempty"`
	InstanceRegion    *string `json:"instanceRegion,omitempty"`
	InstanceAccountId *string `json:"instanceAccountId,omitempty"`
	Status            *string `json:"status,omitempty"`
}
