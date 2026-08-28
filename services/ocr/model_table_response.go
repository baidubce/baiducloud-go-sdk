package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TableResponse struct {
	bce.BaseResponse
	ErrorCode    *int32          `json:"error_code,omitempty"`
	ErrorMsg     *string         `json:"error_msg,omitempty"`
	LogId        *int64          `json:"log_id,omitempty"`
	TableNum     *int32          `json:"table_num,omitempty"`
	TablesResult []*TablesResult `json:"tables_result,omitempty"`
	PdfFileSize  *int32          `json:"pdf_file_size,omitempty"`
	ExcelFile    *string         `json:"excel_file,omitempty"`
}
