package apm

type DescribeMetricDataRequest struct {
	Action                 *string        `json:"-"`
	Metrics                []*MetricQuery `json:"metrics,omitempty"`
	MetricsName            *string        `json:"metrics.name,omitempty"`
	MetricsCompareTo       []*string      `json:"metrics.compareTo,omitempty"`
	MetricsFilters         []*Filter      `json:"metrics.filters,omitempty"`
	BeginDatetime          *string        `json:"beginDatetime,omitempty"`
	EndDatetime            *string        `json:"endDatetime,omitempty"`
	Filters                []*Filter      `json:"filters,omitempty"`
	GroupBy                []*string      `json:"groupBy,omitempty"`
	OrderBy                *string        `json:"orderBy,omitempty"`
	Order                  *string        `json:"order,omitempty"`
	Limit                  *int32         `json:"limit,omitempty"`
	PeriodSeconds          *int32         `json:"periodSeconds,omitempty"`
	ReserveEmptyDimensions *bool          `json:"reserveEmptyDimensions,omitempty"`
}
