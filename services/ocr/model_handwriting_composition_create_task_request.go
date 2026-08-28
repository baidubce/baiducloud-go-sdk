package ocr

type HandwritingCompositionCreateTaskRequest struct {
	Image                *string `json:"image,omitempty"`
	Url                  *string `json:"url,omitempty"`
	PdfFile              *string `json:"pdf_file,omitempty"`
	RecognizeGranularity *string `json:"recognize_granularity,omitempty"`
	PdfFileNum           *int32  `json:"pdf_file_num,omitempty"`
}
