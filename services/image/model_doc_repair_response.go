package image

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DocRepairResponse struct {
	bce.BaseResponse
	LogId     *int64  `json:"log_id,omitempty"`
	ErrorCode *int32  `json:"error_code,omitempty"`
	ErrorMsg  *string `json:"error_msg,omitempty"`
	Result    *string `json:"result,omitempty"`
}
