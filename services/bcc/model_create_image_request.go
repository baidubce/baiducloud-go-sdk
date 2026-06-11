package bcc

type CreateImageRequest struct {
	ImageName  *string `json:"imageName,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	SnapshotId *string `json:"snapshotId,omitempty"`
	EncryptKey *string `json:"encryptKey,omitempty"`
	RelateCds  *bool   `json:"relateCds,omitempty"`
	Detection  *bool   `json:"detection,omitempty"`
}
