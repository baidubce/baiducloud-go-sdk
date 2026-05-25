package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAHighlyAvailableVirtualIpResponse struct {
	bce.BaseResponse
	HaVipId *string `json:"haVipId,omitempty"`
}
