package ccr

type RegistryCredential struct {
	AccessKey    *string `json:"accessKey,omitempty"`
	AccessSecret *string `json:"accessSecret,omitempty"`
	CcrType      *string `json:"type,omitempty"`
}
