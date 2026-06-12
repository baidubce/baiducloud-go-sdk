package aihc

type CreateDatasetVersionRequest struct {
	DatasetId   *string `json:"-"`
	Description *string `json:"description,omitempty"`
	StoragePath *string `json:"storagePath,omitempty"`
	MountPath   *string `json:"mountPath,omitempty"`
}
