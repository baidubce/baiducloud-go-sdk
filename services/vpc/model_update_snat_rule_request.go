package vpc

type UpdateSnatRuleRequest struct {
	NatId            *string   `json:"-"`
	RuleId           *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	RuleName         *string   `json:"ruleName,omitempty"`
	SourceCIDR       *string   `json:"sourceCIDR,omitempty"`
	PublicIpsAddress []*string `json:"publicIpsAddress,omitempty"`
}
