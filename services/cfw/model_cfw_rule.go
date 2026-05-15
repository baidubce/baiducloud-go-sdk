package cfw

type CfwRule struct {
	CfwId         *string `json:"cfwId,omitempty"`
	CfwRuleId     *string `json:"cfwRuleId,omitempty"`
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
