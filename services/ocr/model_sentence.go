package ocr

type Sentence struct {
	Bbox         []*HandwritingGetBBox `json:"bbox,omitempty"`
	SentenceId   *string               `json:"sentenceId,omitempty"`
	Text         *string               `json:"text,omitempty"`
	LineSegments []*LineSegment        `json:"lineSegments,omitempty"`
}
