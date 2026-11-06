package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SlowLogDownloadDetailsResponse struct {
	bce.BaseResponse
	URL             *string `json:"URL,omitempty"`
	DownloadExpires *string `json:"downloadExpires,omitempty"`
	DataBackupType  *string `json:"dataBackupType,omitempty"`
}
