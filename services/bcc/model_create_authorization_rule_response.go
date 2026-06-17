package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAuthorizationRuleResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	RuleId    *string `json:"ruleId,omitempty"`
}
