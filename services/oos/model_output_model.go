package oos

type OutputModel struct {
	Name        *string `json:"name,omitempty"`
	OosType     *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
}
