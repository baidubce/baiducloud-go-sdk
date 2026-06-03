package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateL2PolicyResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	PolicyId  *string `json:"policyId,omitempty"`
}
