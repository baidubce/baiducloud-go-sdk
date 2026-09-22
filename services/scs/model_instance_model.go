package scs

type InstanceModel struct {
	InstanceId         *string       `json:"instanceId,omitempty"`
	InstanceName       *string       `json:"instanceName,omitempty"`
	InstanceStatus     *string       `json:"instanceStatus,omitempty"`
	IsolatedStatus     *string       `json:"isolatedStatus,omitempty"`
	ClusterType        *string       `json:"clusterType,omitempty"`
	Engine             *string       `json:"engine,omitempty"`
	EngineVersion      *string       `json:"engineVersion,omitempty"`
	VnetIp             *string       `json:"vnetIp,omitempty"`
	Domain             *string       `json:"domain,omitempty"`
	Port               *int32        `json:"port,omitempty"`
	InstanceCreateTime *string       `json:"instanceCreateTime,omitempty"`
	Capacity           *float32      `json:"capacity,omitempty"`
	UsedCapacity       *float64      `json:"usedCapacity,omitempty"`
	PaymentTiming      *string       `json:"paymentTiming,omitempty"`
	ZoneNames          []*string     `json:"zoneNames,omitempty"`
	DiskFlavor         *int32        `json:"diskFlavor,omitempty"`
	Eip                *string       `json:"eip,omitempty"`
	InstanceExpireTime *string       `json:"instanceExpireTime,omitempty"`
	ReplicationNum     *int32        `json:"replicationNum,omitempty"`
	NodeType           *string       `json:"nodeType,omitempty"`
	StoreType          *int32        `json:"storeType,omitempty"`
	ShardNum           *int32        `json:"shardNum,omitempty"`
	Tags               []*Tag        `json:"tags,omitempty"`
	ResourceGroupId    *string       `json:"resourceGroupId,omitempty"`
	ResourceGroupName  *string       `json:"resourceGroupName,omitempty"`
	OrderStatus        *string       `json:"orderStatus,omitempty"`
	DeployIdList       []*string     `json:"deployIdList,omitempty"`
	Vpc                *VpcInfo      `json:"vpc,omitempty"`
	Subnets            []*SubnetInfo `json:"subnets,omitempty"`
}
