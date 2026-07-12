package oos

type Property struct {
	Name          *string        `json:"name,omitempty"`
	Required      *bool          `json:"required,omitempty"`
	OosType       *string        `json:"type,omitempty"`
	Label         *string        `json:"label,omitempty"`
	Description   *string        `json:"description,omitempty"`
	Multiple      *bool          `json:"multiple,omitempty"`
	Options       []*interface{} `json:"options,omitempty"`
	SelectOptions []*Option      `json:"selectOptions,omitempty"`
	DefaultValue  *interface{}   `json:"defaultValue,omitempty"`
}
