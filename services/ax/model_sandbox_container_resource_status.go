package ax

type SandboxContainerResourceStatus struct {
	Name      *string          `json:"name,omitempty"`
	Desired   *SandboxResource `json:"desired,omitempty"`
	Allocated *SandboxResource `json:"allocated,omitempty"`
	Current   *SandboxResource `json:"current,omitempty"`
}
