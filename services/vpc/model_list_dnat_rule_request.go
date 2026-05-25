package vpc

type ListDnatRuleRequest struct {
	NatId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
