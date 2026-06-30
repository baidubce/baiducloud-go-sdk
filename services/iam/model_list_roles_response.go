package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListRolesResponse struct {
	bce.BaseResponse
	Roles []*RoleModel `json:"roles,omitempty"`
}
