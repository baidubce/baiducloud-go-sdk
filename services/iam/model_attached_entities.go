package iam

type AttachedEntities struct {
	Id         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	IamType    *string `json:"type,omitempty"`
	AttachTime *string `json:"attachTime,omitempty"`
}
