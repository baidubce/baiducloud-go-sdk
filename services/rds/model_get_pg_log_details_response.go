package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPgLogDetailsResponse struct {
	bce.BaseResponse
	PglogId          *string `json:"pglogId,omitempty"`
	PglogSizeInBytes *int64  `json:"pglogSizeInBytes,omitempty"`
	PglogStartTime   *string `json:"pglogStartTime,omitempty"`
	PglogEndTime     *string `json:"pglogEndTime,omitempty"`
	DownloadUrl      *string `json:"downloadUrl,omitempty"`
	DownloadExpires  *string `json:"downloadExpires,omitempty"`
}
