package image

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RemoveMoireResponse struct {
	bce.BaseResponse
	ErrorCode      *int32  `json:"error_code,omitempty"`
	ErrorMsg       *string `json:"error_msg,omitempty"`
	LogId          *int64  `json:"log_id,omitempty"`
	ImageProcessed *string `json:"image_processed,omitempty"`
	PdfFileSize    *int32  `json:"pdf_file_size,omitempty"`
}
