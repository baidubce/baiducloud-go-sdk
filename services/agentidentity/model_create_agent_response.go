package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAgentResponse struct {
	bce.BaseResponse
	Id                              *string   `json:"id,omitempty"`
	BceDomainId                     *string   `json:"bceDomainId,omitempty"`
	BceUserId                       *string   `json:"bceUserId,omitempty"`
	Name                            *string   `json:"name,omitempty"`
	Description                     *string   `json:"description,omitempty"`
	Extra                           *string   `json:"extra,omitempty"`
	AllowedResourceOauth2ReturnUrls []*string `json:"allowedResourceOauth2ReturnUrls,omitempty"`
	CreatedAt                       *string   `json:"createdAt,omitempty"`
	UpdatedAt                       *string   `json:"updatedAt,omitempty"`
}
