package iam

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListTheSubjectsGrantedPermissionsResponse struct {
	bce.BaseResponse
	Entities   []*AttachedEntities `json:"entities,omitempty"`
	Id         *string             `json:"id,omitempty"`
	Name       *string             `json:"name,omitempty"`
	IamType    *string             `json:"type,omitempty"`
	AttachTime *string             `json:"attach_time,omitempty"`
}
