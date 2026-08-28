package ocr

type SealRegField struct {
	Words        *string  `json:"words,omitempty"`
	FlattenImage *string  `json:"flatten_image,omitempty"`
	Probability  *float64 `json:"probability,omitempty"`
}
