package scs

type ClusterItem struct {
	ClusterShowId     *string         `json:"clusterShowId,omitempty"`
	ClusterName       *string         `json:"clusterName,omitempty"`
	Region            *string         `json:"region,omitempty"`
	ClusterStatus     *string         `json:"clusterStatus,omitempty"`
	ClusterEngine     *string         `json:"clusterEngine,omitempty"`
	CreateTime        *string         `json:"createTime,omitempty"`
	TotalCapacityInGb *float32        `json:"totalCapacityInGb,omitempty"`
	UsedCapacityInGb  *float32        `json:"usedCapacityInGb,omitempty"`
	ExpiredTime       *string         `json:"expiredTime,omitempty"`
	ShardList         []*string       `json:"shardList,omitempty"`
	SyncFlow          []*SyncFlowItem `json:"syncFlow,omitempty"`
}
