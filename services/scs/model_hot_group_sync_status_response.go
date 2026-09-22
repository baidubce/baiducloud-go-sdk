package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type HotGroupSyncStatusResponse struct {
	bce.BaseResponse
	Followers []*Followers `json:"followers,omitempty"`
}
