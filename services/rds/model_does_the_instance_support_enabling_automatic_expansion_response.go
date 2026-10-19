package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DoesTheInstanceSupportEnablingAutomaticExpansionResponse struct {
	bce.BaseResponse
	SupportEnableDiskAutoResize *int32 `json:"supportEnableDiskAutoResize,omitempty"`
}
