package aigw

type RuleItem struct {
	MatchCondition *interface{} `json:"match_condition,omitempty"`
	LimitConfig    *interface{} `json:"limit_config,omitempty"`
}
