package nlp

type EcnetItem struct {
	CorrectQuery *string        `json:"correct_query,omitempty"`
	Score        *float64       `json:"score,omitempty"`
	VecFragment  []*VecFragment `json:"vec_fragment,omitempty"`
}
