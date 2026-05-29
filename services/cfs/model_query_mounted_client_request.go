package cfs

type QueryMountedClientRequest struct {
	FsId    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
