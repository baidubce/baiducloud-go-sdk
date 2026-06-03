package ccr

type RobotPermissionAccess struct {
	Resource *string `json:"resource,omitempty"`
	Action   *string `json:"action,omitempty"`
}
