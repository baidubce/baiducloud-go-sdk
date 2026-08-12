package ax

type ReleaseResult struct {
	SandboxId *string `json:"sandboxId,omitempty"`
	Success   *bool   `json:"success,omitempty"`
	AxError   *string `json:"error,omitempty"`
}
