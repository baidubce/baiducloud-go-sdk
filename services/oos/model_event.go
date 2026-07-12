package oos

type Event struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Service     *string   `json:"service,omitempty"`
	Region      *string   `json:"region,omitempty"`
	Resource    *Resource `json:"resource,omitempty"`
}
