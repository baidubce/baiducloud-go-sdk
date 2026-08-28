package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type FerryTicketResponse struct {
	bce.BaseResponse
	LogId          *int64                  `json:"log_id,omitempty"`
	ErrorCode      *int32                  `json:"error_code,omitempty"`
	ErrorMsg       *string                 `json:"error_msg,omitempty"`
	WordsResultNum *int32                  `json:"words_result_num,omitempty"`
	WordsResult    *FerryTicketWordsResult `json:"words_result,omitempty"`
}
