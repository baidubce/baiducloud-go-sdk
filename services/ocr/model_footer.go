package ocr

type Footer struct {
	Location []*TablePoint `json:"location,omitempty"`
	Words    *string       `json:"words,omitempty"`
}
