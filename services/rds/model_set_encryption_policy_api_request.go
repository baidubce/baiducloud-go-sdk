package rds

type SetEncryptionPolicyApiRequest struct {
	InstanceId   *string                       `json:"-"`
	BackupPolicy []*UpdateEncryptPolicyRequest `json:"backupPolicy,omitempty"`
}
