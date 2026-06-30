package iam

type AccessKey struct {
	Id          *string `json:"id,omitempty"`
	Secret      *string `json:"secret,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	Description *string `json:"description,omitempty"`
}
