package bcm

type AlarmResource struct {
	Scope        *string            `json:"scope,omitempty"`
	ResourceType *string            `json:"resourceType,omitempty"`
	Region       *string            `json:"region,omitempty"`
	Identifiers  *map[string]string `json:"identifiers,omitempty"`
	Properties   *map[string]string `json:"properties,omitempty"`
}
