package vpc

type DnatRule struct {
	RuleId           *string `json:"ruleId,omitempty"`
	RuleName         *string `json:"ruleName,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	Protocol         *string `json:"protocol,omitempty"`
	PublicPort       *string `json:"publicPort,omitempty"`
	PrivatePort      *string `json:"privatePort,omitempty"`
	Status           *string `json:"status,omitempty"`
	PrivatePortRange *string `json:"privatePortRange,omitempty"`
	PublicPortRange  *string `json:"publicPortRange,omitempty"`
}
