package scs

type CreateParameterTemplateRequest struct {
	Name          *string       `json:"name,omitempty"`
	Engine        *string       `json:"engine,omitempty"`
	EngineVersion *string       `json:"engineVersion,omitempty"`
	ClusterType   *string       `json:"clusterType,omitempty"`
	TemplateType  *int32        `json:"templateType,omitempty"`
	Comment       *string       `json:"comment,omitempty"`
	Parameters    []*Parameters `json:"parameters,omitempty"`
}
