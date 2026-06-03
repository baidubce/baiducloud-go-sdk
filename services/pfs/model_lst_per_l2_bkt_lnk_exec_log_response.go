package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type LstPerL2BktLnkExecLogResponse struct {
	bce.BaseResponse
	RequestId    *string        `json:"requestId,omitempty"`
	BucketLinkId *string        `json:"bucketLinkId,omitempty"`
	InstanceId   *string        `json:"instanceId,omitempty"`
	ExecuteInfos []*ExecuteInfo `json:"executeInfos,omitempty"`
}
