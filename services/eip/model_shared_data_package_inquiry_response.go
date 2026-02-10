package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SharedDataPackageInquiryResponse struct {
	bce.BaseResponse
	Price *string `json:"price,omitempty"`
}
