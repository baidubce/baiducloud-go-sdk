package bcm

type DescribeMetricDataLatestTopRequest struct {
	Action              *string   `json:"-"`
	Scope               *string   `json:"scope,omitempty"`
	Region              *string   `json:"region,omitempty"`
	EndDatetime         *string   `json:"endDatetime,omitempty"`
	MetricName          *string   `json:"metricName,omitempty"`
	Filters             []*Filter `json:"filters,omitempty"`
	Limit               *int32    `json:"limit,omitempty"`
	Asc                 *bool     `json:"asc,omitempty"`
	PeriodSeconds       *int32    `json:"periodSeconds,omitempty"`
	AggregationOverTime *string   `json:"aggregationOverTime,omitempty"`
}
