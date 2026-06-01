package et

type RemoveDedicatedChannelRouteParametersRequest struct {
	EtId         *string   `json:"-"`
	EtChannelId  *string   `json:"-"`
	ClientToken  *string   `json:"-"`
	Networks     []*string `json:"networks,omitempty"`
	Ipv6Networks []*string `json:"ipv6Networks,omitempty"`
	RouteType    *string   `json:"routeType,omitempty"`
}
