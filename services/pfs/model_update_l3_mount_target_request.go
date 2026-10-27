package pfs

type UpdateL3MountTargetRequest struct {
	Action          *string `json:"-"`
	InstanceId      *string `json:"instanceId,omitempty"`
	MountTargetId   *string `json:"mountTargetId,omitempty"`
	AdmissionStatus *int32  `json:"admissionStatus,omitempty"`
	Name            *string `json:"name,omitempty"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
}
