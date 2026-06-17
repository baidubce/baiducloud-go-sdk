package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLogStoreResponse struct {
	bce.BaseResponse
	Success               *bool   `json:"success,omitempty"`
	Code                  *string `json:"code,omitempty"`
	CreationDateTime      *string `json:"creationDateTime,omitempty"`
	LastModifiedTime      *string `json:"lastModifiedTime,omitempty"`
	Project               *string `json:"project,omitempty"`
	LogStoreName          *string `json:"logStoreName,omitempty"`
	Retention             *int32  `json:"retention,omitempty"`
	ShortID               *string `json:"shortID,omitempty"`
	ShardCount            *int32  `json:"shardCount,omitempty"`
	MaxShardCount         *int32  `json:"maxShardCount,omitempty"`
	DisableShardAutoSplit *bool   `json:"disableShardAutoSplit,omitempty"`
	IndexEnabled          *bool   `json:"indexEnabled,omitempty"`
	HotRetention          *int32  `json:"hotRetention,omitempty"`
	Index                 *Index  `json:"index,omitempty"`
}
