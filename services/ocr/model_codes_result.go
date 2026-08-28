package ocr

type CodesResult struct {
	OcrType  *string         `json:"type,omitempty"`
	Text     []*string       `json:"text,omitempty"`
	Location *QRCodeLocation `json:"location,omitempty"`
}
