package aigw

type TimeoutPolicy struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Timeout *int32 `json:"timeout,omitempty"`
}
