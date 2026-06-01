package apm

type DescribeServicesMetricsRequest struct {
	BeginDatetime *string         `json:"beginDatetime,omitempty"`
	EndDatetime   *string         `json:"endDatetime,omitempty"`
	Services      []*string       `json:"services,omitempty"`
	Metrics       []*string       `json:"metrics,omitempty"`
	MetricFilters []*MetricFilter `json:"metricFilters,omitempty"`
}
