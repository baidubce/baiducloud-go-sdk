package face

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UserAddResponse struct {
	bce.BaseResponse
	Result    *FaceRegResult `json:"result,omitempty"`
	ErrorCode *int32         `json:"error_code,omitempty"`
	ErrorMsg  *string        `json:"error_msg,omitempty"`
	LogId     *int64         `json:"log_id,omitempty"`
}
