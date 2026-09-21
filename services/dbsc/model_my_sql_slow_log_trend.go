package dbsc

type MySQLSlowLogTrend struct {
	Value     *int32  `json:"value,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}
