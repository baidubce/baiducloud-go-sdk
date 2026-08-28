package ocr

type HandwritingCandidate struct {
	Word *string `json:"word,omitempty"`
	Prob *string `json:"prob,omitempty"`
}
