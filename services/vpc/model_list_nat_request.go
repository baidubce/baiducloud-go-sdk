package vpc

type ListNatRequest struct {
	VpcId   *string `json:"-"`
	NatId   *string `json:"-"`
	Name    *string `json:"-"`
	Ip      *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
