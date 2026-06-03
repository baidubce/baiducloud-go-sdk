package rapidfs

type CceClusterInfo struct {
	ClusterId   *string   `json:"clusterId,omitempty"`
	ClusterName *string   `json:"clusterName,omitempty"`
	Zones       []*string `json:"zones,omitempty"`
}
