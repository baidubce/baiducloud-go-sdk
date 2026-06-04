package privatezone

type UnbindVpcToRuleRequest struct {
	RuleId     *string      `json:"-"`
	ClienToken *string      `json:"-"`
	VpcRegions []*VpcRegion `json:"vpcRegions,omitempty"`
}
