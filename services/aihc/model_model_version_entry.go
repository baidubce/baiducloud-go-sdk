package aihc

type ModelVersionEntry struct {
	Id            *string `json:"id,omitempty"`
	Version       *string `json:"version,omitempty"`
	Source        *string `json:"source,omitempty"`
	StorageBucket *string `json:"storageBucket,omitempty"`
	StoragePath   *string `json:"storagePath,omitempty"`
	ModelMetrics  *string `json:"modelMetrics,omitempty"`
	Description   *string `json:"description,omitempty"`
}
