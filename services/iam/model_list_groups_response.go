package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListGroupsResponse struct {
	bce.BaseResponse
	Groups []*GroupModel `json:"groups,omitempty"`
}
