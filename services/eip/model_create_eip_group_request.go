package eip

type CreateEipGroupRequest struct {
	ClientToken     *string     `json:"-"`
	RouteType       *string     `json:"routeType,omitempty"`
	EipCount        *int32      `json:"eipCount,omitempty"`
	Eipv6Count      *int32      `json:"eipv6Count,omitempty"`
	BandwidthInMbps *int32      `json:"bandwidthInMbps,omitempty"`
	Billing         *Billing    `json:"billing,omitempty"`
	Name            *string     `json:"name,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	ResourceGroupId *string     `json:"resourceGroupId,omitempty"`
}
