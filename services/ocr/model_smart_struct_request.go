package ocr

type SmartStructRequest struct {
	Image          *string `json:"image,omitempty"`
	Url            *string `json:"url,omitempty"`
	PdfFile        *string `json:"pdf_file,omitempty"`
	PdfFileNum     *int32  `json:"pdf_file_num,omitempty"`
	ReturnRelation *bool   `json:"return_relation,omitempty"`
}
