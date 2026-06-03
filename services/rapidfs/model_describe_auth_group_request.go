package rapidfs

type DescribeAuthGroupRequest struct {
	Action      *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	AuthGroupId *string `json:"authGroupId,omitempty"`
}
