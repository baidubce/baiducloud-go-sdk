package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SyncGroupDelayInfoResponse struct {
	bce.BaseResponse
	DelayInfo []*DelayInfoConsoleItem `json:"delayInfo,omitempty"`
}
