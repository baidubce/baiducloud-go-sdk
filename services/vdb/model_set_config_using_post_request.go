package vdb

type SetConfigUsingPOSTRequest struct {
	InstanceId        *string `json:"-"`
	EngineType        *string `json:"-"`
	AutoBackupConfig  *string `json:"autoBackupConfig,omitempty"`
	AutoBackupEnabled *bool   `json:"autoBackupEnabled,omitempty"`
	IsEncrypt         *string `json:"isEncrypt,omitempty"`
}
