package nlp

type TopicTagItem struct {
	Score *float32 `json:"score,omitempty"`
	Tag   *string  `json:"tag,omitempty"`
}
