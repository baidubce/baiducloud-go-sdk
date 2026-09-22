package face

import "github.com/baidubce/baiducloud-go-sdk/bce"

type FaceGetListResponse struct {
	bce.BaseResponse
	Timestamp *int64             `json:"timestamp,omitempty"`
	Result    *FaceGetListResult `json:"result,omitempty"`
	ErrorCode *int32             `json:"error_code,omitempty"`
	ErrorMsg  *string            `json:"error_msg,omitempty"`
	LogId     *int64             `json:"log_id,omitempty"`
}
