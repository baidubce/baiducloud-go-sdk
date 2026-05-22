package cfw

type DeleteCfwRuleRequest struct {
	CfwId      *string   `json:"-"`
	CfwRuleIds []*string `json:"cfwRuleIds,omitempty"`
}
