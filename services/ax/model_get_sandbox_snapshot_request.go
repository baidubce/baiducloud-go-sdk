package ax

type GetSandboxSnapshotRequest struct {
	SandboxID  *string `json:"-"`
	SnapshotID *string `json:"-"`
}
