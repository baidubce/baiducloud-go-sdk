package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type MultiIdcardResponse struct {
	bce.BaseResponse
	ErrorCode   *int32             `json:"error_code,omitempty"`
	ErrorMsg    *string            `json:"error_msg,omitempty"`
	LogId       *int64             `json:"log_id,omitempty"`
	WordsResult []*MultiIdcardItem `json:"words_result,omitempty"`
}
