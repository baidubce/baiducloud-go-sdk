package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type TableGranularityQueryUserListApiResponse struct {
	bce.BaseResponse
	Users []*TableLevelAccountListResponse `json:"users,omitempty"`
}
