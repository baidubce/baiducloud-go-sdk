package rds

type EncryptPolicyRequest struct {
	EncryptEnable            *bool   `json:"encryptEnable,omitempty"`
	KeyManagementType        *string `json:"keyManagementType,omitempty"`
	KeyManagementServiceName *string `json:"keyManagementServiceName,omitempty"`
	SecretKeyID              *string `json:"secretKeyID,omitempty"`
}
