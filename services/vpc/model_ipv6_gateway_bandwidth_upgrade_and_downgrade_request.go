package vpc

type Ipv6GatewayBandwidthUpgradeAndDowngradeRequest struct {
	GatewayId       *string `json:"-"`
	ClientToken     *string `json:"-"`
	BandwidthInMbps *int32  `json:"bandwidthInMbps,omitempty"`
}
