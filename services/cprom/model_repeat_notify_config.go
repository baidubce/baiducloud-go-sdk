package cprom

type RepeatNotifyConfig struct {
	Enabled      *bool   `json:"enabled,omitempty"`
	IntervalHour *int32  `json:"intervalHour,omitempty"`
	IntervalMin  *int32  `json:"intervalMin,omitempty"`
	MaxCount     *int32  `json:"maxCount,omitempty"`
	Strategy     *string `json:"strategy,omitempty"`
}
