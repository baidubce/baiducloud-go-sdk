package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateIpv6GatewayResponse struct {
	bce.BaseResponse
	GatewayId *string `json:"gatewayId,omitempty"`
}
