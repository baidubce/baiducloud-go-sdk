package face

type FaceDetectOcclusion struct {
	Nose        *float64 `json:"nose,omitempty"`
	Mouth       *float64 `json:"mouth,omitempty"`
	LeftEye     *float64 `json:"left_eye,omitempty"`
	RightEye    *float64 `json:"right_eye,omitempty"`
	LeftCheek   *float64 `json:"left_cheek,omitempty"`
	RightCheek  *float64 `json:"right_cheek,omitempty"`
	ChinContour *float64 `json:"chin_contour,omitempty"`
}
