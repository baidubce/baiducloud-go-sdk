package as

type EphemeralDisk struct {
	StorageType *string `json:"storageType,omitempty"`
	SizeInGB    *int32  `json:"sizeInGB,omitempty"`
}
