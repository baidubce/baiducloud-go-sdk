package face

type UserCopyRequest struct {
	UserId     *string `json:"user_id,omitempty"`
	SrcGroupId *string `json:"src_group_id,omitempty"`
	DstGroupId *string `json:"dst_group_id,omitempty"`
}
