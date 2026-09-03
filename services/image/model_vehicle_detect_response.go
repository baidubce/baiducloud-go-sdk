package image

import "github.com/baidubce/baiducloud-go-sdk/bce"

type VehicleDetectResponse struct {
	bce.BaseResponse
	ErrorCode   *int32         `json:"error_code,omitempty"`
	ErrorMsg    *string        `json:"error_msg,omitempty"`
	LogId       *int64         `json:"log_id,omitempty"`
	VehicleNum  *VehicleNumber `json:"vehicle_num,omitempty"`
	VehicleInfo []*VehicleInfo `json:"vehicle_info,omitempty"`
}
