package vpc

type UpdateAclRulesRequest struct {
	AclRuleId            *string `json:"-"`
	ClientToken          *string `json:"-"`
	Description          *string `json:"description,omitempty"`
	Protocol             *string `json:"protocol,omitempty"`
	SourceIpAddress      *string `json:"sourceIpAddress,omitempty"`
	DestinationIpAddress *string `json:"destinationIpAddress,omitempty"`
	SourcePort           *string `json:"sourcePort,omitempty"`
	DestinationPort      *string `json:"destinationPort,omitempty"`
	Position             *int32  `json:"position,omitempty"`
	Action               *string `json:"action,omitempty"`
}
