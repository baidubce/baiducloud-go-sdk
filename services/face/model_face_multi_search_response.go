package face

import "github.com/baidubce/baiducloud-go-sdk/bce"

type FaceMultiSearchResponse struct {
	bce.BaseResponse
	Timestamp *int64           `json:"timestamp,omitempty"`
	Result    *FaceMultiResult `json:"result,omitempty"`
	ErrorCode *int32           `json:"error_code,omitempty"`
	ErrorMsg  *string          `json:"error_msg,omitempty"`
	LogId     *int64           `json:"log_id,omitempty"`
}
