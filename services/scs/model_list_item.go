package scs

type ListItem struct {
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
	UserName            *string        `json:"userName,omitempty"`
	UpdateStatus        *int32         `json:"updateStatus,omitempty"`
	Extra               *string        `json:"extra,omitempty"`
	UserType            *int32         `json:"userType,omitempty"`
}
