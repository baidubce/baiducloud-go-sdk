package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type LogDetailsResponse struct {
	bce.BaseResponse
	LogId          *string `json:"logId,omitempty"`
	DownloadUrl    *string `json:"downloadUrl,omitempty"`
	DownloadExpire *string `json:"downloadExpire,omitempty"`
}
