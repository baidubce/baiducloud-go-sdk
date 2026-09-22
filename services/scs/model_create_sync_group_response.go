package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSyncGroupResponse struct {
	bce.BaseResponse
	SyncGroupShowId *string `json:"syncGroupShowId,omitempty"`
}
