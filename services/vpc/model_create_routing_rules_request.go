package vpc

type CreateRoutingRulesRequest struct {
	ClientToken        *string    `json:"-"`
	RouteTableId       *string    `json:"routeTableId,omitempty"`
	SourceAddress      *string    `json:"sourceAddress,omitempty"`
	DestinationAddress *string    `json:"destinationAddress,omitempty"`
	NexthopId          *string    `json:"nexthopId,omitempty"`
	NexthopType        *string    `json:"nexthopType,omitempty"`
	NextHopList        []*NextHop `json:"nextHopList,omitempty"`
	Description        *string    `json:"description,omitempty"`
}
