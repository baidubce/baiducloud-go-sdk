package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteL3AccessRuleResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
