package cfw

type CreateCfwRuleRequest struct {
	CfwId    *string       `json:"-"`
	CfwRules []*CreateRule `json:"cfwRules,omitempty"`
}
