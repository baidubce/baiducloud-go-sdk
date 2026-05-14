package snic

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryAvailablePublicServicesResponse struct {
	bce.BaseResponse
	Services []*string `json:"services,omitempty"`
}
