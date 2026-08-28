package ocr

type HandwritingRequest struct {
	Image                *string `json:"image,omitempty"`
	Url                  *string `json:"url,omitempty"`
	PdfFile              *string `json:"pdf_file,omitempty"`
	PdfFileNum           *int32  `json:"pdf_file_num,omitempty"`
	OfdFile              *string `json:"ofd_file,omitempty"`
	OfdFileNum           *int32  `json:"ofd_file_num,omitempty"`
	RecognizeGranularity *string `json:"recognize_granularity,omitempty"`
	EngGranularity       *string `json:"eng_granularity,omitempty"`
	Probability          *bool   `json:"probability,omitempty"`
	DetectDirection      *bool   `json:"detect_direction,omitempty"`
	DetectAlteration     *bool   `json:"detect_alteration,omitempty"`
	LanguageType         *string `json:"language_type,omitempty"`
}
