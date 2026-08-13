package aigw

type RequestRateLimitRule struct {
	MatchCondition *map[string]string `json:"match_condition,omitempty"`
	LimitConfig    *map[string]string `json:"limit_config,omitempty"`
}
