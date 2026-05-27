package csn

type QueryCsnListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
