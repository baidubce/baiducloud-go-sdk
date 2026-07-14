package as

type ExecRuleV2Request struct {
	GroupId  *string `json:"-"`
	ExecRule *string `json:"-"`
	RuleId   *string `json:"ruleId,omitempty"`
}
