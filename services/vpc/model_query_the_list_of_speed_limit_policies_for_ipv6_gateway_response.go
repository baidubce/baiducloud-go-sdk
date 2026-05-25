package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheListOfSpeedLimitPoliciesForIpv6GatewayResponse struct {
	bce.BaseResponse
	RateLimitRules []*RateLimitRule `json:"rateLimitRules,omitempty"`
	Marker         *string          `json:"marker,omitempty"`
	IsTruncated    *bool            `json:"isTruncated,omitempty"`
	NextMarker     *string          `json:"nextMarker,omitempty"`
	MaxKeys        *int32           `json:"maxKeys,omitempty"`
}
