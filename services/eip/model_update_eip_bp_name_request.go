package eip

type UpdateEipBpNameRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
}
