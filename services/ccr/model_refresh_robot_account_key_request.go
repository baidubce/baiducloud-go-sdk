package ccr

type RefreshRobotAccountKeyRequest struct {
	InstanceId *string `json:"-"`
	RobotID    *string `json:"-"`
	Secret     *string `json:"secret,omitempty"`
}
