package vpc

type DeleteSnatRuleRequest struct {
	NatId       *string `json:"-"`
	RuleId      *string `json:"-"`
	ClientToken *string `json:"-"`
}
