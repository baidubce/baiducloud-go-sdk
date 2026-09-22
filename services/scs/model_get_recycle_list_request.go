package scs

type GetRecycleListRequest struct {
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
