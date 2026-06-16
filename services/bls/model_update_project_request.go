package bls

type UpdateProjectRequest struct {
	Uuid        *string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	Top         *bool   `json:"top,omitempty"`
}
