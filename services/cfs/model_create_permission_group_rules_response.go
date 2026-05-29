package cfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreatePermissionGroupRulesResponse struct {
	bce.BaseResponse
	AccessRuleId *int32 `json:"accessRuleId,omitempty"`
}
