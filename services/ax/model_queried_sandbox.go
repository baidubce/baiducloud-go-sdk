package ax

type QueriedSandbox struct {
	SandboxID       *string            `json:"sandboxID,omitempty"`
	State           *string            `json:"state,omitempty"`
	Metadata        *map[string]string `json:"metadata,omitempty"`
	CpuCount        *int32             `json:"cpuCount,omitempty"`
	MemoryMB        *int32             `json:"memoryMB,omitempty"`
	StartedAt       *string            `json:"startedAt,omitempty"`
	EndAt           *string            `json:"endAt,omitempty"`
	EnvdVersion     *string            `json:"envdVersion,omitempty"`
	EnvdAccessToken *string            `json:"envdAccessToken,omitempty"`
	TemplateID      *string            `json:"templateID,omitempty"`
	ImagePath       *string            `json:"imagePath,omitempty"`
}
