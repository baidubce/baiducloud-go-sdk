package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QrcodeResponse struct {
	bce.BaseResponse
	ErrorCode      *int32         `json:"error_code,omitempty"`
	ErrorMsg       *string        `json:"error_msg,omitempty"`
	LogId          *int64         `json:"log_id,omitempty"`
	CodesResultNum *int32         `json:"codes_result_num,omitempty"`
	CodesResult    []*CodesResult `json:"codes_result,omitempty"`
}
