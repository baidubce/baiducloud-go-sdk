package bci

type VolumeMount struct {
	Name      *string `json:"name,omitempty"`
	BciType   *string `json:"type,omitempty"`
	MountPath *string `json:"mountPath,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}
