package vpc

type ListHaVipRequest struct {
	VpcId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
