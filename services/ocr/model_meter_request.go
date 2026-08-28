package ocr

type MeterRequest struct {
	Image        *string `json:"image,omitempty"`
	Url          *string `json:"url,omitempty"`
	Probability  *bool   `json:"probability,omitempty"`
	PolyLocation *bool   `json:"poly_location,omitempty"`
}
