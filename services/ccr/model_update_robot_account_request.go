package ccr

type UpdateRobotAccountRequest struct {
	InstanceId  *string            `json:"-"`
	RobotID     *string            `json:"-"`
	Disable     *bool              `json:"disable,omitempty"`
	Duration    *int32             `json:"duration,omitempty"`
	Description *string            `json:"description,omitempty"`
	Permissions []*RobotPermission `json:"permissions,omitempty"`
}
