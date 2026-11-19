package pfs

type MountTargetBindConfig struct {
	MountTargetId *string `json:"mountTargetId,omitempty"`
	MountPath     *string `json:"mountPath,omitempty"`
}
