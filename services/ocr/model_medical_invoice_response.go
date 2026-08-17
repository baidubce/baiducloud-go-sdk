package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type MedicalInvoiceResponse struct {
	bce.BaseResponse
	ErrorCode      *int32                     `json:"error_code,omitempty"`
	ErrorMsg       *string                    `json:"error_msg,omitempty"`
	LogId          *int64                     `json:"log_id,omitempty"`
	WordsResultNum *int32                     `json:"words_result_num,omitempty"`
	InvoiceType    *string                    `json:"InvoiceType,omitempty"`
	Province       *string                    `json:"Province,omitempty"`
	WordsResult    *MedicalInvoiceWordsResult `json:"words_result,omitempty"`
}
