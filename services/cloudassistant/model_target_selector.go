package cloudassistant

type TargetSelector struct {
	InstanceType    *string       `json:"instanceType,omitempty"`
	Tags            []*Tag        `json:"tags,omitempty"`
	ImportInstances *TargetImport `json:"importInstances,omitempty"`
}
