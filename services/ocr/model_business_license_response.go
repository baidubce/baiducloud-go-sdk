package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BusinessLicenseResponse struct {
	bce.BaseResponse
	ErrorCode      *int32                  `json:"error_code,omitempty"`
	ErrorMsg       *string                 `json:"error_msg,omitempty"`
	LogId          *int64                  `json:"log_id,omitempty"`
	Direction      *int32                  `json:"direction,omitempty"`
	RiskType       *string                 `json:"risk_type,omitempty"`
	WordsResultNum *int32                  `json:"words_result_num,omitempty"`
	WordsResult    *interface{}            `json:"words_result,omitempty"`
	CardQuality    *BusinessLicenseQuality `json:"card_quality,omitempty"`
}
