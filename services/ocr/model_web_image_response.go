package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type WebImageResponse struct {
	bce.BaseResponse
	ErrorCode      *int32                 `json:"error_code,omitempty"`
	ErrorMsg       *string                `json:"error_msg,omitempty"`
	LogId          *int64                 `json:"log_id,omitempty"`
	Language       *int32                 `json:"language,omitempty"`
	Direction      *int32                 `json:"direction,omitempty"`
	WordsResult    []*WebImageWordsResult `json:"words_result,omitempty"`
	WordsResultNum *int32                 `json:"words_result_num,omitempty"`
	PdfFileSize    *int32                 `json:"pdf_file_size,omitempty"`
}
