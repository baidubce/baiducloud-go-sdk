package bcm

type Metric struct {
	Name                *string   `json:"name,omitempty"`
	Label               *string   `json:"label,omitempty"`
	ResourceIdentifiers []*string `json:"resourceIdentifiers,omitempty"`
	MetricDimensions    []*string `json:"metricDimensions,omitempty"`
	Period              *float64  `json:"period,omitempty"`
	PeriodUnit          *string   `json:"periodUnit,omitempty"`
	Unit                *string   `json:"unit,omitempty"`
}
