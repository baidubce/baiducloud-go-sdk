package pfs

type ListPfsRequest struct {
	MaxKeys   *int32  `json:"-"`
	Manner    *string `json:"-"`
	Marker    *string `json:"-"`
	FilterTag *string `json:"-"`
}
