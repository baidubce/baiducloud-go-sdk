package face

type EyeStatus struct {
	LeftEye  *float64 `json:"left_eye,omitempty"`
	RightEye *float64 `json:"right_eye,omitempty"`
}
