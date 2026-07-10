package blb

type AppBackendServerForCreate struct {
	InstanceId *string `json:"instanceId,omitempty"`
	Weight     *int32  `json:"weight,omitempty"`
	Desc       *string `json:"desc,omitempty"`
}
