package aigw

type FallbackConfig struct {
	Enabled            *bool   `json:"enabled,omitempty"`
	ServiceName        *string `json:"serviceName,omitempty"`
	ModelNameMode      *string `json:"modelNameMode,omitempty"`
	SpecifiedModelName *string `json:"specifiedModelName,omitempty"`
}
