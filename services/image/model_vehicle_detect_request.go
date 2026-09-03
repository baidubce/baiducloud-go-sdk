package image

type VehicleDetectRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
	Area  *string `json:"area,omitempty"`
}
