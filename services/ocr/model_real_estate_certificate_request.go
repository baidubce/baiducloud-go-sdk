package ocr

type RealEstateCertificateRequest struct {
	Image       *string `json:"image,omitempty"`
	Url         *string `json:"url,omitempty"`
	PdfFile     *string `json:"pdf_file,omitempty"`
	PdfFileNum  *int32  `json:"pdf_file_num,omitempty"`
	Probability *bool   `json:"probability,omitempty"`
	Location    *bool   `json:"location,omitempty"`
}
