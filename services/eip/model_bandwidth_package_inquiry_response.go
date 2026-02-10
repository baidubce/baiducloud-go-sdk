package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BandwidthPackageInquiryResponse struct {
	bce.BaseResponse
	Prices *map[string]string `json:"prices,omitempty"`
}
