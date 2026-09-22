package scs

type ShardBandwidth struct {
	ShardName         *string `json:"shardName,omitempty"`
	NodeBandwidthInMB *int32  `json:"nodeBandwidthInMB,omitempty"`
}
