package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateHaVipResponse struct {
	bce.BaseResponse
	HaVipId *string `json:"haVipId,omitempty"`
}
