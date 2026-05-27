package csn

type QueryRouteTableListRequest struct {
	CsnId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
