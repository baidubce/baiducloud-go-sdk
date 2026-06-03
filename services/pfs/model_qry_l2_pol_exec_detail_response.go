package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QryL2PolExecDetailResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
	Report    *string `json:"report,omitempty"`
}
