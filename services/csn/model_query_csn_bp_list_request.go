package csn

type QueryCsnBpListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
