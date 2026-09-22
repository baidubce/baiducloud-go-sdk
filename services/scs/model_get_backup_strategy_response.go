package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetBackupStrategyResponse struct {
	bce.BaseResponse
	BackupTime *string `json:"backupTime,omitempty"`
	BackupDays *string `json:"backupDays,omitempty"`
	ExpireDay  *int32  `json:"expireDay,omitempty"`
	IsEncrypt  *string `json:"isEncrypt,omitempty"`
}
