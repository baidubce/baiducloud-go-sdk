package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteAIGatewayResponse struct {
	bce.BaseResponse
	Result *string `json:"result,omitempty"`
}
