package pfs

type AvailableMountServiceInfo struct {
	MountTargetId       *string                    `json:"mountTargetId,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	State               *int32                     `json:"state,omitempty"`
	SubnetId            *string                    `json:"subnetId,omitempty"`
	VersionNumber       *string                    `json:"versionNumber,omitempty"`
	VpcId               *string                    `json:"vpcId,omitempty"`
	ZoneName            *string                    `json:"zoneName,omitempty"`
	MountTargetBindList []*MountTargetBindInfoItem `json:"mountTargetBindList,omitempty"`
}
