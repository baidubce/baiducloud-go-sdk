package cce

type Instance struct {
	Spec      *interface{} `json:"spec,omitempty"`
	Status    *interface{} `json:"status,omitempty"`
	CreatedAt *string      `json:"createdAt,omitempty"`
	UpdatedAt *string      `json:"updatedAt,omitempty"`
}
