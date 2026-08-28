package ocr

type TableContent struct {
	PolyLocation []*TablePoint `json:"poly_location,omitempty"`
	Word         *string       `json:"word,omitempty"`
}
