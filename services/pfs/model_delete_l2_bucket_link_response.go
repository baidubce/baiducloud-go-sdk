package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteL2BucketLinkResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
