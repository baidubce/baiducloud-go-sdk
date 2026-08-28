package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ParserTaskResponse struct {
	bce.BaseResponse
	ErrorCode *int32        `json:"error_code,omitempty"`
	ErrorMsg  *string       `json:"error_msg,omitempty"`
	LogId     *string       `json:"log_id,omitempty"`
	Result    *ParserResult `json:"result,omitempty"`
}
