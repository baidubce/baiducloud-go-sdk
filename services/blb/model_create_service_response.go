package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateServiceResponse struct {
	bce.BaseResponse
	Service *string `json:"service,omitempty"`
}
