package aihc

type CreateModelVersionRequest struct {
	ModelId       *string `json:"-"`
	StorageBucket *string `json:"storageBucket,omitempty"`
	StoragePath   *string `json:"storagePath,omitempty"`
	Description   *string `json:"description,omitempty"`
	Source        *string `json:"source,omitempty"`
	ModelMetrics  *string `json:"modelMetrics,omitempty"`
}
