package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ObtainSslEncryptionInformationResponse struct {
	bce.BaseResponse
	SslState *SslState `json:"sslState,omitempty"`
}
