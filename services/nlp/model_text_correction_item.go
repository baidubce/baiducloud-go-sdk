package nlp

type TextCorrectionItem struct {
	Text         *string                 `json:"text,omitempty"`
	CorrectQuery *string                 `json:"correct_query,omitempty"`
	ContentLen   *int32                  `json:"content_len,omitempty"`
	Details      []*TextCorrectionDetail `json:"details,omitempty"`
	ErrorNum     *int32                  `json:"error_num,omitempty"`
}
