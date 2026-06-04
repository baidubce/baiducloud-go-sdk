package privatezone

type BindVpcToRuleRequest struct {
	RuleId     *string      `json:"-"`
	ClienToken *string      `json:"-"`
	VpcRegions []*VpcRegion `json:"vpcRegions,omitempty"`
}
