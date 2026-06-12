package aihc

type ModifyModelRequest struct {
	ModelId     *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
