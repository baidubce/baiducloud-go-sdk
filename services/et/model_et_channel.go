package et

type EtChannel struct {
	Id                  *string     `json:"id,omitempty"`
	Name                *string     `json:"name,omitempty"`
	Status              *string     `json:"status,omitempty"`
	BgpStatus           *string     `json:"bgpStatus,omitempty"`
	BgpRouteLimit       *int32      `json:"bgpRouteLimit,omitempty"`
	Ipv6BgpStatus       *string     `json:"ipv6BgpStatus,omitempty"`
	BaiduAddress        *string     `json:"baiduAddress,omitempty"`
	AuthorizedUsers     []*string   `json:"authorizedUsers,omitempty"`
	Description         *string     `json:"description,omitempty"`
	Networks            []*string   `json:"networks,omitempty"`
	CustomerAddress     *string     `json:"customerAddress,omitempty"`
	RouteType           *string     `json:"routeType,omitempty"`
	VlanId              *int32      `json:"vlanId,omitempty"`
	EnableIpv6          *int32      `json:"enableIpv6,omitempty"`
	BaiduIpv6Address    *string     `json:"baiduIpv6Address,omitempty"`
	CustomerIpv6Address *string     `json:"customerIpv6Address,omitempty"`
	Ipv6Networks        []*string   `json:"ipv6Networks,omitempty"`
	SendInterval        *int32      `json:"sendInterval,omitempty"`
	ReceivInterval      *int32      `json:"receivInterval,omitempty"`
	DetectMultiplier    *int32      `json:"detectMultiplier,omitempty"`
	Tags                []*TagModel `json:"tags,omitempty"`
}
