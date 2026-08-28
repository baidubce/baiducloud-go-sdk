package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BirthCertificateResponse struct {
	bce.BaseResponse
	ErrorCode      *int32                  `json:"error_code,omitempty"`
	ErrorMsg       *string                 `json:"error_msg,omitempty"`
	LogId          *int64                  `json:"log_id,omitempty"`
	Direction      *int32                  `json:"direction,omitempty"`
	WordsResultNum *int32                  `json:"words_result_num,omitempty"`
	WordsResult    *BirthCertificateResult `json:"words_result,omitempty"`
}
