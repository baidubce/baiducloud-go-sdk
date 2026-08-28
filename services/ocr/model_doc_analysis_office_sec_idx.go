package ocr

type DocAnalysisOfficeSecIdx struct {
	Idx     []*float64 `json:"idx,omitempty"`
	ParaIdx []*float64 `json:"para_idx,omitempty"`
	RowIdx  []*float64 `json:"row_idx,omitempty"`
	ColIdx  []*float64 `json:"col_idx,omitempty"`
}
