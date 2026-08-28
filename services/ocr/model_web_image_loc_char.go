package ocr

type WebImageLocChar struct {
	Char     *string              `json:"char,omitempty"`
	Location *WebImageLocLocation `json:"location,omitempty"`
}
