package pfs

type AttachedMountTargetInfo struct {
	MountTargetId   *string `json:"mountTargetId,omitempty"`
	MountTargetName *string `json:"mountTargetName,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty"`
	InstanceName    *string `json:"instanceName,omitempty"`
}
