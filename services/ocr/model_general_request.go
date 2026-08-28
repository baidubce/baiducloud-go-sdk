package ocr

type GeneralRequest struct {
	Image                *string `json:"image,omitempty"`
	Url                  *string `json:"url,omitempty"`
	PdfFile              *string `json:"pdf_file,omitempty"`
	PdfFileNum           *int32  `json:"pdf_file_num,omitempty"`
	OfdFile              *string `json:"ofd_file,omitempty"`
	OfdFileNum           *int32  `json:"ofd_file_num,omitempty"`
	RecognizeGranularity *string `json:"recognize_granularity,omitempty"`
	LanguageType         *string `json:"language_type,omitempty"`
	DetectDirection      *bool   `json:"detect_direction,omitempty"`
	DetectLanguage       *bool   `json:"detect_language,omitempty"`
	Paragraph            *bool   `json:"paragraph,omitempty"`
	VertexesLocation     *bool   `json:"vertexes_location,omitempty"`
	Probability          *bool   `json:"probability,omitempty"`
}
