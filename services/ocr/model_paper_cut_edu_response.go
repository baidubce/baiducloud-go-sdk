package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type PaperCutEduResponse struct {
	bce.BaseResponse
	ErrorCode       *int32       `json:"error_code,omitempty"`
	ErrorMsg        *string      `json:"error_msg,omitempty"`
	LogId           *int64       `json:"log_id,omitempty"`
	Direction       *int32       `json:"direction,omitempty"`
	QusResultNum    *int32       `json:"qus_result_num,omitempty"`
	QusFigure       []*QusFigure `json:"qus_figure,omitempty"`
	QusResult       []*QusResult `json:"qus_result,omitempty"`
	PdfFileSize     *int32       `json:"pdf_file_size,omitempty"`
	ProcessedStatus *string      `json:"processed_status,omitempty"`
}
