package scs

type CreateInstanceWhiteGroupRequest struct {
	InstanceId    *string   `json:"-"`
	GroupName     *string   `json:"groupName,omitempty"`
	ClusterIpList []*string `json:"clusterIpList,omitempty"`
}
