package nlp

type EmotionSubitem struct {
	Label *string  `json:"label,omitempty"`
	Prob  *float64 `json:"prob,omitempty"`
}
