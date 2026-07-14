package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetWorkloadAccessTokenResponse struct {
	bce.BaseResponse
	Token    *string `json:"token,omitempty"`
	ExpireAt *string `json:"expireAt,omitempty"`
}
