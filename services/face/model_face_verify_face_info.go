package face

type FaceVerifyFaceInfo struct {
	FaceToken       *string                    `json:"face_token,omitempty"`
	Location        *FaceVerifyLocation        `json:"location,omitempty"`
	FaceProbability *float64                   `json:"face_probability,omitempty"`
	Angle           *FaceVerifyAngle           `json:"angle,omitempty"`
	Age             *float64                   `json:"age,omitempty"`
	Expression      *FaceVerifyTypeProbability `json:"expression,omitempty"`
	FaceShape       *FaceVerifyTypeProbability `json:"face_shape,omitempty"`
	Gender          *FaceVerifyTypeProbability `json:"gender,omitempty"`
	Glasses         *FaceVerifyTypeProbability `json:"glasses,omitempty"`
	FaceType        *FaceVerifyTypeProbability `json:"face_type,omitempty"`
	Landmark        []*FaceVerifyPoint         `json:"landmark,omitempty"`
	Landmark72      []*FaceVerifyPoint         `json:"landmark72,omitempty"`
	Quality         *FaceVerifyQuality         `json:"quality,omitempty"`
	Liveness        *FaceVerifyLiveness        `json:"liveness,omitempty"`
	Beauty          *float64                   `json:"beauty,omitempty"`
	Spoofing        *float64                   `json:"spoofing,omitempty"`
	NotSpoofing     *float64                   `json:"not_spoofing,omitempty"`
}
