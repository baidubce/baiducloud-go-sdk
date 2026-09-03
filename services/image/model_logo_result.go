package image

type LogoResult struct {
	Location    *LogoLocation `json:"location,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Probability *float32      `json:"probability,omitempty"`
	ImageType   *int32        `json:"type,omitempty"`
}
