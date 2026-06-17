package bcm

type DescribeDimensionValuesRequest struct {
	Action        *string   `json:"-"`
	Scope         *string   `json:"scope,omitempty"`
	ResourceType  *string   `json:"resourceType,omitempty"`
	Region        *string   `json:"region,omitempty"`
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	MetricName    *string   `json:"metricName,omitempty"`
	DimensionKey  *string   `json:"dimensionKey,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
}
