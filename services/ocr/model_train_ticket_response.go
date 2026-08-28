package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TrainTicketResponse struct {
	bce.BaseResponse
	LogId          *int64                  `json:"log_id,omitempty"`
	ErrorCode      *int32                  `json:"error_code,omitempty"`
	ErrorMsg       *string                 `json:"error_msg,omitempty"`
	Direction      *int32                  `json:"direction,omitempty"`
	WordsResultNum *int32                  `json:"words_result_num,omitempty"`
	WordsResult    *TrainTicketWordsResult `json:"words_result,omitempty"`
	PdfFileSize    *int32                  `json:"pdf_file_size,omitempty"`
}
