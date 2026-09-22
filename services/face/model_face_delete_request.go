package face

type FaceDeleteRequest struct {
	UserId    *string `json:"user_id,omitempty"`
	GroupId   *string `json:"group_id,omitempty"`
	FaceToken *string `json:"face_token,omitempty"`
	LogId     *int64  `json:"log_id,omitempty"`
}
