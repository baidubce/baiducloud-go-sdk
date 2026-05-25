package vpc

type UpdateIpv6GatewaySpeedLimitPolicyRequest struct {
	GatewayId              *string `json:"-"`
	RateLimitRuleId        *string `json:"-"`
	ClientToken            *string `json:"-"`
	IngressBandwidthInMbps *int32  `json:"ingressBandwidthInMbps,omitempty"`
	EgressBandwidthInMbps  *int32  `json:"egressBandwidthInMbps,omitempty"`
}
