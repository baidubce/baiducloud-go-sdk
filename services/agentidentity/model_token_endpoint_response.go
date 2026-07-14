package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TokenEndpointResponse struct {
	bce.BaseResponse
	AccessToken      *string `json:"access_token,omitempty"`
	IdToken          *string `json:"id_token,omitempty"`
	RefreshToken     *string `json:"refresh_token,omitempty"`
	TokenType        *string `json:"token_type,omitempty"`
	ExpiresIn        *int32  `json:"expires_in,omitempty"`
	RefreshExpiresIn *int32  `json:"refresh_expires_in,omitempty"`
}
