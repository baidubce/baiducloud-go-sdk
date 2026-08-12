package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ConnectSandboxResponse struct {
	bce.BaseResponse
	SandboxID       *string `json:"sandboxID,omitempty"`
	TemplateID      *string `json:"templateID,omitempty"`
	EnvdAccessToken *string `json:"envdAccessToken,omitempty"`
	Domain          *string `json:"domain,omitempty"`
}
