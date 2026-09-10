package cce

type Autoscaler struct {
	ClusterID   *string   `json:"clusterID,omitempty"`
	ClusterName *string   `json:"clusterName,omitempty"`
	CaConfig    *CAConfig `json:"caConfig,omitempty"`
}
