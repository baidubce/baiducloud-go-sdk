package ocr

type DocAnalysisOfficeChar struct {
	Char          *string                   `json:"char,omitempty"`
	CharProb      *float64                  `json:"char_prob,omitempty"`
	CharsLocation *DocAnaysisOfficeLocation `json:"chars_location,omitempty"`
}
