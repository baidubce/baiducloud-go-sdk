package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateRateLimitRuleResponse struct {
	bce.BaseResponse
	RateLimitRuleId *string `json:"rateLimitRuleId,omitempty"`
}
