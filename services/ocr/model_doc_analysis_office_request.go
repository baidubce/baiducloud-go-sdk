package ocr

type DocAnalysisOfficeRequest struct {
	Image                 *string `json:"image,omitempty"`
	Url                   *string `json:"url,omitempty"`
	PdfFile               *string `json:"pdf_file,omitempty"`
	PdfFileNum            *int32  `json:"pdf_file_num,omitempty"`
	OfdFile               *string `json:"ofd_file,omitempty"`
	OfdFileNum            *int32  `json:"ofd_file_num,omitempty"`
	LanguageType          *string `json:"language_type,omitempty"`
	ResultType            *string `json:"result_type,omitempty"`
	CharProbability       *bool   `json:"char_probability,omitempty"`
	DetectDirection       *bool   `json:"detect_direction,omitempty"`
	LineProbability       *bool   `json:"line_probability,omitempty"`
	DispLinePoly          *bool   `json:"disp_line_poly,omitempty"`
	WordsType             *string `json:"words_type,omitempty"`
	LayoutAnalysis        *bool   `json:"layout_analysis,omitempty"`
	RecgTables            *bool   `json:"recg_tables,omitempty"`
	RecogSeal             *bool   `json:"recog_seal,omitempty"`
	RecgFormula           *bool   `json:"recg_formula,omitempty"`
	EraseSeal             *bool   `json:"erase_seal,omitempty"`
	DispUnderlineAnalysis *bool   `json:"disp_underline_analysis,omitempty"`
}
