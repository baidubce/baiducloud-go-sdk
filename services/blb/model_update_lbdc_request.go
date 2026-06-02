package blb

type UpdateLbdcRequest struct {
	Id          *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Desc        *string `json:"desc,omitempty"`
}
