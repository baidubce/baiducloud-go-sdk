package apm

type DescribeTraceMetricDataRequest struct {
	Action        *string             `json:"-"`
	BeginDatetime *string             `json:"beginDatetime,omitempty"`
	EndDatetime   *string             `json:"endDatetime,omitempty"`
	Metrics       []*TraceMetricQuery `json:"metrics,omitempty"`
	MetricsName   *string             `json:"metrics.name,omitempty"`
	Filters       []*Filter           `json:"filters,omitempty"`
	GroupBy       []*string           `json:"groupBy,omitempty"`
	PeriodSeconds *int32              `json:"periodSeconds,omitempty"`
	Aggregate     []*string           `json:"aggregate,omitempty"`
}
