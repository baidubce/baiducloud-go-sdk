package ocr

type SealResult struct {
	Color       *string         `json:"color,omitempty"`
	SealImage   *string         `json:"seal_image,omitempty"`
	Location    *SealLocation   `json:"location,omitempty"`
	Probability *float64        `json:"probability,omitempty"`
	OcrType     *string         `json:"type,omitempty"`
	Major       *SealRegField   `json:"major,omitempty"`
	Minor       []*SealRegField `json:"minor,omitempty"`
}
