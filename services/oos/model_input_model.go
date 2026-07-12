package oos

type InputModel struct {
	Name        *string        `json:"name,omitempty"`
	Required    *bool          `json:"required,omitempty"`
	OosType     *string        `json:"type,omitempty"`
	Description *string        `json:"description,omitempty"`
	Options     []*interface{} `json:"options,omitempty"`
	OosDefault  *interface{}   `json:"default,omitempty"`
}
