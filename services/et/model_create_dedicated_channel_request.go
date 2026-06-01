package et

type CreateDedicatedChannelRequest struct {
	EtId                *string     `json:"-"`
	ClientToken         *string     `json:"-"`
	AuthorizedUsers     []*string   `json:"authorizedUsers,omitempty"`
	Description         *string     `json:"description,omitempty"`
	BaiduAddress        *string     `json:"baiduAddress,omitempty"`
	Name                *string     `json:"name,omitempty"`
	Networks            []*string   `json:"networks,omitempty"`
	CustomerAddress     *string     `json:"customerAddress,omitempty"`
	RouteType           *string     `json:"routeType,omitempty"`
	VlanId              *int32      `json:"vlanId,omitempty"`
	BgpAsn              *string     `json:"bgpAsn,omitempty"`
	BgpKey              *string     `json:"bgpKey,omitempty"`
	EnableIpv6          *int32      `json:"enableIpv6,omitempty"`
	BaiduIpv6Address    *string     `json:"baiduIpv6Address,omitempty"`
	CustomerIpv6Address *string     `json:"customerIpv6Address,omitempty"`
	Ipv6Networks        []*string   `json:"ipv6Networks,omitempty"`
	Tags                []*TagModel `json:"tags,omitempty"`
}
