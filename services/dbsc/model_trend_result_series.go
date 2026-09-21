package dbsc

type TrendResultSeries struct {
	Name *string        `json:"name,omitempty"`
	Data []*interface{} `json:"data,omitempty"`
}
