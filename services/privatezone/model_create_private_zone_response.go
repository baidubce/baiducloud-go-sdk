package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreatePrivateZoneResponse struct {
	bce.BaseResponse
	ZoneId *string `json:"zoneId,omitempty"`
}
