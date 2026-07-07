package agentidentity

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UserinfoEndpointResponse struct {
	bce.BaseResponse
	Sub         *string `json:"sub,omitempty"`
	Username    *string `json:"username,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}
