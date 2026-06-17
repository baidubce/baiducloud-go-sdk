package bcm

type Timeseries struct {
	MetricName *string      `json:"metricName,omitempty"`
	Dimensions []*Dimension `json:"dimensions,omitempty"`
	DataPoints []*DataPoint `json:"dataPoints,omitempty"`
}
