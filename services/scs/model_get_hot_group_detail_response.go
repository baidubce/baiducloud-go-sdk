package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetHotGroupDetailResponse struct {
	bce.BaseResponse
	Leader          *Leader      `json:"leader,omitempty"`
	Followers       []*Followers `json:"followers,omitempty"`
	GroupId         *string      `json:"groupId,omitempty"`
	GroupName       *string      `json:"groupName,omitempty"`
	GroupStatus     *string      `json:"groupStatus,omitempty"`
	ClusterNum      *int32       `json:"clusterNum,omitempty"`
	GroupCreateTime *string      `json:"groupCreateTime,omitempty"`
	ForbidWrite     *int32       `json:"forbidWrite,omitempty"`
	GroupType       *string      `json:"groupType,omitempty"`
}
