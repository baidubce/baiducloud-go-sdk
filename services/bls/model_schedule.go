package bls

type Schedule struct {
	IntervalMinute *int32 `json:"intervalMinute,omitempty"`
	FixTimeMinute  *int32 `json:"fixTimeMinute,omitempty"`
	DayOfWeek      *int32 `json:"dayOfWeek,omitempty"`
}
