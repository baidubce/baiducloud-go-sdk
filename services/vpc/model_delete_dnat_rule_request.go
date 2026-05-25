package vpc

type DeleteDnatRuleRequest struct {
	NatId       *string `json:"-"`
	RuleId      *string `json:"-"`
	ClientToken *string `json:"-"`
}
