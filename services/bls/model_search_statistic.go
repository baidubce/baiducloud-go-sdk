package bls

type SearchStatistic struct {
	Interval  *int32    `json:"interval,omitempty"`
	StartTime *string   `json:"startTime,omitempty"`
	EndTime   *string   `json:"endTime,omitempty"`
	Histogram []*Bucket `json:"histogram,omitempty"`
}
