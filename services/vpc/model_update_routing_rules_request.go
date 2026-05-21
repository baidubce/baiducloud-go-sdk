package vpc

type UpdateRoutingRulesRequest struct {
	RouteRuleId        *string `json:"-"`
	ClientToken        *string `json:"-"`
	SourceAddress      *string `json:"sourceAddress,omitempty"`
	DestinationAddress *string `json:"destinationAddress,omitempty"`
	NexthopId          *string `json:"nexthopId,omitempty"`
	Description        *string `json:"description,omitempty"`
}
