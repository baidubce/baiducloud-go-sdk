package scs

type ShardLog struct {
	ShardShowId *string    `json:"shardShowId,omitempty"`
	TotalNum    *int32     `json:"totalNum,omitempty"`
	LogItem     []*LogItem `json:"logItem,omitempty"`
	ShardId     *int32     `json:"shardId,omitempty"`
}
