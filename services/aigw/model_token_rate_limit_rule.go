package aigw

type TokenRateLimitRule struct {
	MatchCondition *map[string]string `json:"match_condition,omitempty"`
	LimitConfig    *TokenLimitConfig  `json:"limit_config,omitempty"`
}
