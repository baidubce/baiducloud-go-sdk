package face

type FaceMergeRequest struct {
	Version       *string  `json:"version,omitempty"`
	Alpha         *float32 `json:"alpha,omitempty"`
	ImageTemplate *string  `json:"image_template,omitempty"`
	ImageTarget   *string  `json:"image_target,omitempty"`
	MergeDegree   *string  `json:"merge_degree,omitempty"`
	Position      *int32   `json:"position,omitempty"`
	Language      *int32   `json:"language,omitempty"`
}
