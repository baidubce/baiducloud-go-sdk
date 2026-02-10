package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type EipInquiryResponse struct {
	bce.BaseResponse
	Prices *map[string]string `json:"prices,omitempty"`
}
