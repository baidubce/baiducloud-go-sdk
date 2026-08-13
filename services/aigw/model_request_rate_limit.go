package aigw

type RequestRateLimit struct {
	RuleName  *string     `json:"rule_name,omitempty"`
	Enabled   *bool       `json:"enabled,omitempty"`
	RuleItems []*RuleItem `json:"rule_items,omitempty"`
}
