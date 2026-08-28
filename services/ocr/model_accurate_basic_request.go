package ocr

type AccurateBasicRequest struct {
	Image                     *string `json:"image,omitempty"`
	Url                       *string `json:"url,omitempty"`
	PdfFile                   *string `json:"pdf_file,omitempty"`
	PdfFileNum                *int32  `json:"pdf_file_num,omitempty"`
	OfdFile                   *string `json:"ofd_file,omitempty"`
	OfdFileNum                *int32  `json:"ofd_file_num,omitempty"`
	LanguageType              *string `json:"language_type,omitempty"`
	DetectDirection           *bool   `json:"detect_direction,omitempty"`
	Paragraph                 *bool   `json:"paragraph,omitempty"`
	Probability               *bool   `json:"probability,omitempty"`
	MultidirectionalRecognize *bool   `json:"multidirectional_recognize,omitempty"`
}
