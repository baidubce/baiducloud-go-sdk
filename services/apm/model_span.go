package apm

type Span struct {
	TraceId      *string            `json:"traceId,omitempty"`
	SpanId       *string            `json:"spanId,omitempty"`
	ParentSpanId *string            `json:"parentSpanId,omitempty"`
	Name         *string            `json:"name,omitempty"`
	Service      *string            `json:"service,omitempty"`
	Start        *int64             `json:"start,omitempty"`
	End          *int64             `json:"end,omitempty"`
	Duration     *int64             `json:"duration,omitempty"`
	Host         *string            `json:"host,omitempty"`
	StatusCode   *string            `json:"statusCode,omitempty"`
	Kind         *string            `json:"kind,omitempty"`
	Attributes   *map[string]string `json:"attributes,omitempty"`
	Resource     *map[string]string `json:"resource,omitempty"`
	Events       []*SpanEvent       `json:"events,omitempty"`
	Height       *int32             `json:"height,omitempty"`
	SubSpans     []*Span            `json:"subSpans,omitempty"`
}
