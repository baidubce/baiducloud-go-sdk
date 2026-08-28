package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type VehicleRegCertificateResponse struct {
	bce.BaseResponse
	ErrorCode   *int32                                   `json:"error_code,omitempty"`
	ErrorMsg    *string                                  `json:"error_msg,omitempty"`
	LogId       *int64                                   `json:"log_id,omitempty"`
	Direction   *string                                  `json:"direction,omitempty"`
	WordsResult *VehicleRegisterCertificationwordsResult `json:"words_result,omitempty"`
}
