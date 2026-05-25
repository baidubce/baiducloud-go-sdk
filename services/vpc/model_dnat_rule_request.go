package vpc

type DnatRuleRequest struct {
	RuleName         *string `json:"ruleName,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	Protocol         *string `json:"protocol,omitempty"`
	PublicPort       *int32  `json:"publicPort,omitempty"`
	PrivatePort      *int32  `json:"privatePort,omitempty"`
	PublicPortRange  *string `json:"publicPortRange,omitempty"`
	PrivatePortRange *string `json:"privatePortRange,omitempty"`
}
