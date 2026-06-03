package rapidfs

type DeleteAuthGroupRequest struct {
	Action      *string `json:"-"`
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	AuthGroupId *string `json:"authGroupId,omitempty"`
}
