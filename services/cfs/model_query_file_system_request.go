package cfs

type QueryFileSystemRequest struct {
	UserId    *string `json:"-"`
	FsId      *string `json:"-"`
	Marker    *string `json:"-"`
	MaxKeys   *int32  `json:"-"`
	FilterTag *string `json:"-"`
}
