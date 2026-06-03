package rapidfs

type IDCCacheNodeInfo struct {
	SshUser     *string `json:"sshUser,omitempty"`
	SshPassword *string `json:"sshPassword,omitempty"`
	SshPort     *int32  `json:"sshPort,omitempty"`
}
