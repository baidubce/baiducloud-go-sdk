package dbsc

type TrendResultBase struct {
	Category []*string            `json:"category,omitempty"`
	Series   []*TrendResultSeries `json:"series,omitempty"`
}
