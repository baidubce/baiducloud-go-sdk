package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateL2BucketLinkResponse struct {
	bce.BaseResponse
	RequestId    *string `json:"requestId,omitempty"`
	BucketLinkId *string `json:"bucketLinkId,omitempty"`
}
