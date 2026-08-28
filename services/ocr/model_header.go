package ocr

type Header struct {
	Location []*TablePoint `json:"location,omitempty"`
	Words    *string       `json:"words,omitempty"`
}
