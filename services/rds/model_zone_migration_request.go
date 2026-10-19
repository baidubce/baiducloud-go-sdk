package rds

type ZoneMigrationRequest struct {
	InstanceId    *string      `json:"-"`
	MasterAzone   *string      `json:"master\_azone,omitempty"`
	BackupAzone   *string      `json:"backup\_azone,omitempty"`
	ZoneNames     []*string    `json:"zoneNames,omitempty"`
	Subnets       []*SubnetMap `json:"subnets,omitempty"`
	EffectiveTime *string      `json:"effectiveTime,omitempty"`
}
