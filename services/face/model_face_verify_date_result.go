package face

type FaceVerifyDateResult struct {
	VerifyStatus *int32   `json:"verify_status,omitempty"`
	VerifyScore  *float32 `json:"verify_score,omitempty"`
}
