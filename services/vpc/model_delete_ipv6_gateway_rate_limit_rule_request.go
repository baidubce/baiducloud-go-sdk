package vpc

type DeleteIpv6GatewayRateLimitRuleRequest struct {
	GatewayId       *string `json:"-"`
	RateLimitRuleId *string `json:"-"`
	ClientToken     *string `json:"-"`
}
