package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetResourceOauth2tokenResponse struct {
	bce.BaseResponse
	AccessToken      *string `json:"accessToken,omitempty"`
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
	SessionUri       *string `json:"sessionUri,omitempty"`
	SessionStatus    *string `json:"sessionStatus,omitempty"`
}
