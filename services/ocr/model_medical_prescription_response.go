package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type MedicalPrescriptionResponse struct {
	bce.BaseResponse
	ErrorCode           *int32                          `json:"error_code,omitempty"`
	ErrorMsg            *string                         `json:"error_msg,omitempty"`
	LogId               *int64                          `json:"log_id,omitempty"`
	CommonDataResultNum *int32                          `json:"CommonData_result_num,omitempty"`
	CostDetailRowNum    *int32                          `json:"CostDetail_row_num,omitempty"`
	WordsResult         *MedicalPrescriptionWordsResult `json:"words_result,omitempty"`
}
