package ccr

type Robot struct {
	Id           *int32             `json:"id,omitempty"`
	Name         *string            `json:"name,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Level        *string            `json:"level,omitempty"`
	Disable      *bool              `json:"disable,omitempty"`
	Duration     *int32             `json:"duration,omitempty"`
	ExpiresAt    *int32             `json:"expiresAt,omitempty"`
	CreationTime *string            `json:"creationTime,omitempty"`
	UpdateTime   *string            `json:"updateTime,omitempty"`
	Permissions  []*RobotPermission `json:"permissions,omitempty"`
}
