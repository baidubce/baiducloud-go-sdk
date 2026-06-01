package et

type UpdateDedicatedChannelRouteRulesRequest struct {
	EtId        *string `json:"-"`
	EtChannelId *string `json:"-"`
	RouteRuleId *string `json:"-"`
	ClientToken *string `json:"-"`
	Description *string `json:"description,omitempty"`
}
