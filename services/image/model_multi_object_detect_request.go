package image

type MultiObjectDetectRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
}
