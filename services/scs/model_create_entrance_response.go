package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEntranceResponse struct {
	bce.BaseResponse
	Success *bool `json:"success,omitempty"`
	Result  *bool `json:"result,omitempty"`
}
