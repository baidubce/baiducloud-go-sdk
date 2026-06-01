package apm

type SpanMessage struct {
	Role       *string   `json:"role,omitempty"`
	Content    *string   `json:"content,omitempty"`
	ContentRef []*string `json:"contentRef,omitempty"`
}
