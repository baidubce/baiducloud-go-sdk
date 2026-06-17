package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TermsEnumResponse struct {
	bce.BaseResponse
	Terms    []*string `json:"terms,omitempty"`
	BlsError *Error    `json:"error,omitempty"`
	Status   *int32    `json:"status,omitempty"`
}
