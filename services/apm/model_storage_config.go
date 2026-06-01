package apm

type StorageConfig struct {
	TraceRetentionDays  *int32 `json:"traceRetentionDays,omitempty"`
	MetricRetentionDays *int32 `json:"metricRetentionDays,omitempty"`
	DorisRetentionDays  *int32 `json:"dorisRetentionDays,omitempty"`
}
