package vpc

type CreateIpv6GatewaySpeedLimitPolicyRequest struct {
	GatewayId              *string `json:"-"`
	ClientToken            *string `json:"-"`
	Ipv6Address            *string `json:"ipv6Address,omitempty"`
	IngressBandwidthInMbps *int32  `json:"ingressBandwidthInMbps,omitempty"`
	EgressBandwidthInMbps  *int32  `json:"egressBandwidthInMbps,omitempty"`
}
