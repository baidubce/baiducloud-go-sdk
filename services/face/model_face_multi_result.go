package face

type FaceMultiResult struct {
	FaceNum  *int32               `json:"face_num,omitempty"`
	FaceList []*FaceMultiFaceInfo `json:"face_list,omitempty"`
}
