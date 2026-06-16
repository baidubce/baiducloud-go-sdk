package bcc

type CreateSnapshotShareRequest struct {
	SnapshotId *string   `json:"snapshotId,omitempty"`
	AccountIds []*string `json:"accountIds,omitempty"`
}
