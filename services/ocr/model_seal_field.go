package ocr

type SealField struct {
	Words       *string  `json:"words,omitempty"`
	Probability *float64 `json:"probability,omitempty"`
}
