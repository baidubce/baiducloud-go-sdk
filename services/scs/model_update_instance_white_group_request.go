package scs

type UpdateInstanceWhiteGroupRequest struct {
	InstanceId    *string   `json:"-"`
	GroupName     *string   `json:"groupName,omitempty"`
	NewGroupName  *string   `json:"newGroupName,omitempty"`
	ClusterIpList []*string `json:"clusterIpList,omitempty"`
}
