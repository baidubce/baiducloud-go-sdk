package ocr

type GeneralResponseChar struct {
	Char     *string          `json:"char,omitempty"`
	Location *GeneralLocation `json:"location,omitempty"`
}
