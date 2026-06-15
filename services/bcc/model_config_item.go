package bcc

type ConfigItem struct {
	Cpu      *int32  `json:"cpu,omitempty"`
	Memory   *int32  `json:"memory,omitempty"`
	BccType  *string `json:"type,omitempty"`
	SpecId   *string `json:"specId,omitempty"`
	Spec     *string `json:"spec,omitempty"`
	ZoneName *string `json:"zoneName,omitempty"`
}
