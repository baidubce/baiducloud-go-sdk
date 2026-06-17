package bls

type LogStoreDetail struct {
	CreationDateTime       *string `json:"creationDateTime,omitempty"`
	DisableShardAutoSplit  *bool   `json:"disableShardAutoSplit,omitempty"`
	EnableArchiveRetention *bool   `json:"enableArchiveRetention,omitempty"`
	EnableHotRetention     *bool   `json:"enableHotRetention,omitempty"`
	HotRetention           *int32  `json:"hotRetention,omitempty"`
	IndexEnabled           *bool   `json:"indexEnabled,omitempty"`
	LastModifiedTime       *string `json:"lastModifiedTime,omitempty"`
	LogStoreName           *string `json:"logStoreName,omitempty"`
	LowFrequencyRetention  *int32  `json:"lowFrequencyRetention,omitempty"`
	MaxShardCount          *int32  `json:"maxShardCount,omitempty"`
	Project                *string `json:"project,omitempty"`
	ResourceID             *string `json:"resourceID,omitempty"`
	Retention              *int32  `json:"retention,omitempty"`
	ShardCount             *int32  `json:"shardCount,omitempty"`
	ShortID                *string `json:"shortID,omitempty"`
	Tags                   []*Tag  `json:"tags,omitempty"`
}
