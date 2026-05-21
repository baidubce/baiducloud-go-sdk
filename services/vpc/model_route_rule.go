package vpc

type RouteRule struct {
	RouteRuleId        *string `json:"routeRuleId,omitempty"`
	RouteTableId       *string `json:"routeTableId,omitempty"`
	SourceAddress      *string `json:"sourceAddress,omitempty"`
	DestinationAddress *string `json:"destinationAddress,omitempty"`
	NexthopId          *string `json:"nexthopId,omitempty"`
	NexthopType        *string `json:"nexthopType,omitempty"`
	PathType           *string `json:"pathType,omitempty"`
	Description        *string `json:"description,omitempty"`
}
