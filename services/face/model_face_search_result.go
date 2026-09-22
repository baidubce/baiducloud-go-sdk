package face

type FaceSearchResult struct {
	FaceToken *string               `json:"face_token,omitempty"`
	UserList  []*FaceSearchUserInfo `json:"user_list,omitempty"`
}
