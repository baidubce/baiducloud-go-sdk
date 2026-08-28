package ocr

type DocAnalysisRequest struct {
	Image                 *string `json:"image,omitempty"`
	Url                   *string `json:"url,omitempty"`
	PdfFile               *string `json:"pdf_file,omitempty"`
	PdfFileNum            *int32  `json:"pdf_file_num,omitempty"`
	LanguageType          *string `json:"language_type,omitempty"`
	ResultType            *string `json:"result_type,omitempty"`
	DetectDirection       *bool   `json:"detect_direction,omitempty"`
	LineProbability       *bool   `json:"line_probability,omitempty"`
	DispLinePoly          *bool   `json:"disp_line_poly,omitempty"`
	WordsType             *string `json:"words_type,omitempty"`
	LayoutAnalysis        *bool   `json:"layout_analysis,omitempty"`
	RecgFormula           *bool   `json:"recg_formula,omitempty"`
	RecgLongDivision      *bool   `json:"recg_long_division,omitempty"`
	DispUnderlineAnalysis *bool   `json:"disp_underline_analysis,omitempty"`
	RecgAlter             *bool   `json:"recg_alter,omitempty"`
}
