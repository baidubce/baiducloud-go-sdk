package ax

type BatchReleaseSandboxesRequest struct {
	SandboxIds []*string `json:"sandboxIds,omitempty"`
}
