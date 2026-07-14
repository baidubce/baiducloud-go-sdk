package as

type CdsInfo struct {
	VolumeType   *string `json:"volumeType,omitempty"`
	SizeInGB     *int32  `json:"sizeInGB,omitempty"`
	SnapshotId   *string `json:"snapshotId,omitempty"`
	SnapshotName *string `json:"snapshotName,omitempty"`
}
