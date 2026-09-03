package image

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AiRetouchingQueryTaskResponse struct {
	bce.BaseResponse
	ErrorCode *int32            `json:"error_code,omitempty"`
	ErrorMsg  *string           `json:"error_msg,omitempty"`
	LogId     *int64            `json:"log_id,omitempty"`
	RequestId *int64            `json:"request_id,omitempty"`
	Result    *RetouchingResult `json:"result,omitempty"`
}
