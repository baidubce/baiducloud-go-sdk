package nlp

type EcnetRequest struct {
	Charset *string `json:"-"`
	Text    *string `json:"text,omitempty"`
}
