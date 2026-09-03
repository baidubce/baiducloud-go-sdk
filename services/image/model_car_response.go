package image

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CarResponse struct {
	bce.BaseResponse
	ErrorCode      *int32             `json:"error_code,omitempty"`
	ErrorMsg       *string            `json:"error_msg,omitempty"`
	LogId          *int64             `json:"log_id,omitempty"`
	ColorResult    *string            `json:"color_result,omitempty"`
	Result         []*CarResult       `json:"result,omitempty"`
	Brand          *string            `json:"brand,omitempty"`
	LocationResult *CarLocationResult `json:"location_result,omitempty"`
}
