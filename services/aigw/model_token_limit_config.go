package aigw

type TokenLimitConfig struct {
	TimeUnit         *string             `json:"time_unit,omitempty"`
	TimeWindow       *float32            `json:"time_window,omitempty"`
	TokenAmount      *int32              `json:"token_amount,omitempty"`
	Limits           *map[string]int32   `json:"limits,omitempty"`
	TokenUsageWeight *map[string]float32 `json:"token_usage_weight,omitempty"`
}
