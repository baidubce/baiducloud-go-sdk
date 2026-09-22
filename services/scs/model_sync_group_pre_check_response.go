package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SyncGroupPreCheckResponse struct {
	bce.BaseResponse
	CheckResult       []*CheckSyncGroupResultItem          `json:"checkResult,omitempty"`
	ConnectionResults []*SyncGroupInstanceConnectionResult `json:"connectionResults,omitempty"`
}
