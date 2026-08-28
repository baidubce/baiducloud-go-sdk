package ocr

type SealRecogResult struct {
	Location    *DocAnaysisOfficeLocation `json:"location,omitempty"`
	Probability *float64                  `json:"probability,omitempty"`
	OcrType     *string                   `json:"type,omitempty"`
	Major       *SealField                `json:"major,omitempty"`
	Minor       []*SealField              `json:"minor,omitempty"`
}
