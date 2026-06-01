package apm

type LLMTraceItem struct {
	TraceId      *string   `json:"traceId,omitempty"`
	StartTime    *string   `json:"startTime,omitempty"`
	EndTime      *string   `json:"endTime,omitempty"`
	Service      *string   `json:"service,omitempty"`
	SessionId    *string   `json:"sessionId,omitempty"`
	UserId       *string   `json:"userId,omitempty"`
	InputTokens  *int32    `json:"inputTokens,omitempty"`
	OutputTokens *int32    `json:"outputTokens,omitempty"`
	TotalTokens  *int32    `json:"totalTokens,omitempty"`
	Input        *string   `json:"input,omitempty"`
	Output       *string   `json:"output,omitempty"`
	Models       []*string `json:"models,omitempty"`
	Tools        []*string `json:"tools,omitempty"`
	Duration     *int64    `json:"duration,omitempty"`
}
