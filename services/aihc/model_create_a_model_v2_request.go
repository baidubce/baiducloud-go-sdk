package aihc

type CreateAModelV2Request struct {
	Name             *string            `json:"name,omitempty"`
	Description      *string            `json:"description,omitempty"`
	ModelFormat      *string            `json:"modelFormat,omitempty"`
	Owner            *string            `json:"owner,omitempty"`
	VisibilityScope  *string            `json:"visibilityScope,omitempty"`
	InitVersionEntry *ModelVersionEntry `json:"initVersionEntry,omitempty"`
}
