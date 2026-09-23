package nlp

type NewsSummaryRequest struct {
	Charset       *string `json:"-"`
	Title         *string `json:"title,omitempty"`
	Content       *string `json:"content,omitempty"`
	MaxSummaryLen *int32  `json:"max_summary_len,omitempty"`
}
