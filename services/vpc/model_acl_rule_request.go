package vpc

type AclRuleRequest struct {
	SubnetId             *string `json:"subnetId,omitempty"`
	Description          *string `json:"description,omitempty"`
	Protocol             *string `json:"protocol,omitempty"`
	SourceIpAddress      *string `json:"sourceIpAddress,omitempty"`
	DestinationIpAddress *string `json:"destinationIpAddress,omitempty"`
	SourcePort           *string `json:"sourcePort,omitempty"`
	DestinationPort      *string `json:"destinationPort,omitempty"`
	Position             *int32  `json:"position,omitempty"`
	Direction            *string `json:"direction,omitempty"`
	IpVersion            *int32  `json:"ipVersion,omitempty"`
	Action               *string `json:"action,omitempty"`
}
