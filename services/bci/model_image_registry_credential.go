package bci

type ImageRegistryCredential struct {
	Server   *string `json:"server,omitempty"`
	UserName *string `json:"userName,omitempty"`
	Password *string `json:"password,omitempty"`
}
