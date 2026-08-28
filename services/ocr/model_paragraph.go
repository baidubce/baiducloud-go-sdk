package ocr

type Paragraph struct {
	Bbox        []*HandwritingGetBBox `json:"bbox,omitempty"`
	ParagraphId *string               `json:"paragraphId,omitempty"`
	IsColumn    *int32                `json:"isColumn,omitempty"`
	Text        *string               `json:"text,omitempty"`
	Sentences   []*Sentence           `json:"sentences,omitempty"`
}
