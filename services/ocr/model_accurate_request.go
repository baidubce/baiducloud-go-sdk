package ocr

type AccurateRequest struct {
	Image                     *string `json:"image,omitempty"`
	Url                       *string `json:"url,omitempty"`
	PdfFile                   *string `json:"pdf_file,omitempty"`
	PdfFileNum                *int32  `json:"pdf_file_num,omitempty"`
	OfdFile                   *string `json:"ofd_file,omitempty"`
	OfdFileNum                *int32  `json:"ofd_file_num,omitempty"`
	LanguageType              *string `json:"language_type,omitempty"`
	EngGranularity            *string `json:"eng_granularity,omitempty"`
	RecognizeGranularity      *string `json:"recognize_granularity,omitempty"`
	DetectDirection           *bool   `json:"detect_direction,omitempty"`
	VertexesLocation          *bool   `json:"vertexes_location,omitempty"`
	Paragraph                 *bool   `json:"paragraph,omitempty"`
	Probability               *bool   `json:"probability,omitempty"`
	CharProbability           *bool   `json:"char_probability,omitempty"`
	MultidirectionalRecognize *bool   `json:"multidirectional_recognize,omitempty"`
}
