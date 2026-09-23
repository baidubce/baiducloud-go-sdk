package nlp

type EntityAnalysis struct {
	Mention    *string   `json:"mention,omitempty"`
	Category   *Category `json:"category,omitempty"`
	Confidence *float64  `json:"confidence,omitempty"`
	Desc       *string   `json:"desc,omitempty"`
	Status     *string   `json:"status,omitempty"`
}
