package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetRoleListResponse struct {
	bce.BaseResponse
	Roles []*InstanceRoleModel `json:"roles,omitempty"`
}
