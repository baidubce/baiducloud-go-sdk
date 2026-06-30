package bci

type ContainerSecurityContext struct {
	Capabilities           *Capabilities `json:"capabilities,omitempty"`
	RunAsUser              *int64        `json:"runAsUser,omitempty"`
	RunAsGroup             *int64        `json:"runAsGroup,omitempty"`
	RunAsNonRoot           *bool         `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem *bool         `json:"readOnlyRootFilesystem,omitempty"`
}
