package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BindRoleResponse struct {
	bce.BaseResponse
	FailInstances            []*InstancePassRoleFailModel    `json:"failInstances,omitempty"`
	InstanceRoleAssociations []*InstanceRoleAssociationModel `json:"instanceRoleAssociations,omitempty"`
}
