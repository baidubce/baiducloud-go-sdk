package ocr

type Body struct {
	CellLocation []*TablePoint   `json:"cell_location,omitempty"`
	RowStart     *int32          `json:"row_start,omitempty"`
	RowEnd       *int32          `json:"row_end,omitempty"`
	ColStart     *int32          `json:"col_start,omitempty"`
	ColEnd       *int32          `json:"col_end,omitempty"`
	Words        *string         `json:"words,omitempty"`
	Contents     []*TableContent `json:"contents,omitempty"`
}
