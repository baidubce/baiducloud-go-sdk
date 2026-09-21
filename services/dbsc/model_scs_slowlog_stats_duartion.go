package dbsc

type SCSSlowlogStatsDuartion struct {
	Start *int64  `json:"start,omitempty"`
	End   *int64  `json:"end,omitempty"`
	Title *string `json:"title,omitempty"`
	Count *int64  `json:"count,omitempty"`
}
