package vpc

type DeleteIpv6OnlyAccessPolicyRequest struct {
	GatewayId        *string `json:"-"`
	EgressOnlyRuleId *string `json:"-"`
	ClientToken      *string `json:"-"`
}
