package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryIpv6GatewayResponse struct {
	bce.BaseResponse
	GatewayId       *string           `json:"gatewayId,omitempty"`
	Name            *string           `json:"name,omitempty"`
	BandwidthInMbps *int32            `json:"bandwidthInMbps,omitempty"`
	VpcId           *string           `json:"vpcId,omitempty"`
	EgressOnlyRules []*EgressOnlyRule `json:"egressOnlyRules,omitempty"`
	RateLimitRules  []*RateLimitRule  `json:"rateLimitRules,omitempty"`
	DeleteProtect   *bool             `json:"deleteProtect,omitempty"`
}
