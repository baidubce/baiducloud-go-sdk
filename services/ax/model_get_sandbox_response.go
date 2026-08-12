package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSandboxResponse struct {
	bce.BaseResponse
	SandboxID       *string            `json:"sandboxID,omitempty"`
	State           *string            `json:"state,omitempty"`
	Metadata        *map[string]string `json:"metadata,omitempty"`
	TemplateID      *string            `json:"templateID,omitempty"`
	Alias           *string            `json:"alias,omitempty"`
	ClientID        *string            `json:"clientID,omitempty"`
	Domain          *string            `json:"domain,omitempty"`
	EnvdAccessToken *string            `json:"envdAccessToken,omitempty"`
	EnvdVersion     *string            `json:"envdVersion,omitempty"`
	CpuCount        *int32             `json:"cpuCount,omitempty"`
	MemoryMB        *int32             `json:"memoryMB,omitempty"`
	DiskSizeMB      *int32             `json:"diskSizeMB,omitempty"`
	StartedAt       *string            `json:"startedAt,omitempty"`
	EndAt           *string            `json:"endAt,omitempty"`
	VpcDomain       *string            `json:"vpcDomain,omitempty"`
}
