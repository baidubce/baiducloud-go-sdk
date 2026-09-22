package face

type FaceVerifyOcclusion struct {
	LeftEye    *float64 `json:"left_eye,omitempty"`
	RightEye   *float64 `json:"right_eye,omitempty"`
	Nose       *float64 `json:"nose,omitempty"`
	Mouth      *float64 `json:"mouth,omitempty"`
	LeftCheek  *float64 `json:"left_cheek,omitempty"`
	RightCheek *float64 `json:"right_cheek,omitempty"`
	Chin       *float64 `json:"chin,omitempty"`
}
