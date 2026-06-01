package apm

type DescribeLLMMetricDataRequest struct {
	BeginDatetime *string        `json:"beginDatetime,omitempty"`
	EndDatetime   *string        `json:"endDatetime,omitempty"`
	Metrics       []*MetricQuery `json:"metrics,omitempty"`
	Filters       []*Filter      `json:"filters,omitempty"`
	GroupBy       []*string      `json:"groupBy,omitempty"`
	OrderBy       *string        `json:"orderBy,omitempty"`
	Order         *string        `json:"order,omitempty"`
	Limit         *int32         `json:"limit,omitempty"`
	PeriodSeconds *int32         `json:"periodSeconds,omitempty"`
	Aggregate     []*string      `json:"aggregate,omitempty"`
}
