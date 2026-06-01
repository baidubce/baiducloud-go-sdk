package apm

type SessionTrace struct {
	StartTime    *string   `json:"startTime,omitempty"`
	EndTime      *string   `json:"endTime,omitempty"`
	Duration     *int32    `json:"duration,omitempty"`
	UserId       *string   `json:"userId,omitempty"`
	TraceId      *string   `json:"traceId,omitempty"`
	InputTokens  *int32    `json:"inputTokens,omitempty"`
	OutputTokens *int32    `json:"outputTokens,omitempty"`
	TotalTokens  *int32    `json:"totalTokens,omitempty"`
	Input        *string   `json:"input,omitempty"`
	Output       *string   `json:"output,omitempty"`
	Models       []*string `json:"models,omitempty"`
	Tools        []*string `json:"tools,omitempty"`
}
