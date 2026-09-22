package face

import "github.com/baidubce/baiducloud-go-sdk/bce"

type FaceLandmarkResponse struct {
	bce.BaseResponse
	ErrorCode *int32          `json:"error_code,omitempty"`
	ErrorMsg  *string         `json:"error_msg,omitempty"`
	LogId     *int64          `json:"log_id,omitempty"`
	Timestamp *int64          `json:"timestamp,omitempty"`
	Cached    *int32          `json:"cached,omitempty"`
	Result    *FaceMarkResult `json:"result,omitempty"`
}
