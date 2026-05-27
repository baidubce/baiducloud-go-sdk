package csn

type CsnRtPropagation struct {
	AttachId       *string `json:"attachId,omitempty"`
	Description    *string `json:"description,omitempty"`
	InstanceId     *string `json:"instanceId,omitempty"`
	InstanceName   *string `json:"instanceName,omitempty"`
	InstanceRegion *string `json:"instanceRegion,omitempty"`
	InstanceType   *string `json:"instanceType,omitempty"`
	Status         *string `json:"status,omitempty"`
}
