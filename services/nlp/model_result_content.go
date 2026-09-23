package nlp

type ResultContent struct {
	Content *string           `json:"content,omitempty"`
	Results []*TxtMonetResult `json:"results,omitempty"`
}
