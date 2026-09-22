package face

type FaceRegResult struct {
	Location  *FaceLocation `json:"location,omitempty"`
	FaceToken *string       `json:"face_token,omitempty"`
}
