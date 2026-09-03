package image

type ImageUnderstandingRequestRequest struct {
	Image    *string `json:"image,omitempty"`
	Url      *string `json:"url,omitempty"`
	Question *string `json:"question,omitempty"`
}
