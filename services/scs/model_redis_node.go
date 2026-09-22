package scs

type RedisNode struct {
	Uuid              *string `json:"uuid,omitempty"`
	NodeShowId        *string `json:"nodeShowId,omitempty"`
	CacheInstanceType *int32  `json:"cacheInstanceType,omitempty"`
	IsReadOnly        *int32  `json:"isReadOnly,omitempty"`
	InGroup           *int32  `json:"inGroup,omitempty"`
	AvailabilityZone  *string `json:"availabilityZone,omitempty"`
	SubnetId          *string `json:"subnetId,omitempty"`
	Status            *int32  `json:"status,omitempty"`
	Weight            *int32  `json:"weight,omitempty"`
	HashName          *string `json:"hashName,omitempty"`
	ShardId           *int32  `json:"shardId,omitempty"`
	NodeId            *int32  `json:"nodeId,omitempty"`
}
