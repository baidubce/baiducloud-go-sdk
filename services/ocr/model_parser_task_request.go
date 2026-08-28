package ocr

type ParserTaskRequest struct {
	FileData           *string `json:"file_data,omitempty"`
	FileUrl            *string `json:"file_url,omitempty"`
	FileName           *string `json:"file_name,omitempty"`
	RecognizeFormula   *bool   `json:"recognize_formula,omitempty"`
	AnalysisChart      *bool   `json:"analysis_chart,omitempty"`
	AngleAdjust        *bool   `json:"angle_adjust,omitempty"`
	ParseImageLayout   *bool   `json:"parse_image_layout,omitempty"`
	LanguageType       *string `json:"language_type,omitempty"`
	SwitchDigitalWidth *string `json:"switch_digital_width,omitempty"`
	HtmlTableFormat    *bool   `json:"html_table_format,omitempty"`
	ReturnDocChunks    *string `json:"return_doc_chunks,omitempty"`
}
