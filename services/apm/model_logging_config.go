package apm

type LoggingConfig struct {
	Enabled      *bool   `json:"enabled,omitempty"`
	Region       *string `json:"region,omitempty"`
	Project      *string `json:"project,omitempty"`
	LogStoreName *string `json:"logStoreName,omitempty"`
	TraceIdIndex *string `json:"traceIdIndex,omitempty"`
	TraceIdKey   *string `json:"traceIdKey,omitempty"`
	SpanIdIndex  *string `json:"spanIdIndex,omitempty"`
	SpanIdKey    *string `json:"spanIdKey,omitempty"`
}
