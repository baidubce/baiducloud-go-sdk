package et

type EtChannelRouteRule struct {
	RouteRuleId *string `json:"routeRuleId,omitempty"`
	IpVersion   *int32  `json:"ipVersion,omitempty"`
	DestAddress *string `json:"destAddress,omitempty"`
	NexthopType *string `json:"nexthopType,omitempty"`
	NexthopId   *string `json:"nexthopId,omitempty"`
	Description *string `json:"description,omitempty"`
}
