package apm

type DescribeLLMTraceRequest struct {
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	TraceId       *string   `json:"traceId,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
	ReturnHeight  *bool     `json:"returnHeight,omitempty"`
}
