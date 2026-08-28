package ocr

type HkMacauTaiwanExitentrypermitRequest struct {
	Image               *string `json:"image,omitempty"`
	Url                 *string `json:"url,omitempty"`
	PdfFile             *string `json:"pdf_file,omitempty"`
	PdfFileNum          *int32  `json:"pdf_file_num,omitempty"`
	ExitentrypermitType *string `json:"exitentrypermit_type,omitempty"`
	Probability         *bool   `json:"probability,omitempty"`
	Location            *bool   `json:"location,omitempty"`
}
