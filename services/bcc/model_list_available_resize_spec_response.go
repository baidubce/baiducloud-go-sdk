package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ListAvailableResizeSpecResponse struct {
	bce.BaseResponse
	SpecList []*string `json:"specList,omitempty"`
}
