package ocr

type DocCropEnhanceRequest struct {
	Image       *string `json:"image,omitempty"`
	Url         *string `json:"url,omitempty"`
	PdfFile     *string `json:"pdf_file,omitempty"`
	PdfFileNum  *int32  `json:"pdf_file_num,omitempty"`
	ScanType    *int32  `json:"scan_type,omitempty"`
	Points      *string `json:"points,omitempty"`
	EnhanceType *int32  `json:"enhance_type,omitempty"`
}
