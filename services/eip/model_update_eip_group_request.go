package eip

type UpdateEipGroupRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
}
