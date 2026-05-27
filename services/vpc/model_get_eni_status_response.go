package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetEniStatusResponse struct {
	bce.BaseResponse
	Status *string `json:"status,omitempty"`
}
