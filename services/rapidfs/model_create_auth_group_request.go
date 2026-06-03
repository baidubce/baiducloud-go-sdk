package rapidfs

type CreateAuthGroupRequest struct {
	ClientToken   *string     `json:"-"`
	AuthGroupName *string     `json:"authGroupName,omitempty"`
	InstanceId    *string     `json:"instanceId,omitempty"`
	Description   *string     `json:"description,omitempty"`
	AuthInfos     []*AuthInfo `json:"authInfos,omitempty"`
}
