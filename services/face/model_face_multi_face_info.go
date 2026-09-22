package face

type FaceMultiFaceInfo struct {
	Location  *FaceMultiLocation   `json:"location,omitempty"`
	FaceToken *string              `json:"face_token,omitempty"`
	UserList  []*FaceMultiUserInfo `json:"user_list,omitempty"`
}
