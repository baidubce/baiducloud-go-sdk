package face

type FaceVerifyResult struct {
	FaceLiveness *float32              `json:"face_liveness,omitempty"`
	Thresholds   *Thresholds           `json:"thresholds,omitempty"`
	FaceList     []*FaceVerifyFaceInfo `json:"face_list,omitempty"`
}
