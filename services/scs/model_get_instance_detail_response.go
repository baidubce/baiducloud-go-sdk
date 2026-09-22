package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceDetailResponse struct {
	bce.BaseResponse
	InstanceId              *string                  `json:"instanceId,omitempty"`
	InstanceName            *string                  `json:"instanceName,omitempty"`
	InstanceStatus          *string                  `json:"instanceStatus,omitempty"`
	ClusterType             *string                  `json:"clusterType,omitempty"`
	Engine                  *string                  `json:"engine,omitempty"`
	EngineVersion           *string                  `json:"engineVersion,omitempty"`
	NodeType                *string                  `json:"nodeType,omitempty"`
	StoreType               *int32                   `json:"storeType,omitempty"`
	ShardNum                *int32                   `json:"shardNum,omitempty"`
	VnetIp                  *string                  `json:"vnetIp,omitempty"`
	Eip                     *string                  `json:"eip,omitempty"`
	Domain                  *string                  `json:"domain,omitempty"`
	PublicDomain            *string                  `json:"publicDomain,omitempty"`
	Port                    *int32                   `json:"port,omitempty"`
	InstanceCreateTime      *string                  `json:"instanceCreateTime,omitempty"`
	InstanceExpireTime      *string                  `json:"instanceExpireTime,omitempty"`
	Capacity                *float32                 `json:"capacity,omitempty"`
	UsedCapacity            *float64                 `json:"usedCapacity,omitempty"`
	PaymentTiming           *string                  `json:"paymentTiming,omitempty"`
	ZoneNames               []*string                `json:"zoneNames,omitempty"`
	ReplicationNum          *int32                   `json:"replicationNum,omitempty"`
	ReplicationInfo         []*ReplicationItem       `json:"replicationInfo,omitempty"`
	DiskFlavor              *int32                   `json:"diskFlavor,omitempty"`
	EnableReadOnly          *int32                   `json:"enableReadOnly,omitempty"`
	VpcId                   *string                  `json:"vpcId,omitempty"`
	Subnets                 []*Subnet                `json:"subnets,omitempty"`
	AutoRenew               *bool                    `json:"autoRenew,omitempty"`
	Entrance                *string                  `json:"entrance,omitempty"`
	EnableSlowLog           *int32                   `json:"enableSlowLog,omitempty"`
	BnsGroup                *string                  `json:"bnsGroup,omitempty"`
	ProxyList               []*ProxyItem             `json:"proxyList,omitempty"`
	CacheClusterInstances   []*CacheClusterNode      `json:"cacheClusterInstances,omitempty"`
	RedisList               []*RedisNode             `json:"redisList,omitempty"`
	Tags                    []*Tag                   `json:"tags,omitempty"`
	ResourceGroupId         *string                  `json:"resourceGroupId,omitempty"`
	ResourceGroupName       *string                  `json:"resourceGroupName,omitempty"`
	FullVersionInfo         *InstanceFullVersionInfo `json:"fullVersionInfo,omitempty"`
	MaintainTime            *MaintainTime            `json:"maintainTime,omitempty"`
	EnableHotkey            *bool                    `json:"enableHotkey,omitempty"`
	OrderStatus             *string                  `json:"orderStatus,omitempty"`
	FeatureSwitches         *FeatureSwitches         `json:"featureSwitches,omitempty"`
	CrossAzNearest          *string                  `json:"crossAzNearest,omitempty"`
	EntranceList            []*EntranceItem          `json:"entranceList,omitempty"`
	SupportSentinelCommands *bool                    `json:"supportSentinelCommands,omitempty"`
}
