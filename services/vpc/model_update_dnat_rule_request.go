package vpc

type UpdateDnatRuleRequest struct {
	NatId            *string `json:"-"`
	RuleId           *string `json:"-"`
	ClientToken      *string `json:"-"`
	RuleName         *string `json:"ruleName,omitempty"`
	Protocol         *string `json:"protocol,omitempty"`
	PublicPort       *int32  `json:"publicPort,omitempty"`
	PrivatePort      *int32  `json:"privatePort,omitempty"`
	PublicPortRange  *string `json:"publicPortRange,omitempty"`
	PrivatePortRange *string `json:"privatePortRange,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
}
