package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateMetaSyncRuleResponse struct {
	bce.BaseResponse
	MetaSyncRuleId *string `json:"metaSyncRuleId,omitempty"`
}
