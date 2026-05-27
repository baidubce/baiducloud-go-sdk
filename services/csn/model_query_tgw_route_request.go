package csn

type QueryTgwRouteRequest struct {
	CsnId   *string `json:"-"`
	TgwId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
