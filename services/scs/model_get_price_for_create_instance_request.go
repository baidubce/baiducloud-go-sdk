package scs

type GetPriceForCreateInstanceRequest struct {
	Engine            *int32  `json:"engine,omitempty"`
	ClusterType       *string `json:"clusterType,omitempty"`
	NodeType          *string `json:"nodeType,omitempty"`
	CacheInstanceType *int32  `json:"cacheInstanceType,omitempty"`
	ShardNum          *int32  `json:"shardNum,omitempty"`
	ReplicationNum    *int32  `json:"replicationNum,omitempty"`
	InstanceNum       *int32  `json:"instanceNum,omitempty"`
	DiskType          *string `json:"diskType,omitempty"`
	DiskFlavor        *int32  `json:"diskFlavor,omitempty"`
	ChargeType        *string `json:"chargeType,omitempty"`
	Period            *int32  `json:"period,omitempty"`
	TimeUnit          *string `json:"timeUnit,omitempty"`
}
