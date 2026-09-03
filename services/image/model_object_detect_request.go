package image

type ObjectDetectRequest struct {
	Image    *string `json:"image,omitempty"`
	WithFace *int32  `json:"with_face,omitempty"`
}
