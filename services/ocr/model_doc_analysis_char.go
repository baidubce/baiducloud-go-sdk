package ocr

type DocAnalysisChar struct {
	Char          *string        `json:"char,omitempty"`
	CharsLocation *CharsLocation `json:"chars_location,omitempty"`
}
