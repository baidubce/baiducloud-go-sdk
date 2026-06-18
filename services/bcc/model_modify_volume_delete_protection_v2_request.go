package bcc

type ModifyVolumeDeleteProtectionV2Request struct {
	VolumeIds              []*string `json:"volumeIds,omitempty"`
	EnableDeleteProtection *bool     `json:"enableDeleteProtection,omitempty"`
}
