package iam

type PolicyModel struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	IamType     *string `json:"type,omitempty"`
	CreateTime  *string `json:"createTime,omitempty"`
	AttachTime  *string `json:"attachTime,omitempty"`
	Description *string `json:"description,omitempty"`
	Document    *string `json:"document,omitempty"`
}
