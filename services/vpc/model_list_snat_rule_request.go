package vpc

type ListSnatRuleRequest struct {
	NatId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
