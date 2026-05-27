package csn

type AddRouteRuleRequest struct {
	CsnRtId     *string `json:"-"`
	ClientToken *string `json:"-"`
	AttachId    *string `json:"attachId,omitempty"`
	DestAddress *string `json:"destAddress,omitempty"`
	RouteType   *string `json:"routeType,omitempty"`
}
