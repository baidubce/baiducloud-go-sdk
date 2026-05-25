package vpc

type ResizeNatRequest struct {
	NatId       *string `json:"-"`
	ClientToken *string `json:"-"`
	CuNum       *int32  `json:"cuNum,omitempty"`
}
