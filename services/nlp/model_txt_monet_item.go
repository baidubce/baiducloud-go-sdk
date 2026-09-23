package nlp

type TxtMonetItem struct {
	Text        *string  `json:"text,omitempty"`
	Prob        *float32 `json:"prob,omitempty"`
	StartOffset *int32   `json:"start_offset,omitempty"`
	EndOffset   *int32   `json:"end_offset,omitempty"`
}
