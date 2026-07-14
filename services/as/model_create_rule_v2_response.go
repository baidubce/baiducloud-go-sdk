package as

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateRuleV2Response struct {
	bce.BaseResponse
	RuleId *string `json:"ruleId,omitempty"`
}
