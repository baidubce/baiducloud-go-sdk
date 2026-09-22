package scs

type ModifyParameterTemplateNameRequest struct {
	TemplateShowId *string `json:"-"`
	Name           *string `json:"name,omitempty"`
}
