package bcm

type AlertMetricRule struct {
	MetricName *string  `json:"metricName,omitempty"`
	Operator   *string  `json:"operator,omitempty"`
	Threshold  *float32 `json:"threshold,omitempty"`
	Statistics *string  `json:"statistics,omitempty"`
	Window     *int32   `json:"window,omitempty"`
}
