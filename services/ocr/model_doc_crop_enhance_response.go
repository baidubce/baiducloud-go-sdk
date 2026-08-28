package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DocCropEnhanceResponse struct {
	bce.BaseResponse
	LogId          *int64                 `json:"log_id,omitempty"`
	ImageProcessed *string                `json:"image_processed,omitempty"`
	Points         []*DocCropEnhancePoint `json:"points,omitempty"`
	PdfFileSize    *int32                 `json:"pdf_file_size,omitempty"`
}
