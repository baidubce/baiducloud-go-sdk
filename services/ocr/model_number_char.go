package ocr

type NumberChar struct {
	Char     *string          `json:"char,omitempty"`
	Location *NumbersLocation `json:"location,omitempty"`
}
