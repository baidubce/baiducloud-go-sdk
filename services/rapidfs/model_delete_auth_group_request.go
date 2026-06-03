package rapidfs

type DeleteAuthGroupRequest struct {
	ClientToken *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	AuthGroupId *string `json:"authGroupId,omitempty"`
}
