package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteInstUserOpAuthorizeRuleResponse struct {
	bce.BaseResponse
	RequestId *string `json:"requestId,omitempty"`
}
