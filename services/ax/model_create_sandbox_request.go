package ax

type CreateSandboxRequest struct {
	TemplateID          *string                   `json:"templateID,omitempty"`
	Timeout             *int32                    `json:"timeout,omitempty"`
	Metadata            *map[string]string        `json:"metadata,omitempty"`
	EnvVars             *map[string]string        `json:"envVars,omitempty"`
	Secure              *bool                     `json:"secure,omitempty"`
	AllowInternetAccess *bool                     `json:"allow_internet_access,omitempty"`
	AutoPause           *bool                     `json:"autoPause,omitempty"`
	AutoResume          *map[string]interface{}   `json:"autoResume,omitempty"`
	RuntimeType         *string                   `json:"runtimeType,omitempty"`
	Mcp                 *map[string]interface{}   `json:"mcp,omitempty"`
	VolumeMounts        []*map[string]interface{} `json:"volumeMounts,omitempty"`
}
