package vpc

type QueryVpnListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
	VpcId   *string `json:"-"`
	Eip     *string `json:"-"`
	Type    *string `json:"-"`
}
