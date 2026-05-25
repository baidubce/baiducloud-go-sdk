package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAnIpv6GatewayResponse struct {
	bce.BaseResponse
	GatewayId *string `json:"gatewayId,omitempty"`
}
