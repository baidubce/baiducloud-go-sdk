package aigw

type TokenRateLimit struct {
	RuleName                        *string           `json:"rule_name,omitempty"`
	Enabled                         *bool             `json:"enabled,omitempty"`
	PreReserveRemainingRatio        *float32          `json:"pre_reserve_remaining_ratio,omitempty"`
	PreReserveHistoryWindowSeconds  *int32            `json:"pre_reserve_history_window_seconds,omitempty"`
	PreReserveSafetyFactor          *float32          `json:"pre_reserve_safety_factor,omitempty"`
	PreReserveEstimationMode        *string           `json:"pre_reserve_estimation_mode,omitempty"`
	PreReserveInitialTokens         *map[string]int32 `json:"pre_reserve_initial_tokens,omitempty"`
	SlidingWindowBucketCount        *int32            `json:"sliding_window_bucket_count,omitempty"`
	PreReserveAdmissionMode         *string           `json:"pre_reserve_admission_mode,omitempty"`
	PreReserveAdmissionBurstSeconds *int32            `json:"pre_reserve_admission_burst_seconds,omitempty"`
	PreReserveRetryJitterMs         *int32            `json:"pre_reserve_retry_jitter_ms,omitempty"`
	RuleItems                       []*RuleItem       `json:"rule_items,omitempty"`
}
