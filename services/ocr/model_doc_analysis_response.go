package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DocAnalysisResponse struct {
	bce.BaseResponse
	ErrorCode       *int32                      `json:"error_code,omitempty"`
	ErrorMsg        *string                     `json:"error_msg,omitempty"`
	LogId           *int64                      `json:"log_id,omitempty"`
	ImgDirection    *int32                      `json:"img_direction,omitempty"`
	ResultsNum      *int32                      `json:"results_num,omitempty"`
	Results         []*DocAnalysisResult        `json:"results,omitempty"`
	FormulaResult   []*DocAnalysisFormulaResult `json:"formula_result,omitempty"`
	WordsResult     []*DocAnalysisWordsResult   `json:"words_result,omitempty"`
	LayoutsNum      *int32                      `json:"layouts_num,omitempty"`
	Layouts         []*DocAnalysisLayout        `json:"layouts,omitempty"`
	SecRows         *int32                      `json:"sec_rows,omitempty"`
	SecCols         *int32                      `json:"sec_cols,omitempty"`
	Sections        []*DocAnalysisSection       `json:"sections,omitempty"`
	LongDivision    []*LongDivision             `json:"long_division,omitempty"`
	LongDivisionNum *int32                      `json:"long_division_num,omitempty"`
	Underline       []*Underline                `json:"underline,omitempty"`
	PdfFileSize     *int32                      `json:"pdf_file_size,omitempty"`
}
