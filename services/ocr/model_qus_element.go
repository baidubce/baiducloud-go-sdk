package ocr

type QusElement struct {
	ElemType        *string       `json:"elem_type,omitempty"`
	ElemProbability *float64      `json:"elem_probability,omitempty"`
	ElemLocation    *ElemLocation `json:"elem_location,omitempty"`
	ElemWord        []*ElemWord   `json:"elem_word,omitempty"`
}
