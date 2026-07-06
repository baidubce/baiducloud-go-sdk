package aihc

type Resource struct {
	Name     *string `json:"name,omitempty"`
	Quantity *int32  `json:"quantity,omitempty"`
}
