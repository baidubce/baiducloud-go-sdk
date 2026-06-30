package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListUserGroupsResponse struct {
	bce.BaseResponse
	Groups []*GroupModel `json:"groups,omitempty"`
}
