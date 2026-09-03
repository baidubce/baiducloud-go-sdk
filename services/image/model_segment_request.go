package image

type SegmentRequest struct {
	Image      *string `json:"image,omitempty"`
	Url        *string `json:"url,omitempty"`
	Method     *string `json:"method,omitempty"`
	ReturnForm *string `json:"return_form,omitempty"`
	RefineMask *bool   `json:"refine_mask,omitempty"`
	Position   *string `json:"position,omitempty"`
}
