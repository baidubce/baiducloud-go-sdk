package rapidfs

type ModifyAuthGroupRequest struct {
	ClientToken       *string     `json:"-"`
	AuthGroupId       *string     `json:"authGroupId,omitempty"`
	InstanceId        *string     `json:"instanceId,omitempty"`
	AuthGroupName     *string     `json:"authGroupName,omitempty"`
	Description       *string     `json:"description,omitempty"`
	AuthInfos         []*AuthInfo `json:"authInfos,omitempty"`
	OriginalAuthInfos []*AuthInfo `json:"originalAuthInfos,omitempty"`
}
