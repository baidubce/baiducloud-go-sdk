package aihc

type ModifyTheModelV2Request struct {
	ModelId     *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
