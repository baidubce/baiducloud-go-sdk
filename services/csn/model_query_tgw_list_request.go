package csn

type QueryTgwListRequest struct {
	CsnId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
