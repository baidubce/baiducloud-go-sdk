package cce

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	CCE_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	// ==== AK/SK 鉴权 ====
	CCE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_CreateAShrinkingNodeGroupTaskV2(t *testing.T) {
	DeleteOption := &DeleteOption{
		DeleteResource:    util.PtrBool(false),
		DeleteCDSSnapshot: util.PtrBool(false),
		MoveOut:           util.PtrBool(false),
	}
	createAShrinkingNodeGroupTaskV2Request := &CreateAShrinkingNodeGroupTaskV2Request{
		ClusterID:            util.PtrString(""),
		InstanceGroupID:      util.PtrString(""),
		InstancesToBeRemoved: []*string{},
		K8sNodesToBeRemoved:  []*string{},
		CleanPolicy:          util.PtrString(""),
		DeleteOption:         DeleteOption,
	}
	result := &CreateAShrinkingNodeGroupTaskV2Response{}
	result, err := CCE_CLIENT.CreateAShrinkingNodeGroupTaskV2(createAShrinkingNodeGroupTaskV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAnAutoscalerV2(t *testing.T) {
	createAnAutoscalerV2Request := &CreateAnAutoscalerV2Request{
		ClusterID: util.PtrString(""),
	}
	result := &CreateAnAutoscalerV2Response{}
	result, err := CCE_CLIENT.CreateAnAutoscalerV2(createAnAutoscalerV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateExpansionNodeGroupTaskV2(t *testing.T) {
	createExpansionNodeGroupTaskV2Request := &CreateExpansionNodeGroupTaskV2Request{
		ClusterID:       util.PtrString(""),
		InstanceGroupID: util.PtrString(""),
		UpToReplicas:    util.PtrInt32(int32(0)),
		UpReplicas:      util.PtrInt32(int32(0)),
	}
	result := &CreateExpansionNodeGroupTaskV2Response{}
	result, err := CCE_CLIENT.CreateExpansionNodeGroupTaskV2(createExpansionNodeGroupTaskV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateNodeGroupV2(t *testing.T) {
	InstanceTemplate := &InstanceTemplate{
		MachineType:               util.PtrString(""),
		InstanceType:              util.PtrString(""),
		InstanceName:              util.PtrString(""),
		VpcConfig:                 nil,
		InstanceResource:          nil,
		CheckGPUDriver:            util.PtrBool(false),
		ImageID:                   util.PtrString(""),
		UserData:                  nil,
		InstanceOS:                nil,
		ScaleDownDisabled:         util.PtrBool(false),
		IsOpenHostnameDomain:      util.PtrBool(false),
		NeedEIP:                   util.PtrBool(false),
		EipOption:                 nil,
		IamRole:                   nil,
		DeployCustomConfig:        nil,
		RuntimeType:               util.PtrString(""),
		RuntimeVersion:            util.PtrString(""),
		DeploySetIDs:              []*string{},
		Labels:                    nil,
		Annotations:               nil,
		Tags:                      []*string{},
		Taints:                    []*string{},
		RelationTag:               util.PtrBool(false),
		InstancePreChargingOption: nil,
	}
	ClusterAutoscalerSpec := &ClusterAutoscalerSpec{
		Enabled:              util.PtrBool(false),
		MinReplicas:          util.PtrInt32(int32(0)),
		MaxReplicas:          util.PtrInt32(int32(0)),
		ScalingGroupPriority: util.PtrInt32(int32(0)),
	}
	createNodeGroupV2Request := &CreateNodeGroupV2Request{
		ClusterID:             util.PtrString(""),
		InstanceGroupName:     util.PtrString(""),
		ClusterRole:           util.PtrString(""),
		ShrinkPolicy:          util.PtrString(""),
		UpdatePolicy:          util.PtrString(""),
		CleanPolicy:           util.PtrString(""),
		InstanceTemplate:      InstanceTemplate,
		Replicas:              util.PtrInt32(int32(0)),
		ClusterAutoscalerSpec: ClusterAutoscalerSpec,
	}
	result := &CreateNodeGroupV2Response{}
	result, err := CCE_CLIENT.CreateNodeGroupV2(createNodeGroupV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateNodesClusterExpansionV2(t *testing.T) {
	createNodesClusterExpansionV2Request := &CreateNodesClusterExpansionV2Request{
		ClusterID:   util.PtrString(""),
		RequestBody: []*InstanceSet{},
	}
	result := &CreateNodesClusterExpansionV2Response{}
	result, err := CCE_CLIENT.CreateNodesClusterExpansionV2(createNodesClusterExpansionV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteNodeGroupV2(t *testing.T) {
	deleteNodeGroupV2Request := &DeleteNodeGroupV2Request{
		ClusterID:          util.PtrString(""),
		InstanceGroupID:    util.PtrString(""),
		DeleteInstances:    util.PtrBool(false),
		ReleaseAllResource: util.PtrBool(false),
	}
	result := &DeleteNodeGroupV2Response{}
	result, err := CCE_CLIENT.DeleteNodeGroupV2(deleteNodeGroupV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteNodesClusterScalingV2(t *testing.T) {
	DeleteOption := &DeleteOption{
		DeleteResource:    util.PtrBool(false),
		DeleteCDSSnapshot: util.PtrBool(false),
		MoveOut:           util.PtrBool(false),
	}
	deleteNodesClusterScalingV2Request := &DeleteNodesClusterScalingV2Request{
		ClusterID:    util.PtrString(""),
		DeleteOption: DeleteOption,
		InstanceIDs:  []*string{},
		ScaleDown:    util.PtrBool(false),
	}
	result := &DeleteNodesClusterScalingV2Response{}
	result, err := CCE_CLIENT.DeleteNodesClusterScalingV2(deleteNodesClusterScalingV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetNodeDetailsV2(t *testing.T) {
	getNodeDetailsV2Request := &GetNodeDetailsV2Request{
		ClusterID:  util.PtrString(""),
		InstanceID: util.PtrString(""),
	}
	result := &GetNodeDetailsV2Response{}
	result, err := CCE_CLIENT.GetNodeDetailsV2(getNodeDetailsV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetNodeGroupDetailsV2(t *testing.T) {
	getNodeGroupDetailsV2Request := &GetNodeGroupDetailsV2Request{
		ClusterID:       util.PtrString(""),
		InstanceGroupID: util.PtrString(""),
	}
	result := &GetNodeGroupDetailsV2Response{}
	result, err := CCE_CLIENT.GetNodeGroupDetailsV2(getNodeGroupDetailsV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPackageListV2(t *testing.T) {
	getPackageListV2Request := &GetPackageListV2Request{
		Type:            util.PtrString(""),
		MachineSpecList: []*string{},
	}
	result := &GetPackageListV2Response{}
	result, err := CCE_CLIENT.GetPackageListV2(getPackageListV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTaskListV2(t *testing.T) {
	getTaskListV2Request := &GetTaskListV2Request{
		TaskType:      util.PtrString(""),
		TargetID:      util.PtrString(""),
		OperationType: util.PtrString(""),
		Phase:         util.PtrString(""),
		Order:         util.PtrString(""),
		OrderBy:       util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &GetTaskListV2Response{}
	result, err := CCE_CLIENT.GetTaskListV2(getTaskListV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTheListOfClusterNodeGroupsV2(t *testing.T) {
	getTheListOfClusterNodeGroupsV2Request := &GetTheListOfClusterNodeGroupsV2Request{
		ClusterID:         util.PtrString(""),
		PageNo:            util.PtrInt32(int32(0)),
		PageSize:          util.PtrInt32(int32(0)),
		KeywordType:       util.PtrString(""),
		Keyword:           util.PtrString(""),
		AutoscalerEnabled: util.PtrString(""),
		ChargingType:      util.PtrString(""),
	}
	result := &GetTheListOfClusterNodeGroupsV2Response{}
	result, err := CCE_CLIENT.GetTheListOfClusterNodeGroupsV2(getTheListOfClusterNodeGroupsV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTheListOfClusterNodesV2(t *testing.T) {
	getTheListOfClusterNodesV2Request := &GetTheListOfClusterNodesV2Request{
		ClusterID:   util.PtrString(""),
		KeywordType: util.PtrString(""),
		Keyword:     util.PtrString(""),
		OrderBy:     util.PtrString(""),
		Order:       util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &GetTheListOfClusterNodesV2Response{}
	result, err := CCE_CLIENT.GetTheListOfClusterNodesV2(getTheListOfClusterNodesV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyIGAutoScaler(t *testing.T) {
	modifyIGAutoScalerRequest := &ModifyIGAutoScalerRequest{
		ClusterID:            util.PtrString(""),
		InstanceGroupID:      util.PtrString(""),
		Enabled:              util.PtrBool(false),
		MinReplicas:          util.PtrInt32(int32(0)),
		MaxReplicas:          util.PtrInt32(int32(0)),
		ScalingGroupPriority: util.PtrInt32(int32(0)),
	}
	result := &ModifyIGAutoScalerResponse{}
	result, err := CCE_CLIENT.ModifyIGAutoScaler(modifyIGAutoScalerRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyNodeGroupNodeShrinkProtectionStatusV2(t *testing.T) {
	modifyNodeGroupNodeShrinkProtectionStatusV2Request := &ModifyNodeGroupNodeShrinkProtectionStatusV2Request{
		ClusterID:         util.PtrString(""),
		InstanceIDs:       []*string{},
		ScaleDownDisabled: util.PtrBool(false),
	}
	result := &ModifyNodeGroupNodeShrinkProtectionStatusV2Response{}
	result, err := CCE_CLIENT.ModifyNodeGroupNodeShrinkProtectionStatusV2(modifyNodeGroupNodeShrinkProtectionStatusV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTheNumberOfNodeReplicasInANodeGroupV2(t *testing.T) {
	DeleteOption := &DeleteOption{
		DeleteResource:    util.PtrBool(false),
		DeleteCDSSnapshot: util.PtrBool(false),
		MoveOut:           util.PtrBool(false),
	}
	modifyTheNumberOfNodeReplicasInANodeGroupV2Request := &ModifyTheNumberOfNodeReplicasInANodeGroupV2Request{
		ClusterID:       util.PtrString(""),
		InstanceGroupID: util.PtrString(""),
		Replicas:        util.PtrInt32(int32(0)),
		InstanceIDs:     []*string{},
		DeleteInstance:  util.PtrBool(false),
		DeleteOption:    DeleteOption,
	}
	result := &ModifyTheNumberOfNodeReplicasInANodeGroupV2Response{}
	result, err := CCE_CLIENT.ModifyTheNumberOfNodeReplicasInANodeGroupV2(modifyTheNumberOfNodeReplicasInANodeGroupV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MoveIntoAnExistingNodeV2(t *testing.T) {
	moveIntoAnExistingNodeV2Request := &MoveIntoAnExistingNodeV2Request{
		ClusterID:                          util.PtrString(""),
		InstanceGroupID:                    util.PtrString(""),
		InCluster:                          util.PtrBool(false),
		UseInstanceGroupConfig:             util.PtrBool(false),
		UseInstanceGroupConfigWithDiskInfo: util.PtrBool(false),
		InstallGpuDriver:                   util.PtrBool(false),
		ExistedInstances:                   []*InstanceSet{},
		ExistedInstancesInCluster:          []*ExistedInstanceInCluster{},
	}
	result := &MoveIntoAnExistingNodeV2Response{}
	result, err := CCE_CLIENT.MoveIntoAnExistingNodeV2(moveIntoAnExistingNodeV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTheConfigurationOfAutoscalerV2(t *testing.T) {
	queryTheConfigurationOfAutoscalerV2Request := &QueryTheConfigurationOfAutoscalerV2Request{
		ClusterID: util.PtrString(""),
	}
	result := &QueryTheConfigurationOfAutoscalerV2Response{}
	result, err := CCE_CLIENT.QueryTheConfigurationOfAutoscalerV2(queryTheConfigurationOfAutoscalerV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RetrieveTheNodeGroupNodeListV2(t *testing.T) {
	retrieveTheNodeGroupNodeListV2Request := &RetrieveTheNodeGroupNodeListV2Request{
		ClusterID:       util.PtrString(""),
		InstanceGroupID: util.PtrString(""),
		PageNo:          util.PtrInt32(int32(0)),
		PageSize:        util.PtrInt32(int32(0)),
	}
	result := &RetrieveTheNodeGroupNodeListV2Response{}
	result, err := CCE_CLIENT.RetrieveTheNodeGroupNodeListV2(retrieveTheNodeGroupNodeListV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StepsToObtainNodeEventsV2(t *testing.T) {
	stepsToObtainNodeEventsV2Request := &StepsToObtainNodeEventsV2Request{
		InstanceID: util.PtrString(""),
	}
	result := &StepsToObtainNodeEventsV2Response{}
	result, err := CCE_CLIENT.StepsToObtainNodeEventsV2(stepsToObtainNodeEventsV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SynchronizeNodeMetadataV2(t *testing.T) {
	synchronizeNodeMetadataV2Request := &SynchronizeNodeMetadataV2Request{
		ClusterID: util.PtrString(""),
	}
	result := &SynchronizeNodeMetadataV2Response{}
	result, err := CCE_CLIENT.SynchronizeNodeMetadataV2(synchronizeNodeMetadataV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAutoscalerConfigurationV2(t *testing.T) {
	CustomConfigs := make(map[string]string)
	updateAutoscalerConfigurationV2Request := &UpdateAutoscalerConfigurationV2Request{
		ClusterID:                        util.PtrString(""),
		Expander:                         util.PtrString(""),
		InstanceGroups:                   []*map[string]interface{}{},
		KubeVersion:                      util.PtrString(""),
		MaxEmptyBulkDelete:               util.PtrInt32(int32(0)),
		ScaleDownDelayAfterAdd:           util.PtrInt32(int32(0)),
		ScaleDownEnabled:                 util.PtrBool(false),
		ScaleDownGPUUtilizationThreshold: util.PtrInt32(int32(0)),
		ScaleDownUnneededTime:            util.PtrInt32(int32(0)),
		ScaleDownUtilizationThreshold:    util.PtrInt32(int32(0)),
		SkipNodesWithLocalStorage:        util.PtrBool(false),
		SkipNodesWithSystemPods:          util.PtrBool(false),
		CustomConfigs:                    nil,
	}
	result := &UpdateAutoscalerConfigurationV2Response{}
	result, err := CCE_CLIENT.UpdateAutoscalerConfigurationV2(updateAutoscalerConfigurationV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateNodeAttributesV2(t *testing.T) {
	Labels := make(map[string]string)
	Annotations := make(map[string]string)
	updateNodeAttributesV2Request := &UpdateNodeAttributesV2Request{
		ClusterID:           util.PtrString(""),
		InstanceID:          util.PtrString(""),
		Labels:              nil,
		Annotations:         nil,
		Taints:              []*Taint{},
		CceInstancePriority: util.PtrInt32(int32(0)),
	}
	result := &UpdateNodeAttributesV2Response{}
	result, err := CCE_CLIENT.UpdateNodeAttributesV2(updateNodeAttributesV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ViewTaskDetailsV2(t *testing.T) {
	viewTaskDetailsV2Request := &ViewTaskDetailsV2Request{
		TaskType: util.PtrString(""),
		TaskID:   util.PtrString(""),
	}
	result := &ViewTaskDetailsV2Response{}
	result, err := CCE_CLIENT.ViewTaskDetailsV2(viewTaskDetailsV2Request)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
