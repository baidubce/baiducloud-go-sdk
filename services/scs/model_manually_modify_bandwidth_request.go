package scs

type ManuallyModifyBandwidthRequest struct {
	InstanceId         *string           `json:"-"`
	ShardBandwidthInfo []*ShardBandwidth `json:"shardBandwidthInfo,omitempty"`
}
