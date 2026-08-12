package ax

type ResumeSandboxRequest struct {
	SandboxID *string `json:"-"`
	Timeout   *int32  `json:"timeout,omitempty"`
	AutoPause *bool   `json:"autoPause,omitempty"`
}
