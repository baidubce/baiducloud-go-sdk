package aigw

type RetryPolicy struct {
	Enabled         *bool   `json:"enabled,omitempty"`
	RetryConditions *string `json:"retryConditions,omitempty"`
	NumRetries      *int32  `json:"numRetries,omitempty"`
}
