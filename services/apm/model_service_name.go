package apm

type ServiceName struct {
	ServiceName        *string `json:"serviceName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty"`
	ServiceDisplayName *string `json:"serviceDisplayName,omitempty"`
	Language           *string `json:"language,omitempty"`
	IncludeLLM         *bool   `json:"includeLLM,omitempty"`
}
