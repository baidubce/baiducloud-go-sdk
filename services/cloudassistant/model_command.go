package cloudassistant

type Command struct {
	CloudassistantType *string      `json:"type,omitempty"`
	Content            *string      `json:"content,omitempty"`
	Scope              *string      `json:"scope,omitempty"`
	EnableParameter    *bool        `json:"enableParameter,omitempty"`
	Parameters         []*Parameter `json:"parameters,omitempty"`
	User               *string      `json:"user,omitempty"`
	WorkDir            *string      `json:"workDir,omitempty"`
	ExecParams         *interface{} `json:"execParams,omitempty"`
	WaitOnAgentMilli   *int32       `json:"waitOnAgentMilli,omitempty"`
}
