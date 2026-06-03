package ccr

type RobotPermission struct {
	Kind      *string                  `json:"kind,omitempty"`
	Namespace *string                  `json:"namespace,omitempty"`
	Access    []*RobotPermissionAccess `json:"access,omitempty"`
}
