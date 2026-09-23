package nlp

type TopicRequest struct {
	Charset *string `json:"-"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}
