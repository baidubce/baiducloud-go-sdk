package vpc

type DeleteIpv6GatewaySpeedLimitPolicyRequest struct {
	GatewayId       *string `json:"-"`
	RateLimitRuleId *string `json:"-"`
	ClientToken     *string `json:"-"`
}
