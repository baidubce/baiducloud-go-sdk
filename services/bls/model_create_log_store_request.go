package bls

type CreateLogStoreRequest struct {
	Project               *string `json:"project,omitempty"`
	LogStoreName          *string `json:"logStoreName,omitempty"`
	Retention             *int32  `json:"retention,omitempty"`
	Tags                  []*Tag  `json:"tags,omitempty"`
	Index                 *Index  `json:"index,omitempty"`
	ShardCount            *int32  `json:"shardCount,omitempty"`
	MaxShardCount         *int32  `json:"maxShardCount,omitempty"`
	DisableShardAutoSplit *bool   `json:"disableShardAutoSplit,omitempty"`
	IndexEnabled          *bool   `json:"indexEnabled,omitempty"`
	HotRetention          *int32  `json:"hotRetention,omitempty"`
	EnableHotRetention    *bool   `json:"EnableHotRetention,omitempty"`
}
