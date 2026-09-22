package scs

type ParameterTemplateDeleteParametersRequest struct {
	TemplateShowId *string   `json:"-"`
	Parameters     []*string `json:"parameters,omitempty"`
}
