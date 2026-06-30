package iam

type UserModel struct {
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	CreateTime    *string `json:"createTime,omitempty"`
	JoinTime      *string `json:"joinTime,omitempty"`
	LastLoginTime *string `json:"lastLoginTime,omitempty"`
	Description   *string `json:"description,omitempty"`
	Enabled       *bool   `json:"enabled,omitempty"`
}
