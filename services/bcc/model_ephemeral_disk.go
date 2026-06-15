package bcc

type EphemeralDisk struct {
	StorageType  *string `json:"storageType,omitempty"`
	SizeInGB     *int32  `json:"sizeInGB,omitempty"`
	FreeSizeInGB *int32  `json:"freeSizeInGB,omitempty"`
}
