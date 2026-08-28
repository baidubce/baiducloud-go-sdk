package ocr

type Line struct {
	LineId      *string                   `json:"lineId,omitempty"`
	Text        *string                   `json:"text,omitempty"`
	Bbox        *HandwritingGetBBox       `json:"bbox,omitempty"`
	ParagraphId *string                   `json:"paragraphId,omitempty"`
	Chars       []*HandwritingGetCharInfo `json:"chars,omitempty"`
}
