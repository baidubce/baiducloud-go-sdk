package face

type FaceDetectResult struct {
	FaceNum         *int32                `json:"face_num,omitempty"`
	FaceList        []*FaceDetectFaceInfo `json:"face_list,omitempty"`
	CorpImageBase64 *string               `json:"corp_image_base64,omitempty"`
}
