package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateIpv6GatewaySpeedLimitPolicyResponse struct {
	bce.BaseResponse
	RateLimitRuleId *string `json:"rateLimitRuleId,omitempty"`
}
