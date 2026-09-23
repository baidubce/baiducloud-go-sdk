package nlp

type SentimentClassifyRequest struct {
	Charset *string `json:"-"`
	Text    *string `json:"text,omitempty"`
}
