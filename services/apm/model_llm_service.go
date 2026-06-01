package apm

type LLMService struct {
	ServiceName          *string      `json:"serviceName,omitempty"`
	ServiceDisplayName   *string      `json:"serviceDisplayName,omitempty"`
	ServiceId            *string      `json:"serviceId,omitempty"`
	Tags                 []*Tag       `json:"tags,omitempty"`
	LlmRequests          *MetricValue `json:"llmRequests,omitempty"`
	LlmRequestsPerSecond *MetricValue `json:"llmRequestsPerSecond,omitempty"`
	LlmErrors            *MetricValue `json:"llmErrors,omitempty"`
	LlmDurationSeconds   *MetricValue `json:"llmDurationSeconds,omitempty"`
	LlmTokens            *MetricValue `json:"llmTokens,omitempty"`
}
