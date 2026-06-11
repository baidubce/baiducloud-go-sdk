package aihc

type DatasetVersionEntry struct {
	Id          *string `json:"id,omitempty"`
	Version     *string `json:"version,omitempty"`
	Description *string `json:"description,omitempty"`
	StoragePath *string `json:"storagePath,omitempty"`
	MountPath   *string `json:"mountPath,omitempty"`
	CreateUser  *string `json:"createUser,omitempty"`
}
