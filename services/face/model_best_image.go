package face

type BestImage struct {
	FaceToken     *string  `json:"face_token,omitempty"`
	Pic           *string  `json:"pic,omitempty"`
	LivenessScore *float32 `json:"liveness_score,omitempty"`
}
