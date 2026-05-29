package bls

type UpdateProjectRequest struct {
	Uuid        *string `json:"uuid,omitempty"`
	Description *bool   `json:"description,omitempty"`
}
