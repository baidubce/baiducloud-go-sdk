package blb

type BackendServerStatus struct {
	InstanceId *string `json:"instanceId,omitempty"`
	Weight     *int32  `json:"weight,omitempty"`
	Status     *string `json:"status,omitempty"`
}
