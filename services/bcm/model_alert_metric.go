package bcm

type AlertMetric struct {
	Metric *AlertMetricValue `json:"metric,omitempty"`
	Rule   *AlertMetricRule  `json:"rule,omitempty"`
}
