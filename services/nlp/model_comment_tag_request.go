package nlp

type CommentTagRequest struct {
	Charset *string `json:"-"`
	Text    *string `json:"text,omitempty"`
	NlpType *int32  `json:"type,omitempty"`
}
