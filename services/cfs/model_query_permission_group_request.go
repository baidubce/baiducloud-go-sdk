package cfs

type QueryPermissionGroupRequest struct {
	AgName  *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
