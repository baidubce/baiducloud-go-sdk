package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchCreationOfPermissionGroupRulesResponse struct {
	bce.BaseResponse
	Responses []*CreateAccessRuleResponse `json:"responses,omitempty"`
}
