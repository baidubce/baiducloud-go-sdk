package vpc

type DeleteIpv6GatewayEgressOnlyRuleRequest struct {
	GatewayId        *string `json:"-"`
	EgressOnlyRuleId *string `json:"-"`
	ClientToken      *string `json:"-"`
}
