package as

type PolicyResource struct {
	Identifiers      []*Dimension `json:"identifiers,omitempty"`
	MetricDimensions []*Dimension `json:"metricDimensions,omitempty"`
}
