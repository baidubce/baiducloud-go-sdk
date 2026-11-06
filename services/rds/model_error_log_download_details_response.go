package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ErrorLogDownloadDetailsResponse struct {
	bce.BaseResponse
	Url             *string `json:"url,omitempty"`
	DownloadExpires *string `json:"downloadExpires,omitempty"`
	DataBackupType  *string `json:"dataBackupType,omitempty"`
}
