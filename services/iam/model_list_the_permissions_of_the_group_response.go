package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListThePermissionsOfTheGroupResponse struct {
	bce.BaseResponse
	Policies []*PolicyModel `json:"policies,omitempty"`
}
