package rapidfs

type AuthGroupInfo struct {
	AuthGroupId   *string     `json:"authGroupId,omitempty"`
	AuthGroupName *string     `json:"authGroupName,omitempty"`
	InstanceId    *string     `json:"instanceId,omitempty"`
	Status        *string     `json:"status,omitempty"`
	Description   *string     `json:"description,omitempty"`
	AuthInfos     []*AuthInfo `json:"authInfos,omitempty"`
}
