package privatezone

type UnbindVpcToRuleRequest struct {
	RuleId      *string      `json:"-"`
	ClientToken *string      `json:"-"`
	VpcRegions  []*VpcRegion `json:"vpcRegions,omitempty"`
}
