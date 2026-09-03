package aigw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateServiceResponse struct {
	bce.BaseResponse
	AddedCount *int32 `json:"addedCount,omitempty"`
}
