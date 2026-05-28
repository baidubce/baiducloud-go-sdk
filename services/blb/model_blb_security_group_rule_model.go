package blb

type BlbSecurityGroupRuleModel struct {
	SecurityGroupRuleId *string `json:"securityGroupRuleId,omitempty"`
	Direction           *string `json:"direction,omitempty"`
	Ethertype           *string `json:"ethertype,omitempty"`
	PortRange           *string `json:"portRange,omitempty"`
	Protocol            *string `json:"protocol,omitempty"`
	SourceGroupId       *string `json:"sourceGroupId,omitempty"`
	SourceIp            *string `json:"sourceIp,omitempty"`
	DestGroupId         *string `json:"destGroupId,omitempty"`
	DestIp              *string `json:"destIp,omitempty"`
}
