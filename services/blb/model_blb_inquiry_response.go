package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BlbInquiryResponse struct {
	bce.BaseResponse
	Prices []*Price `json:"prices,omitempty"`
}
