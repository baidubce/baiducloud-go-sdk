package bcc

type VolumeClusterModel struct {
	ClusterId         *string `json:"clusterId,omitempty"`
	ClusterName       *string `json:"clusterName,omitempty"`
	CreatedTime       *string `json:"createdTime,omitempty"`
	ExpiredTime       *string `json:"expiredTime,omitempty"`
	Status            *string `json:"status,omitempty"`
	LogicalZone       *string `json:"logicalZone,omitempty"`
	ProductType       *string `json:"productType,omitempty"`
	ClusterType       *string `json:"clusterType,omitempty"`
	TotalCapacity     *int32  `json:"totalCapacity,omitempty"`
	UsedCapacity      *int32  `json:"usedCapacity,omitempty"`
	AvailableCapacity *int32  `json:"availableCapacity,omitempty"`
	ExpandingCapacity *int32  `json:"expandingCapacity,omitempty"`
	CreatedVolumeNum  *int32  `json:"createdVolumeNum,omitempty"`
	EnableAutoRenew   *bool   `json:"enableAutoRenew,omitempty"`
}
