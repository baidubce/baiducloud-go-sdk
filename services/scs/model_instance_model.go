package scs

type InstanceModel struct {
	InstanceId         *string   `json:"instanceId,omitempty"`
	InstanceName       *string   `json:"instanceName,omitempty"`
	InstanceStatus     *string   `json:"instanceStatus,omitempty"`
	IsolatedStatus     *string   `json:"isolatedStatus,omitempty"`
	ClusterType        *string   `json:"clusterType,omitempty"`
	Engine             *string   `json:"engine,omitempty"`
	EngineVersion      *string   `json:"engineVersion,omitempty"`
	VnetIp             *string   `json:"vnetIp,omitempty"`
	Domain             *string   `json:"domain,omitempty"`
	Port               *string   `json:"port,omitempty"`
	InstanceCreateTime *string   `json:"instanceCreateTime,omitempty"`
	InstanceExpireTime *string   `json:"instanceExpireTime,omitempty"`
	Capacity           *float32  `json:"capacity,omitempty"`
	UsedCapacity       *float64  `json:"usedCapacity,omitempty"`
	PaymentTiming      *string   `json:"paymentTiming,omitempty"`
	ZoneNames          []*string `json:"zoneNames,omitempty"`
	OrderStatus        *string   `json:"orderStatus,omitempty"`
	DeployIdList       []*string `json:"deployIdList,omitempty"`
}
