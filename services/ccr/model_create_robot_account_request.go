package ccr

type CreateRobotAccountRequest struct {
	InstanceId  *string            `json:"-"`
	Name        *string            `json:"name,omitempty"`
	Secret      *string            `json:"secret,omitempty"`
	Disable     *bool              `json:"disable,omitempty"`
	Duration    *int32             `json:"duration,omitempty"`
	Description *string            `json:"description,omitempty"`
	Permissions []*RobotPermission `json:"permissions,omitempty"`
}
