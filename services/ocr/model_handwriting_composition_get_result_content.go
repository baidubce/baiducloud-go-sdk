package ocr

type HandwritingCompositionGetResultContent struct {
	Lines      [][]*Line    `json:"lines,omitempty"`
	Paragraphs []*Paragraph `json:"paragraphs,omitempty"`
}
