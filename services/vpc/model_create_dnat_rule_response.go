package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDnatRuleResponse struct {
	bce.BaseResponse
	RuleId *string `json:"ruleId,omitempty"`
}
