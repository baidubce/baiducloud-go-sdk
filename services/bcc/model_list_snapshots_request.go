package bcc

type ListSnapshotsRequest struct {
	Marker   *string `json:"-"`
	MaxKeys  *int32  `json:"-"`
	VolumeId *string `json:"-"`
}
