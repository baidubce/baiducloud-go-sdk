package nlp

type AddressRequest struct {
	Charset *string `json:"-"`
	Text    *string `json:"text,omitempty"`
}
