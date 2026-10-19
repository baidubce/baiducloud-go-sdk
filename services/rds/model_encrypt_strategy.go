package rds

type EncryptStrategy struct {
	EncryptEnable            *bool   `json:"encryptEnable,omitempty"`
	KeyManagementServiceName *string `json:"keyManagementServiceName,omitempty"`
	KeyManagementType        *string `json:"keyManagementType,omitempty"`
	SecretKeyID              *string `json:"secretKeyID,omitempty"`
}
