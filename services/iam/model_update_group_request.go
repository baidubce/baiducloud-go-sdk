package iam

type UpdateGroupRequest struct {
	GroupName   *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
