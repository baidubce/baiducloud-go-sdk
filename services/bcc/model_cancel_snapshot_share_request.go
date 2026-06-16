package bcc

type CancelSnapshotShareRequest struct {
	SourceSnapshotId *string   `json:"sourceSnapshotId,omitempty"`
	AccountIds       []*string `json:"accountIds,omitempty"`
	ShareSnapshotId  *string   `json:"shareSnapshotId,omitempty"`
}
