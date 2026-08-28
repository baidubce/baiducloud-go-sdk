package ocr

type Title struct {
	Bbox  []*HandwritingGetBBox     `json:"bbox,omitempty"`
	Text  *string                   `json:"text,omitempty"`
	Chars []*HandwritingGetCharInfo `json:"chars,omitempty"`
}
