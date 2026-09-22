package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SyncGroupDetailResponse struct {
	bce.BaseResponse
	SyncGroupShowId     *string        `json:"syncGroupShowId,omitempty"`
	SyncGroupName       *string        `json:"syncGroupName,omitempty"`
	Status              *string        `json:"status,omitempty"`
	ClusterNum          *int32         `json:"clusterNum,omitempty"`
	NodeType            *string        `json:"nodeType,omitempty"`
	NetConn             *string        `json:"netConn,omitempty"`
	ConfilctResolution  *string        `json:"confilctResolution,omitempty"`
	SyncGroupCreateTime *string        `json:"syncGroupCreateTime,omitempty"`
	SameSpec            *bool          `json:"sameSpec,omitempty"`
	SameShardNum        *bool          `json:"sameShardNum,omitempty"`
	Cluster             []*ClusterItem `json:"cluster,omitempty"`
}
