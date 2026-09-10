package cce

type InstanceGroup struct {
	Spec      *interface{} `json:"spec,omitempty"`
	Status    *interface{} `json:"status,omitempty"`
	CreatedAt *string      `json:"createdAt,omitempty"`
}
