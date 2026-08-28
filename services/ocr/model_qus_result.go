package ocr

type QusResult struct {
	QusType        *string       `json:"qus_type,omitempty"`
	QusProbability *float64      `json:"qus_probability,omitempty"`
	ElemText       *ElemText     `json:"elem_text,omitempty"`
	QusLocation    *QusLocation  `json:"qus_location,omitempty"`
	QusElement     []*QusElement `json:"qus_element,omitempty"`
}
