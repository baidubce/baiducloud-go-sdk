package face

type FaceMarkResult struct {
	FaceNum  *int32              `json:"face_num,omitempty"`
	FaceList []*FaceMarkFaceInfo `json:"face_list,omitempty"`
}
