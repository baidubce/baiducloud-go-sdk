package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetInstanceWhiteGroupResponse struct {
	bce.BaseResponse
	ClusterIPGroups []*ClusterIP `json:"clusterIPGroups,omitempty"`
}
