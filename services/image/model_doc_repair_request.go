package image

type DocRepairRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
}
