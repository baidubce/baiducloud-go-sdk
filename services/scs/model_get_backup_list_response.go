package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetBackupListResponse struct {
	bce.BaseResponse
	TotalCount *int32               `json:"totalCount,omitempty"`
	Backups    []*BatchBackupRecord `json:"backups,omitempty"`
}
