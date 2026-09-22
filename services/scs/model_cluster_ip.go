package scs

type ClusterIP struct {
	GroupName *string   `json:"groupName,omitempty"`
	IpList    []*string `json:"ipList,omitempty"`
}
