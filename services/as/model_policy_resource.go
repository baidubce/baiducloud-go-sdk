package as

type PolicyResource struct {
	Identifiers      []*Dimension `json:"identifiers"`
	MetricDimensions []*Dimension `json:"metricDimensions"`
}
