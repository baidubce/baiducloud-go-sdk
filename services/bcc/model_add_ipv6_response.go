package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddIpv6Response struct {
	bce.BaseResponse
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}
