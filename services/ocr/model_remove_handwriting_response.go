package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RemoveHandwritingResponse struct {
	bce.BaseResponse
	LogId          *int64  `json:"log_id,omitempty"`
	ErrorCode      *int32  `json:"error_code,omitempty"`
	ErrorMsg       *string `json:"error_msg,omitempty"`
	ImageProcessed *string `json:"image_processed,omitempty"`
	PdfFileSize    *int32  `json:"pdf_file_size,omitempty"`
}
