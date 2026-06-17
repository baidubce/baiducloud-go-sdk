package bcm

type DescribeMetricDataRequest struct {
	Action              *string   `json:"-"`
	Scope               *string   `json:"scope,omitempty"`
	ResourceType        *string   `json:"resourceType,omitempty"`
	Region              *string   `json:"region,omitempty"`
	BeginDatetime       *string   `json:"beginDatetime,omitempty"`
	EndDatetime         *string   `json:"endDatetime,omitempty"`
	MetricName          *string   `json:"metricName,omitempty"`
	Filters             []*Filter `json:"filters,omitempty"`
	Limit               *int32    `json:"limit,omitempty"`
	Offset              *int32    `json:"offset,omitempty"`
	PeriodSeconds       *int32    `json:"periodSeconds,omitempty"`
	AggregationOverTime []*string `json:"aggregationOverTime,omitempty"`
}
