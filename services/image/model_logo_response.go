package image

import "github.com/baidubce/baiducloud-go-sdk/bce"

type LogoResponse struct {
	bce.BaseResponse
	ErrorCode *int32        `json:"error_code,omitempty"`
	ErrorMsg  *string       `json:"error_msg,omitempty"`
	LogId     *int64        `json:"log_id,omitempty"`
	ResultNum *int32        `json:"result_num,omitempty"`
	Result    []*LogoResult `json:"result,omitempty"`
}
