package cfw

type UpdateCfwRuleRequest struct {
	CfwId         *string `json:"-"`
	CfwRuleId     *string `json:"-"`
	IpVersion     *int32  `json:"ipVersion,omitempty"`
	Priority      *int32  `json:"priority,omitempty"`
	Protocol      *string `json:"protocol,omitempty"`
	Direction     *string `json:"direction,omitempty"`
	SourceAddress *string `json:"sourceAddress,omitempty"`
	DestAddress   *string `json:"destAddress,omitempty"`
	SourcePort    *string `json:"sourcePort,omitempty"`
	DestPort      *string `json:"destPort,omitempty"`
	Action        *string `json:"action,omitempty"`
	Description   *string `json:"description,omitempty"`
}
