package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchGetResourceApiKeyResponse struct {
	bce.BaseResponse
	Credentials *map[string]*ResourceCredentialDTO `json:"credentials,omitempty"`
}
