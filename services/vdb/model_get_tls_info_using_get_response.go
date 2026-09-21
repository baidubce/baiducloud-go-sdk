package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTLSInfoUsingGETResponse struct {
	bce.BaseResponse
	TlsExpiredDate *string `json:"tlsExpiredDate,omitempty"`
	TlsIssuedDate  *string `json:"tlsIssuedDate,omitempty"`
	TlsStatus      *string `json:"tlsStatus,omitempty"`
	TlsValidDays   *int32  `json:"tlsValidDays,omitempty"`
}
