package bcc

type PassInstanceModel struct {
	Application   *string `json:"application,omitempty"`
	InstanceCount *int32  `json:"instanceCount,omitempty"`
}
