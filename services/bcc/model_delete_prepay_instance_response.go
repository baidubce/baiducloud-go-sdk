package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeletePrepayInstanceResponse struct {
	bce.BaseResponse
	SuccessResources   *InstanceDeleteResultModel `json:"successResources,omitempty"`
	FailResources      *InstanceDeleteResultModel `json:"failResources,omitempty"`
	InstanceRefundFlag *bool                      `json:"instanceRefundFlag,omitempty"`
}
