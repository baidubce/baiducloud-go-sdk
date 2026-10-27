package rapidfs

type MountServiceBinding struct {
	MountServiceId   *string `json:"mountServiceId,omitempty"`
	MountServiceName *string `json:"mountServiceName,omitempty"`
	DataSrcId        *string `json:"dataSrcId,omitempty"`
	DataSrcName      *string `json:"dataSrcName,omitempty"`
	LogicalMountPath *string `json:"logicalMountPath,omitempty"`
	Status           *string `json:"status,omitempty"`
}
