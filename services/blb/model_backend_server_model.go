package blb

type BackendServerModel struct {
	InstanceId *string `json:"instanceId,omitempty"`
	Weight     *int32  `json:"weight,omitempty"`
}
