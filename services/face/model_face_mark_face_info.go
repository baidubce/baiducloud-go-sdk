package face

type FaceMarkFaceInfo struct {
	Location        *FaceMarkLocation `json:"location,omitempty"`
	Angle           *FaceMarkAngle    `json:"angle,omitempty"`
	Age             *float64          `json:"age,omitempty"`
	Gender          *Gender           `json:"gender,omitempty"`
	Landmark72      []*FaceMarkPoint  `json:"landmark72,omitempty"`
	Landmark150     *interface{}      `json:"landmark150,omitempty"`
	Landmark201     *interface{}      `json:"landmark201,omitempty"`
	FaceToken       *string           `json:"face_token,omitempty"`
	FaceProbability *float64          `json:"face_probability,omitempty"`
}
