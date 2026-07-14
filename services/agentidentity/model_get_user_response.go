package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetUserResponse struct {
	bce.BaseResponse
	Id          *string `json:"id,omitempty"`
	Username    *string `json:"username,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Description *string `json:"description,omitempty"`
	Source      *string `json:"source,omitempty"`
	HasPassword *bool   `json:"hasPassword,omitempty"`
	CreatedAt   *string `json:"createdAt,omitempty"`
}
