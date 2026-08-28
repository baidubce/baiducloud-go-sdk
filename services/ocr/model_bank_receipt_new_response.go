package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BankReceiptNewResponse struct {
	bce.BaseResponse
	ErrorCode      *int32       `json:"error_code,omitempty"`
	ErrorMsg       *string      `json:"error_msg,omitempty"`
	LogId          *int64       `json:"log_id,omitempty"`
	PdfFileSize    *int32       `json:"pdf_file_size,omitempty"`
	WordsResultNum *int32       `json:"words_result_num,omitempty"`
	WordsResult    *interface{} `json:"words_result,omitempty"`
}
