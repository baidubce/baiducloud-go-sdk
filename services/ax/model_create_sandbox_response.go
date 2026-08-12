package ax

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSandboxResponse struct {
	bce.BaseResponse
	SandboxID       *string `json:"sandboxID,omitempty"`
	TemplateID      *string `json:"templateID,omitempty"`
	EnvdAccessToken *string `json:"envdAccessToken,omitempty"`
	Domain          *string `json:"domain,omitempty"`
	Alias           *string `json:"alias,omitempty"`
	ClientID        *string `json:"clientID,omitempty"`
	EnvdVersion     *string `json:"envdVersion,omitempty"`
	VpcDomain       *string `json:"vpcDomain,omitempty"`
}
