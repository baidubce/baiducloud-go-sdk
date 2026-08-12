package ax

type SetSandboxTimeoutRequest struct {
	SandboxID *string `json:"-"`
	Timeout   *int32  `json:"timeout,omitempty"`
}
