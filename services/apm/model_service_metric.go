package apm

type ServiceMetric struct {
	ServiceName       *string      `json:"serviceName,omitempty"`
	ServiceId         *string      `json:"serviceId,omitempty"`
	Tags              []*Tag       `json:"tags,omitempty"`
	Requests          *MetricValue `json:"requests,omitempty"`
	RequestsPerSecond *MetricValue `json:"requestsPerSecond,omitempty"`
	Errors            *MetricValue `json:"errors,omitempty"`
	ErrorRate         *MetricValue `json:"errorRate,omitempty"`
	DurationSeconds   *MetricValue `json:"durationSeconds,omitempty"`
}
