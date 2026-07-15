package cloudassistant

type Target struct {
	InstanceType *string `json:"instanceType,omitempty"`
	InstanceId   *string `json:"instanceId,omitempty"`
	InstanceName *string `json:"instanceName,omitempty"`
	InternalIp   *string `json:"internalIp,omitempty"`
	ExternalIp   *string `json:"externalIp,omitempty"`
	Bandwidth    *string `json:"bandwidth,omitempty"`
}
