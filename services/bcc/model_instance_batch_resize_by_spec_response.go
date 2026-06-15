package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type InstanceBatchResizeBySpecResponse struct {
	bce.BaseResponse
	OrderUuidResults []*string `json:"orderUuidResults,omitempty"`
}
