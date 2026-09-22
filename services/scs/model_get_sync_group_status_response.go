package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetSyncGroupStatusResponse struct {
	bce.BaseResponse
	SyncGroupShowId *string           `json:"syncGroupShowId,omitempty"`
	SyncStatus      []*SyncStatusItem `json:"syncStatus,omitempty"`
}
