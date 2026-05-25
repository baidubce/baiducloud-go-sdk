package vpc

type BatchAddSnatRulesRequest struct {
	ClientToken *string            `json:"-"`
	NatId       *string            `json:"natId,omitempty"`
	SnatRules   []*SnatRuleRequest `json:"snatRules,omitempty"`
}
