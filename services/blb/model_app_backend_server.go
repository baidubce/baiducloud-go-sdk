package blb

type AppBackendServer struct {
	InstanceId *string           `json:"instanceId,omitempty"`
	Weight     *int32            `json:"weight,omitempty"`
	PrivateIp  *string           `json:"privateIp,omitempty"`
	PortList   []*AppRsPortModel `json:"portList,omitempty"`
}
