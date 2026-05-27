package csn

type QueryCsnInstanceRequest struct {
	CsnId   *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
