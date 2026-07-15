package cloudassistant

type Parameter struct {
	Name         *string `json:"name,omitempty"`
	Desc         *string `json:"desc,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
}
