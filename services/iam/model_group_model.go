package iam

type GroupModel struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	Description *string `json:"description,omitempty"`
}
