package vpc

type CreateSnatRuleRequest struct {
	NatId            *string   `json:"-"`
	ClientToken      *string   `json:"-"`
	RuleName         *string   `json:"ruleName,omitempty"`
	PublicIpsAddress []*string `json:"publicIpsAddress,omitempty"`
	SourceCIDR       *string   `json:"sourceCIDR,omitempty"`
}
