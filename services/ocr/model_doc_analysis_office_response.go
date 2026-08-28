package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DocAnalysisOfficeResponse struct {
	bce.BaseResponse
	ErrorCode        *int32                      `json:"error_code,omitempty"`
	ErrorMsg         *string                     `json:"error_msg,omitempty"`
	LogId            *int64                      `json:"log_id,omitempty"`
	ImgDirection     *int32                      `json:"img_direction,omitempty"`
	ResultsNum       *int32                      `json:"results_num,omitempty"`
	Results          []*DocAnalysisOfficeResult  `json:"results,omitempty"`
	LayoutsNum       *int32                      `json:"layouts_num,omitempty"`
	Layouts          []*DocAnalysisOfficeLayout  `json:"layouts,omitempty"`
	SecRows          *int32                      `json:"sec_rows,omitempty"`
	SecCols          *int32                      `json:"sec_cols,omitempty"`
	Sections         []*DocAnalysisOfficeSection `json:"sections,omitempty"`
	TableNum         *int32                      `json:"table_num,omitempty"`
	TablesResult     []*TableResult              `json:"tables_result,omitempty"`
	SealRecogNum     *int32                      `json:"seal_recog_num,omitempty"`
	SealRecogResults []*SealRecogResult          `json:"seal_recog_results,omitempty"`
	FormulaResult    []*FormulaResult            `json:"formula_result,omitempty"`
	Underline        []*interface{}              `json:"underline,omitempty"`
	PdfFileSize      *int32                      `json:"pdf_file_size,omitempty"`
	OfdFileSize      *string                     `json:"ofd_file_size,omitempty"`
}
