package ocr

type HandwritingChar struct {
	Char       *string                 `json:"char,omitempty"`
	Candidates []*HandwritingCandidate `json:"candidates,omitempty"`
	Location   *HandWritingLocation    `json:"location,omitempty"`
}
