package cce

type MoveIntoAnExistingNodeV2Request struct {
	ClusterID                          *string                     `json:"-"`
	InstanceGroupID                    *string                     `json:"-"`
	InCluster                          *bool                       `json:"inCluster,omitempty"`
	UseInstanceGroupConfig             *bool                       `json:"useInstanceGroupConfig,omitempty"`
	UseInstanceGroupConfigWithDiskInfo *bool                       `json:"useInstanceGroupConfigWithDiskInfo,omitempty"`
	InstallGpuDriver                   *bool                       `json:"installGpuDriver,omitempty"`
	ExistedInstances                   []*InstanceSet              `json:"existedInstances,omitempty"`
	ExistedInstancesInCluster          []*ExistedInstanceInCluster `json:"existedInstancesInCluster,omitempty"`
}
