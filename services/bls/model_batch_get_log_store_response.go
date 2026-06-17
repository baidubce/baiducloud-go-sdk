package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BatchGetLogStoreResponse struct {
	bce.BaseResponse
	Code    *string           `json:"code,omitempty"`
	Success *bool             `json:"success,omitempty"`
	Result  []*LogStoreDetail `json:"result,omitempty"`
}
