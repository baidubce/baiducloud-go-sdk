package vpc

type CreateAnIpv6GatewayRequest struct {
	ClientToken     *string     `json:"-"`
	VpcId           *string     `json:"vpcId,omitempty"`
	Name            *string     `json:"name,omitempty"`
	BandwidthInMbps *int32      `json:"bandwidthInMbps,omitempty"`
	Billing         *Billing    `json:"billing,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	ResourceGroupId *string     `json:"resourceGroupId,omitempty"`
	DeleteProtect   *bool       `json:"deleteProtect,omitempty"`
}
