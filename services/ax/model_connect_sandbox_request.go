package ax

type ConnectSandboxRequest struct {
	SandboxID  *string `json:"-"`
	Timeout    *int32  `json:"timeout,omitempty"`
	SnapshotID *string `json:"snapshotID,omitempty"`
}
