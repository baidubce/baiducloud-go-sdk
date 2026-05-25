package vpc

type AddIpv6OnlyOutboundAndNoInboundPolicyRequest struct {
	GatewayId   *string `json:"-"`
	ClientToken *string `json:"-"`
	Cidr        *string `json:"cidr,omitempty"`
}
