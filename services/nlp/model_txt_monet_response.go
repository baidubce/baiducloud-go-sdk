package nlp

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TxtMonetResponse struct {
	bce.BaseResponse
	ErrorCode   *int32           `json:"error_code,omitempty"`
	ErrorMsg    *string          `json:"error_msg,omitempty"`
	LogId       *int64           `json:"log_id,omitempty"`
	ResultsList []*ResultContent `json:"results_list,omitempty"`
}
