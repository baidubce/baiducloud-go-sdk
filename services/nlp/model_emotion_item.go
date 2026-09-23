package nlp

type EmotionItem struct {
	Label    *string           `json:"label,omitempty"`
	Prob     *float64          `json:"prob,omitempty"`
	Subitems []*EmotionSubitem `json:"subitems,omitempty"`
	Replies  []*string         `json:"replies,omitempty"`
}
