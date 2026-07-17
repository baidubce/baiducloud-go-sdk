package cloudassistant

type ActionFilter struct {
	Id                 *string           `json:"id,omitempty"`
	Name               *string           `json:"name,omitempty"`
	CloudassistantType *string           `json:"type,omitempty"`
	Command            *CommandFilter    `json:"command,omitempty"`
	FileUpload         *FileUploadFilter `json:"fileUpload,omitempty"`
}
