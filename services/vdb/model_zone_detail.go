package vdb

type ZoneDetail struct {
	ApiZoneNames    []*string `json:"apiZoneNames,omitempty"`
	Available       *bool     `json:"available,omitempty"`
	DefaultSubnetId *string   `json:"defaultSubnetId,omitempty"`
	MaxCpuCount     *int32    `json:"maxCpuCount,omitempty"`
	MaxMemory       *int32    `json:"maxMemory,omitempty"`
	MaxStorage      *int32    `json:"maxStorage,omitempty"`
	StockState      *int32    `json:"stockState,omitempty"`
	ZoneNameStr     *string   `json:"zoneNameStr,omitempty"`
	ZoneNames       []*string `json:"zoneNames,omitempty"`
}
