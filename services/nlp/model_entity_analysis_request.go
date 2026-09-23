package nlp

type EntityAnalysisRequest struct {
	Text    *string `json:"text,omitempty"`
	Mention *string `json:"mention,omitempty"`
}
