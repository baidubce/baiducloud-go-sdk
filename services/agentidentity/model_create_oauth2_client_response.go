package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateOauth2ClientResponse struct {
	bce.BaseResponse
	Id           *string `json:"id,omitempty"`
	ClientId     *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	Name         *string `json:"name,omitempty"`
	ClientType   *string `json:"clientType,omitempty"`
	CreatedAt    *string `json:"createdAt,omitempty"`
}
