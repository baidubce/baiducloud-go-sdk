package bcc

type CreateVolumeClusterRequest struct {
	ZoneName        *string     `json:"zoneName,omitempty"`
	ClusterName     *string     `json:"clusterName,omitempty"`
	ClusterSizeInGB *int32      `json:"clusterSizeInGB,omitempty"`
	StorageType     *string     `json:"storageType,omitempty"`
	PurchaseCount   *int32      `json:"purchaseCount,omitempty"`
	Billing         *Billing    `json:"billing,omitempty"`
	Tags            []*TagModel `json:"tags,omitempty"`
}
