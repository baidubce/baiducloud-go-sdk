package bcc

type CreateSnapshotRequest struct {
	VolumeId        *string     `json:"volumeId,omitempty"`
	SnapshotName    *string     `json:"snapshotName,omitempty"`
	Desc            *string     `json:"desc,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
	RetentionInDays *int32      `json:"retentionInDays,omitempty"`
}
