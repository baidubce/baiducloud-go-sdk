package vpc

type SnatRule struct {
	RuleId           *string   `json:"ruleId,omitempty"`
	RuleName         *string   `json:"ruleName,omitempty"`
	PublicIpsAddress []*string `json:"publicIpsAddress,omitempty"`
	SourceCIDR       *string   `json:"sourceCIDR,omitempty"`
	Status           *string   `json:"status,omitempty"`
}
