package apm

type DescribeTraceRequest struct {
	Action        *string   `json:"-"`
	SpanDatetime  *string   `json:"spanDatetime,omitempty"`
	BeginDatetime *string   `json:"beginDatetime,omitempty"`
	EndDatetime   *string   `json:"endDatetime,omitempty"`
	TraceId       *string   `json:"traceId,omitempty"`
	Filters       []*Filter `json:"filters,omitempty"`
	ReturnHeight  *bool     `json:"returnHeight,omitempty"`
}
