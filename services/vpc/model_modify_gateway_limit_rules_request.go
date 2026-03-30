package vpc

type ModifyGatewayLimitRulesRequest struct {
	GlrId       *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Bandwidth   *int32  `json:"bandwidth,omitempty"`
}
