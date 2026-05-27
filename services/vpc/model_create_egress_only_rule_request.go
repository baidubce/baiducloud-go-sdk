package vpc

type CreateEgressOnlyRuleRequest struct {
	GatewayId   *string `json:"-"`
	ClientToken *string `json:"-"`
	Cidr        *string `json:"cidr,omitempty"`
}
