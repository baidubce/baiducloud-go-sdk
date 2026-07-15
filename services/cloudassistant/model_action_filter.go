package cloudassistant

type ActionFilter struct {
	CloudassistantType *string        `json:"type,omitempty"`
	Command            *CommandFilter `json:"command,omitempty"`
	InstanceType       *string        `json:"instanceType,omitempty"`
	Id                 *string        `json:"id,omitempty"`
	Name               *string        `json:"name,omitempty"`
	FileUpload         *interface{}   `json:"fileUpload,omitempty"`
}
