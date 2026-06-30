package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListTheUserSPermissionsResponse struct {
	bce.BaseResponse
	Policies []*PolicyModel `json:"policies,omitempty"`
}
