package vpc

type NatBindEipRequest struct {
	NatId       *string   `json:"-"`
	ClientToken *string   `json:"-"`
	BindEips    []*string `json:"bindEips,omitempty"`
}
