package pfs

type DeleteL3MountTargetRequest struct {
	Action        *string `json:"-"`
	InstanceId    *string `json:"instanceId,omitempty"`
	MountTargetId *string `json:"mountTargetId,omitempty"`
}
