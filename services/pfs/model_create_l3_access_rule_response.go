package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateL3AccessRuleResponse struct {
	bce.BaseResponse
	RequestId    *string                   `json:"requestId,omitempty"`
	AccessRuleId *int32                    `json:"accessRuleId,omitempty"`
	Responses    []*AccessRuleCreateResult `json:"responses,omitempty"`
}
