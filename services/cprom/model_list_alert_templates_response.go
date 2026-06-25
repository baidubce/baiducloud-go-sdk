package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListAlertTemplatesResponse struct {
	bce.BaseResponse
	RuleTemplates []*RuleTemplate `json:"ruleTemplates,omitempty"`
}
