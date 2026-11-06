package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ObtainSslEncryptionCertificateResponse struct {
	bce.BaseResponse
	CaContent *string `json:"caContent,omitempty"`
}
