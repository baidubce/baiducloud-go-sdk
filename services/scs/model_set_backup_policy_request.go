package scs

type SetBackupPolicyRequest struct {
	InstanceId *string `json:"-"`
	BackupTime *string `json:"backupTime,omitempty"`
	BackupDays *string `json:"backupDays,omitempty"`
	ExpireDay  *int32  `json:"expireDay,omitempty"`
	IsEncrypt  *string `json:"isEncrypt,omitempty"`
}
