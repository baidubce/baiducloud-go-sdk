package face

type FaceDetectFaceInfo struct {
	Location        *FaceDetectLocation        `json:"location,omitempty"`
	Angle           *FaceDetectAngle           `json:"angle,omitempty"`
	Age             *float64                   `json:"age,omitempty"`
	Expression      *FaceDetectTypeProbability `json:"expression,omitempty"`
	Gender          *FaceDetectTypeProbability `json:"gender,omitempty"`
	Glasses         *FaceDetectTypeProbability `json:"glasses,omitempty"`
	Emotion         *FaceDetectTypeProbability `json:"emotion,omitempty"`
	Mask            *MaskInfo                  `json:"mask,omitempty"`
	Landmark        []*FaceDetectPoint         `json:"landmark,omitempty"`
	Landmark72      []*FaceDetectPoint         `json:"landmark72,omitempty"`
	Landmark150     *interface{}               `json:"landmark150,omitempty"`
	Quality         *FaceDetectQuality         `json:"quality,omitempty"`
	Liveness        *FaceDetectLiveness        `json:"liveness,omitempty"`
	Spoofing        *float64                   `json:"spoofing,omitempty"`
	FaceToken       *string                    `json:"face_token,omitempty"`
	FaceProbability *float64                   `json:"face_probability,omitempty"`
	FaceShape       *FaceDetectTypeProbability `json:"face_shape,omitempty"`
	EyeStatus       *EyeStatus                 `json:"eye_status,omitempty"`
	FaceType        *FaceDetectTypeProbability `json:"face_type,omitempty"`
	NotSpoofing     *float64                   `json:"not_spoofing,omitempty"`
}
