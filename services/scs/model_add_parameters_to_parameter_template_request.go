package scs

type AddParametersToParameterTemplateRequest struct {
	TemplateShowId *string       `json:"-"`
	Parameters     []*Parameters `json:"parameters,omitempty"`
}
