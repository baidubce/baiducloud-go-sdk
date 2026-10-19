package rds

type ZoneDetail struct {
	ZoneNames       []*string `json:"zoneNames,omitempty"`
	ApiZoneNames    []*string `json:"apiZoneNames,omitempty"`
	Available       *bool     `json:"available,omitempty"`
	ZoneNameStr     *string   `json:"zoneNameStr,omitempty"`
	DefaultSubnetId *string   `json:"defaultSubnetId,omitempty"`
	MaxCpuCount     *int32    `json:"maxCpuCount,omitempty"`
	MaxMemory       *int32    `json:"maxMemory,omitempty"`
	MaxStorage      *int32    `json:"maxStorage,omitempty"`
	SupportNvmeDisk *bool     `json:"supportNvmeDisk,omitempty"`
	ZoneType        *string   `json:"zoneType,omitempty"`
}
