package vpc

type RateLimitRule struct {
	RateLimitRuleId        *string `json:"rateLimitRuleId,omitempty"`
	Ipv6Address            *string `json:"ipv6Address,omitempty"`
	IngressBandwidthInMbps *int32  `json:"ingressBandwidthInMbps,omitempty"`
	EgressBandwidthInMbps  *int32  `json:"egressBandwidthInMbps,omitempty"`
}
