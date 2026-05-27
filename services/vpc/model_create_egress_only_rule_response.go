package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEgressOnlyRuleResponse struct {
	bce.BaseResponse
	EgressOnlyRuleId *string `json:"egressOnlyRuleId,omitempty"`
}
