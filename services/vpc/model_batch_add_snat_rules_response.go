package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchAddSnatRulesResponse struct {
	bce.BaseResponse
	SnatRuleIds []*string `json:"snatRuleIds,omitempty"`
}
