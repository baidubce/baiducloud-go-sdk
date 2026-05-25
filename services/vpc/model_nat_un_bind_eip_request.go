package vpc

type NatUnBindEipRequest struct {
	NatId       *string   `json:"-"`
	ClientToken *string   `json:"-"`
	BindEips    []*string `json:"bindEips,omitempty"`
}
