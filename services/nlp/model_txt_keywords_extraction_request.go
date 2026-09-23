package nlp

type TxtKeywordsExtractionRequest struct {
	Charset *string   `json:"-"`
	Text    []*string `json:"text,omitempty"`
	Num     *int32    `json:"num,omitempty"`
}
