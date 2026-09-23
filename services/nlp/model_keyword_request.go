package nlp

type KeywordRequest struct {
	Charset *string `json:"-"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}
