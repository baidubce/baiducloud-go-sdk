package scs

type ParameterTemplateModifyParametersRequest struct {
	TemplateShowId *string       `json:"-"`
	Parameters     []*Parameters `json:"parameters,omitempty"`
}
