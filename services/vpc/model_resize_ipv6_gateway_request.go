package vpc

type ResizeIpv6GatewayRequest struct {
	GatewayId       *string `json:"-"`
	ClientToken     *string `json:"-"`
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
}
