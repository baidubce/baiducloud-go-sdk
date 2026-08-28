package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DocClassifyResponse struct {
	bce.BaseResponse
	LogId       *int64                    `json:"log_id,omitempty"`
	ErrorCode   *int32                    `json:"error_code,omitempty"`
	ErrorMsg    *string                   `json:"error_msg,omitempty"`
	WordsResult []*DocClassifyWordsResult `json:"words_result,omitempty"`
}
