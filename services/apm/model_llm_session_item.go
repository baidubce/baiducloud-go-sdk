package apm

type LLMSessionItem struct {
	SessionId    *string `json:"sessionId,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	EndTime      *string `json:"endTime,omitempty"`
	Duration     *int64  `json:"duration,omitempty"`
	UserId       *string `json:"userId,omitempty"`
	TraceCount   *int32  `json:"traceCount,omitempty"`
	InputTokens  *int32  `json:"inputTokens,omitempty"`
	OutputTokens *int32  `json:"outputTokens,omitempty"`
	TotalTokens  *int32  `json:"totalTokens,omitempty"`
}
