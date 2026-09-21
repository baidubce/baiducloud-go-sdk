package dbsc

type MySQLSlowLogStatsByDuration struct {
	Start *int32  `json:"start,omitempty"`
	End   *int32  `json:"end,omitempty"`
	Title *string `json:"title,omitempty"`
	Count *int32  `json:"count,omitempty"`
}
