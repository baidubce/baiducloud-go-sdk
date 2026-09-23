package nlp

type TextCorrectionRequest struct {
	Charset *string `json:"-"`
	Text    *string `json:"text,omitempty"`
}
