package vpc

type ModifyNatRequest struct {
	NatId       *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
}
