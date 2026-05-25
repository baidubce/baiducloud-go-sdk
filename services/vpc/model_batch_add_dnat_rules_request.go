package vpc

type BatchAddDnatRulesRequest struct {
	NatId       *string            `json:"-"`
	ClientToken *string            `json:"-"`
	Rules       []*DnatRuleRequest `json:"rules,omitempty"`
}
