package apm

type SpanContentRef struct {
	Content    *string   `json:"content,omitempty"`
	ContentRef []*string `json:"contentRef,omitempty"`
}
