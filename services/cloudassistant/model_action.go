package cloudassistant

type Action struct {
	Id                     *string     `json:"id,omitempty"`
	Ref                    *string     `json:"ref,omitempty"`
	CloudassistantType     *string     `json:"type,omitempty"`
	Name                   *string     `json:"name,omitempty"`
	Alias                  *string     `json:"alias,omitempty"`
	Description            *string     `json:"description,omitempty"`
	TimeoutSecond          *int32      `json:"timeoutSecond,omitempty"`
	Command                *Command    `json:"command,omitempty"`
	FileUpload             *FileUpload `json:"fileUpload,omitempty"`
	SupportedInstanceTypes []*string   `json:"supportedInstanceTypes,omitempty"`
	CreatedTimestamp       *int64      `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp       *int64      `json:"updatedTimestamp,omitempty"`
}
