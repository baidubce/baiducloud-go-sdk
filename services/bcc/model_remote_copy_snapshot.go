package bcc

type RemoteCopySnapshot struct {
	Region     *string `json:"region,omitempty"`
	SnapshotId *string `json:"snapshotId,omitempty"`
}
