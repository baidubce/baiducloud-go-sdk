package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SmartStructResponse struct {
	bce.BaseResponse
	ErrorCode   *int32                  `json:"error_code,omitempty"`
	ErrorMsg    *string                 `json:"error_msg,omitempty"`
	LogId       *int64                  `json:"log_id,omitempty"`
	PdfFileSize *int32                  `json:"pdf_file_size,omitempty"`
	ObjectIdNum *int32                  `json:"object_id_num,omitempty"`
	WordsResult *SmartStructWordsResult `json:"words_result,omitempty"`
}
