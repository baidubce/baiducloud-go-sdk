package bls

type UpdateLogStoreRequest struct {
	LogStoreName          *string `json:"-"`
	Project               *string `json:"-"`
	Retention             *int32  `json:"retention,omitempty"`
	Tags                  []*Tag  `json:"tags,omitempty"`
	ShardCount            *int32  `json:"shardCount,omitempty"`
	MaxShardCount         *int32  `json:"maxShardCount,omitempty"`
	DisableShardAutoSplit *bool   `json:"disableShardAutoSplit,omitempty"`
	IndexEnabled          *bool   `json:"indexEnabled,omitempty"`
	HotRetention          *int32  `json:"hotRetention,omitempty"`
}
