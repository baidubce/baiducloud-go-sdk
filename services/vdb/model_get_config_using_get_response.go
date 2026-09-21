package vdb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetConfigUsingGETResponse struct {
	bce.BaseResponse
	AutoBackupConfig  *string `json:"autoBackupConfig,omitempty"`
	AutoBackupEnabled *bool   `json:"autoBackupEnabled,omitempty"`
	IsEncrypt         *string `json:"isEncrypt,omitempty"`
}
