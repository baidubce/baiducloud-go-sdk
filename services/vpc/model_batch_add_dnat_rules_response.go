package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchAddDnatRulesResponse struct {
	bce.BaseResponse
	RuleIds []*string `json:"ruleIds,omitempty"`
}
