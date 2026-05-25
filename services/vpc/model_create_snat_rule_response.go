package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSnatRuleResponse struct {
	bce.BaseResponse
	RuleId *string `json:"ruleId,omitempty"`
}
