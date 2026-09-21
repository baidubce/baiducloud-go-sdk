package dbsc

type SCSSlowLogSummaryItem struct {
	Fingerprint  *string  `json:"fingerprint,omitempty"`
	ExecuteTimes *int64   `json:"executeTimes,omitempty"`
	DurationSum  *float32 `json:"durationSum,omitempty"`
	DurationMax  *float32 `json:"durationMax,omitempty"`
	DurationAvg  *float32 `json:"durationAvg,omitempty"`
}
