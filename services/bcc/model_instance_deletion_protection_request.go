package bcc

type InstanceDeletionProtectionRequest struct {
	InstanceId         *string `json:"-"`
	DeletionProtection *int32  `json:"deletionProtection,omitempty"`
}
