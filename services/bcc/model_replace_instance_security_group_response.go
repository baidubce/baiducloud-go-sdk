package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ReplaceInstanceSecurityGroupResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
