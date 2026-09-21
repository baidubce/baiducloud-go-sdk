package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTLSCertificateUsingGETResponse struct {
	bce.BaseResponse
	CaContent *string `json:"caContent,omitempty"`
}
