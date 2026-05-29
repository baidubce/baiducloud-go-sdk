package cfs

type QueryMountingTargetRequest struct {
	FId     *string `json:"-"`
	MountId *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
