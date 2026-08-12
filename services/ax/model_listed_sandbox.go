package ax

type ListedSandbox struct {
	SandboxID   *string            `json:"sandboxID,omitempty"`
	State       *string            `json:"state,omitempty"`
	Metadata    *map[string]string `json:"metadata,omitempty"`
	TemplateID  *string            `json:"templateID,omitempty"`
	CpuCount    *int32             `json:"cpuCount,omitempty"`
	MemoryMB    *int32             `json:"memoryMB,omitempty"`
	DiskSizeMB  *int32             `json:"diskSizeMB,omitempty"`
	StartedAt   *string            `json:"startedAt,omitempty"`
	EndAt       *string            `json:"endAt,omitempty"`
	EnvdVersion *string            `json:"envdVersion,omitempty"`
	ClientID    *string            `json:"clientID,omitempty"`
	Alias       *string            `json:"alias,omitempty"`
}
