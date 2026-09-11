package cce

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "cce." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V2 = "v2"

	CONSTANT_EVENT = "event"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_CLUSTER = "cluster"

	CONSTANT_INSTANCEGROUPS = "instancegroups"

	CONSTANT_CLUSTER_I_D = "[clusterID]"

	CONSTANT_INSTANCEGROUP = "instancegroup"

	CONSTANT_INSTANCE_GROUP_I_D = "[instanceGroupID]"

	CONSTANT_ATTACH_INSTANCES = "attachInstances"

	CONSTANT_INSTANCE_SCALE_DOWN = "instanceScaleDown"

	CONSTANT_AUTOSCALER = "autoscaler"

	CONSTANT_TASK = "task"

	CONSTANT_SCALEUP = "scaleup"

	CONSTANT_REPLICAS = "replicas"

	CONSTANT_INSTANCES = "instances"

	CONSTANT_TASKS = "tasks"

	CONSTANT_SYNC = "sync"

	CONSTANT_API = "api"

	CONSTANT_CCE = "cce"

	CONSTANT_ARTIFACT_SERVICE = "artifact-service"

	CONSTANT_V1 = "v1"

	CONSTANT_MACHINE_SPECS = "machine-specs"

	CONSTANT_SCALEDOWN = "scaledown"
)

// Client of cce service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getCreateAShrinkingNodeGroupTaskV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID + bce.URI_PREFIX + CONSTANT_SCALEDOWN
}
func getCreateAnAutoscalerV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_AUTOSCALER + bce.URI_PREFIX + ClusterID
}
func getCreateExpansionNodeGroupTaskV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID + bce.URI_PREFIX + CONSTANT_SCALEUP
}
func getCreateNodeGroupV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP
}
func getDeleteNodeGroupV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID
}
func getDeleteNodesClusterScalingV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCES
}
func getGetNodeDetailsV2Uri(ClusterID string, InstanceID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceID
}
func getGetNodeGroupDetailsV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID
}
func getGetPackageListV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_CCE + bce.URI_PREFIX + CONSTANT_ARTIFACT_SERVICE + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MACHINE_SPECS
}
func getGetTaskListV2Uri(TaskType string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TASKS + bce.URI_PREFIX + TaskType
}
func getGetTheListOfClusterNodeGroupsV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUPS
}
func getGetTheListOfClusterNodesV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCES
}
func getModifyIGAutoScalerUri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID + bce.URI_PREFIX + CONSTANT_AUTOSCALER
}
func getModifyNodeGroupNodeShrinkProtectionStatusV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCE_SCALE_DOWN
}
func getModifyTheNumberOfNodeReplicasInANodeGroupV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID + bce.URI_PREFIX + CONSTANT_REPLICAS
}
func getMoveIntoAnExistingNodeV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + CONSTANT_CLUSTER_I_D + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + CONSTANT_INSTANCE_GROUP_I_D + bce.URI_PREFIX + CONSTANT_ATTACH_INSTANCES
}
func getQueryTheConfigurationOfAutoscalerV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_AUTOSCALER + bce.URI_PREFIX + ClusterID
}
func getRetrieveTheNodeGroupNodeListV2Uri(ClusterID string, InstanceGroupID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCEGROUP + bce.URI_PREFIX + InstanceGroupID + bce.URI_PREFIX + CONSTANT_INSTANCES
}
func getStepsToObtainNodeEventsV2Uri(InstanceID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_EVENT + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceID
}
func getSynchronizeNodeMetadataV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SYNC + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCES
}
func getUpdateAutoscalerConfigurationV2Uri(ClusterID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_AUTOSCALER + bce.URI_PREFIX + ClusterID
}
func getUpdateNodeAttributesV2Uri(ClusterID string, InstanceID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER + bce.URI_PREFIX + ClusterID + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceID
}
func getViewTaskDetailsV2Uri(TaskType string, TaskID string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + TaskType + bce.URI_PREFIX + TaskID
}
