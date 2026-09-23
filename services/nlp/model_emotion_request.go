package nlp

type EmotionRequest struct {
	Charset *string `json:"-"`
	Text    *string `json:"text,omitempty"`
	Scene   *string `json:"scene,omitempty"`
}
