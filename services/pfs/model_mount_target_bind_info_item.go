package pfs

type MountTargetBindInfoItem struct {
	MountTargetBindId *string `json:"mountTargetBindId,omitempty"`
	MountTargetId     *string `json:"mountTargetId,omitempty"`
	MountTargetName   *string `json:"mountTargetName,omitempty"`
	InstanceId        *string `json:"instanceId,omitempty"`
	InstanceName      *string `json:"instanceName,omitempty"`
	Region            *string `json:"region,omitempty"`
	Zone              *string `json:"zone,omitempty"`
	LocalMountPath    *string `json:"localMountPath,omitempty"`
	BindTime          *string `json:"bindTime,omitempty"`
	State             *int32  `json:"state,omitempty"`
	IsPrimaryInstance *bool   `json:"isPrimaryInstance,omitempty"`
}
