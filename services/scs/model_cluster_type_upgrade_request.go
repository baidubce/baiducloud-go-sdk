package scs

type ClusterTypeUpgradeRequest struct {
	InstanceId      *string            `json:"-"`
	IsDefer         *bool              `json:"isDefer,omitempty"`
	NodeType        *string            `json:"nodeType,omitempty"`
	ShardNum        *int32             `json:"shardNum,omitempty"`
	ReplicationInfo []*ReplicationItem `json:"replicationInfo,omitempty"`
}
