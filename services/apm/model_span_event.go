package apm

type SpanEvent struct {
	Name       *string            `json:"name,omitempty"`
	Timestamp  *int64             `json:"timestamp,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
}
