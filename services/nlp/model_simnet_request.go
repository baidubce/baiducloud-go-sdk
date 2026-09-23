package nlp

type SimnetRequest struct {
	Charset *string `json:"-"`
	Text1   *string `json:"text_1,omitempty"`
	Text2   *string `json:"text_2,omitempty"`
	Model   *string `json:"model,omitempty"`
}
