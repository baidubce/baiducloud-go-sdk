package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAPrivateZoneResponse struct {
	bce.BaseResponse
	ZoneId *string `json:"zoneId,omitempty"`
}
