package ax

type CreateSandboxSnapshotRequest struct {
	SandboxID *string `json:"-"`
	Name      *string `json:"name,omitempty"`
}
