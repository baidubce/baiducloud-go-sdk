package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateDedicatedGatewayResponse struct {
	bce.BaseResponse
	EtGatewayId *string `json:"etGatewayId,omitempty"`
}
