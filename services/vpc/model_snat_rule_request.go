package vpc

type SnatRuleRequest struct {
	RuleName         *string   `json:"ruleName,omitempty"`
	PublicIpsAddress []*string `json:"publicIpsAddress,omitempty"`
	SourceCIDR       *string   `json:"sourceCIDR,omitempty"`
}
