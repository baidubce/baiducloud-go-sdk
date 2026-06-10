package bcc

type RollbackVolumeRequest struct {
	VolumeId   *string `json:"-"`
	SnapshotId *string `json:"snapshotId,omitempty"`
}
