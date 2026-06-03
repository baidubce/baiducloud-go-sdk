package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type UpdPerL2BktLnkInfoResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
