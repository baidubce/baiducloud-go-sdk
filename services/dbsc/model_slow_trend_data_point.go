package dbsc

type SlowTrendDataPoint struct {
	Value     *int64  `json:"value,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}
