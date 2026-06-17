package bls

type Template struct {
	Retention             *int32    `json:"retention,omitempty"`
	ShardCount            *int32    `json:"shardCount,omitempty"`
	DisableShardAutoSplit *bool     `json:"disableShardAutoSplit,omitempty"`
	MaxShardCount         *int32    `json:"maxShardCount,omitempty"`
	EnableHotRetention    *bool     `json:"enableHotRetention,omitempty"`
	HotRetention          *int32    `json:"hotRetention,omitempty"`
	Index                 *Index    `json:"index,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	ProjectPatterns       []*string `json:"projectPatterns,omitempty"`
	LogstorePatterns      []*string `json:"logstorePatterns,omitempty"`
	Priority              *int32    `json:"priority,omitempty"`
	CreatedTimestamp      *string   `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp      *string   `json:"updatedTimestamp,omitempty"`
}
