package face

import "github.com/baidubce/baiducloud-go-sdk/bce"

type FaceDeleteResponse struct {
	bce.BaseResponse
	ErrorCode *int32  `json:"error_code,omitempty"`
	ErrorMsg  *string `json:"error_msg,omitempty"`
	LogId     *int64  `json:"log_id,omitempty"`
}
