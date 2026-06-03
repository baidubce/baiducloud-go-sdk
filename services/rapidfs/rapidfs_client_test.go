package rapidfs

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
	RAPIDFS_CLIENT *Client
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

	RAPIDFS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddCacheNodes(t *testing.T) {
	addCacheNodesRequest := &AddCacheNodesRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		RapidfsType: util.PtrString(""),
		CacheNodes:  []*AddCacheNodeInfo{},
	}
	result := &AddCacheNodesResponse{}
	result, err := RAPIDFS_CLIENT.AddCacheNodes(addCacheNodesRequest)
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
func TestClient_CancelCacheRuleJob(t *testing.T) {
	cancelCacheRuleJobRequest := &CancelCacheRuleJobRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		CacheRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.CancelCacheRuleJob(cancelCacheRuleJobRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CancelMetaSyncJob(t *testing.T) {
	cancelMetaSyncJobRequest := &CancelMetaSyncJobRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.CancelMetaSyncJob(cancelMetaSyncJobRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CheckBeforeAddCacheNodes(t *testing.T) {
	checkBeforeAddCacheNodesRequest := &CheckBeforeAddCacheNodesRequest{
		Action:      util.PtrString(""),
		InstanceId:  util.PtrString(""),
		RapidfsType: util.PtrString(""),
		CacheNodes:  []*AddCacheNodeInfo{},
	}
	result := &CheckBeforeAddCacheNodesResponse{}
	result, err := RAPIDFS_CLIENT.CheckBeforeAddCacheNodes(checkBeforeAddCacheNodesRequest)
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
func TestClient_CheckBeforeCreateInstance(t *testing.T) {
	checkBeforeCreateInstanceRequest := &CheckBeforeCreateInstanceRequest{
		Action:             util.PtrString(""),
		Zone:               util.PtrString(""),
		VpcId:              util.PtrString(""),
		SubnetId:           util.PtrString(""),
		ManagedMode:        util.PtrString(""),
		MetaSpec:           util.PtrString(""),
		DataSpec:           util.PtrString(""),
		RapidfsType:        util.PtrString(""),
		CapacityTiB:        util.PtrInt32(int32(0)),
		CceClusterId:       util.PtrString(""),
		AihcResourcePoolId: util.PtrString(""),
		K8sControllerId:    util.PtrString(""),
		K8sControllerToken: util.PtrString(""),
	}
	result := &CheckBeforeCreateInstanceResponse{}
	result, err := RAPIDFS_CLIENT.CheckBeforeCreateInstance(checkBeforeCreateInstanceRequest)
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
func TestClient_CreateAndAssignTag(t *testing.T) {
	createAndAssignTagRequest := &CreateAndAssignTagRequest{
		Action:       util.PtrString(""),
		ClientToken:  util.PtrString(""),
		TagResources: []*TagResource{},
	}
	err := RAPIDFS_CLIENT.CreateAndAssignTag(createAndAssignTagRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAuthGroup(t *testing.T) {
	createAuthGroupRequest := &CreateAuthGroupRequest{
		Action:        util.PtrString(""),
		ClientToken:   util.PtrString(""),
		AuthGroupName: util.PtrString(""),
		InstanceId:    util.PtrString(""),
		Description:   util.PtrString(""),
		AuthInfos:     []*AuthInfo{},
	}
	result := &CreateAuthGroupResponse{}
	result, err := RAPIDFS_CLIENT.CreateAuthGroup(createAuthGroupRequest)
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
func TestClient_CreateCacheRule(t *testing.T) {
	createCacheRuleRequest := &CreateCacheRuleRequest{
		Action:          util.PtrString(""),
		ClientToken:     util.PtrString(""),
		InstanceId:      util.PtrString(""),
		DataSrcId:       util.PtrString(""),
		CacheRuleName:   util.PtrString(""),
		RapidfsType:     util.PtrString(""),
		Directory:       util.PtrString(""),
		ExecuteOnCreate: util.PtrBool(false),
		Description:     util.PtrString(""),
	}
	result := &CreateCacheRuleResponse{}
	result, err := RAPIDFS_CLIENT.CreateCacheRule(createCacheRuleRequest)
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
func TestClient_CreateInstance(t *testing.T) {
	createInstanceRequest := &CreateInstanceRequest{
		Action:                      util.PtrString(""),
		ClientToken:                 util.PtrString(""),
		InstanceName:                util.PtrString(""),
		Description:                 util.PtrString(""),
		Zone:                        util.PtrString(""),
		VpcId:                       util.PtrString(""),
		SubnetId:                    util.PtrString(""),
		ManagedMode:                 util.PtrString(""),
		MetaSpec:                    util.PtrString(""),
		DataSpec:                    util.PtrString(""),
		RapidfsType:                 util.PtrString(""),
		CapacityTiB:                 util.PtrInt32(int32(0)),
		CceClusterId:                util.PtrString(""),
		AihcResourcePoolId:          util.PtrString(""),
		K8sControllerId:             util.PtrString(""),
		K8sControllerToken:          util.PtrString(""),
		TokenRefreshIntervalMinutes: util.PtrInt32(int32(0)),
		Tags:                        []*Tag{},
	}
	result := &CreateInstanceResponse{}
	result, err := RAPIDFS_CLIENT.CreateInstance(createInstanceRequest)
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
func TestClient_CreateMetaSyncRule(t *testing.T) {
	createMetaSyncRuleRequest := &CreateMetaSyncRuleRequest{
		Action:           util.PtrString(""),
		ClientToken:      util.PtrString(""),
		InstanceId:       util.PtrString(""),
		DataSrcId:        util.PtrString(""),
		MetaSyncRuleName: util.PtrString(""),
		RapidfsType:      util.PtrString(""),
		Directory:        util.PtrString(""),
		IntervalMinutes:  util.PtrInt32(int32(0)),
		ExecuteOnCreate:  util.PtrBool(false),
		EnableOnCreate:   util.PtrBool(false),
		Description:      util.PtrString(""),
	}
	result := &CreateMetaSyncRuleResponse{}
	result, err := RAPIDFS_CLIENT.CreateMetaSyncRule(createMetaSyncRuleRequest)
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
func TestClient_DeleteAuthGroup(t *testing.T) {
	deleteAuthGroupRequest := &DeleteAuthGroupRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		AuthGroupId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.DeleteAuthGroup(deleteAuthGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCacheRule(t *testing.T) {
	deleteCacheRuleRequest := &DeleteCacheRuleRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		CacheRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.DeleteCacheRule(deleteCacheRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstance(t *testing.T) {
	deleteInstanceRequest := &DeleteInstanceRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		Token:       util.PtrString(""),
	}
	result := &DeleteInstanceResponse{}
	result, err := RAPIDFS_CLIENT.DeleteInstance(deleteInstanceRequest)
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
func TestClient_DeleteMetaSyncRule(t *testing.T) {
	deleteMetaSyncRuleRequest := &DeleteMetaSyncRuleRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.DeleteMetaSyncRule(deleteMetaSyncRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeAihcResourcePools(t *testing.T) {
	describeAihcResourcePoolsRequest := &DescribeAihcResourcePoolsRequest{
		Action:  util.PtrString(""),
		VpcId:   util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		Marker:  util.PtrString(""),
	}
	result := &DescribeAihcResourcePoolsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeAihcResourcePools(describeAihcResourcePoolsRequest)
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
func TestClient_DescribeAllocatableDataSrcQuota(t *testing.T) {
	describeAllocatableDataSrcQuotaRequest := &DescribeAllocatableDataSrcQuotaRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		DataSrcId:  util.PtrString(""),
	}
	result := &DescribeAllocatableDataSrcQuotaResponse{}
	result, err := RAPIDFS_CLIENT.DescribeAllocatableDataSrcQuota(describeAllocatableDataSrcQuotaRequest)
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
func TestClient_DescribeAuthGroup(t *testing.T) {
	describeAuthGroupRequest := &DescribeAuthGroupRequest{
		Action:      util.PtrString(""),
		InstanceId:  util.PtrString(""),
		AuthGroupId: util.PtrString(""),
	}
	result := &DescribeAuthGroupResponse{}
	result, err := RAPIDFS_CLIENT.DescribeAuthGroup(describeAuthGroupRequest)
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
func TestClient_DescribeAuthGroups(t *testing.T) {
	describeAuthGroupsRequest := &DescribeAuthGroupsRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		Filters:    []*Filter{},
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeAuthGroupsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeAuthGroups(describeAuthGroupsRequest)
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
func TestClient_DescribeCacheDeployGroup(t *testing.T) {
	describeCacheDeployGroupRequest := &DescribeCacheDeployGroupRequest{
		Action:               util.PtrString(""),
		InstanceId:           util.PtrString(""),
		CacheDeployGroupName: util.PtrString(""),
		CacheDeployGroupNS:   util.PtrString(""),
	}
	result := &DescribeCacheDeployGroupResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheDeployGroup(describeCacheDeployGroupRequest)
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
func TestClient_DescribeCacheDeployGroups(t *testing.T) {
	describeCacheDeployGroupsRequest := &DescribeCacheDeployGroupsRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeCacheDeployGroupsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheDeployGroups(describeCacheDeployGroupsRequest)
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
func TestClient_DescribeCacheNode(t *testing.T) {
	describeCacheNodeRequest := &DescribeCacheNodeRequest{
		Action:      util.PtrString(""),
		InstanceId:  util.PtrString(""),
		CacheNodeId: util.PtrString(""),
	}
	result := &DescribeCacheNodeResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheNode(describeCacheNodeRequest)
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
func TestClient_DescribeCacheNodeBccCandidates(t *testing.T) {
	describeCacheNodeBccCandidatesRequest := &DescribeCacheNodeBccCandidatesRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		VpcId:      util.PtrString(""),
		Filters:    []*Filter{},
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeCacheNodeBccCandidatesResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheNodeBccCandidates(describeCacheNodeBccCandidatesRequest)
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
func TestClient_DescribeCacheNodeQuota(t *testing.T) {
	describeCacheNodeQuotaRequest := &DescribeCacheNodeQuotaRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	result := &DescribeCacheNodeQuotaResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheNodeQuota(describeCacheNodeQuotaRequest)
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
func TestClient_DescribeCacheNodes(t *testing.T) {
	describeCacheNodesRequest := &DescribeCacheNodesRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		Filters:    []*Filter{},
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeCacheNodesResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheNodes(describeCacheNodesRequest)
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
func TestClient_DescribeCacheRule(t *testing.T) {
	describeCacheRuleRequest := &DescribeCacheRuleRequest{
		Action:      util.PtrString(""),
		InstanceId:  util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		CacheRuleId: util.PtrString(""),
	}
	result := &DescribeCacheRuleResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheRule(describeCacheRuleRequest)
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
func TestClient_DescribeCacheRuleJobs(t *testing.T) {
	describeCacheRuleJobsRequest := &DescribeCacheRuleJobsRequest{
		Action:      util.PtrString(""),
		InstanceId:  util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		CacheRuleId: util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
		Marker:      util.PtrString(""),
	}
	result := &DescribeCacheRuleJobsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheRuleJobs(describeCacheRuleJobsRequest)
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
func TestClient_DescribeCacheRules(t *testing.T) {
	describeCacheRulesRequest := &DescribeCacheRulesRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		Filters:    []*Filter{},
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeCacheRulesResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCacheRules(describeCacheRulesRequest)
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
func TestClient_DescribeCceClusters(t *testing.T) {
	describeCceClustersRequest := &DescribeCceClustersRequest{
		Action:  util.PtrString(""),
		VpcId:   util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		Marker:  util.PtrString(""),
	}
	result := &DescribeCceClustersResponse{}
	result, err := RAPIDFS_CLIENT.DescribeCceClusters(describeCceClustersRequest)
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
func TestClient_DescribeDataSrc(t *testing.T) {
	describeDataSrcRequest := &DescribeDataSrcRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		DataSrcId:  util.PtrString(""),
	}
	result := &DescribeDataSrcResponse{}
	result, err := RAPIDFS_CLIENT.DescribeDataSrc(describeDataSrcRequest)
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
func TestClient_DescribeDataSrcs(t *testing.T) {
	describeDataSrcsRequest := &DescribeDataSrcsRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		Filters:    []*Filter{},
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeDataSrcsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeDataSrcs(describeDataSrcsRequest)
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
func TestClient_DescribeInstance(t *testing.T) {
	describeInstanceRequest := &DescribeInstanceRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	result := &DescribeInstanceResponse{}
	result, err := RAPIDFS_CLIENT.DescribeInstance(describeInstanceRequest)
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
func TestClient_DescribeInstances(t *testing.T) {
	describeInstancesRequest := &DescribeInstancesRequest{
		Action:  util.PtrString(""),
		Filters: []*Filter{},
		MaxKeys: util.PtrInt32(int32(0)),
		Marker:  util.PtrString(""),
	}
	result := &DescribeInstancesResponse{}
	result, err := RAPIDFS_CLIENT.DescribeInstances(describeInstancesRequest)
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
func TestClient_DescribeMetaSyncJobs(t *testing.T) {
	describeMetaSyncJobsRequest := &DescribeMetaSyncJobsRequest{
		Action:         util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
		MaxKeys:        util.PtrInt32(int32(0)),
		Marker:         util.PtrString(""),
	}
	result := &DescribeMetaSyncJobsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeMetaSyncJobs(describeMetaSyncJobsRequest)
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
func TestClient_DescribeMetaSyncRule(t *testing.T) {
	describeMetaSyncRuleRequest := &DescribeMetaSyncRuleRequest{
		Action:         util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
	}
	result := &DescribeMetaSyncRuleResponse{}
	result, err := RAPIDFS_CLIENT.DescribeMetaSyncRule(describeMetaSyncRuleRequest)
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
func TestClient_DescribeMetaSyncRules(t *testing.T) {
	describeMetaSyncRulesRequest := &DescribeMetaSyncRulesRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		Filters:    []*Filter{},
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &DescribeMetaSyncRulesResponse{}
	result, err := RAPIDFS_CLIENT.DescribeMetaSyncRules(describeMetaSyncRulesRequest)
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
func TestClient_DescribeOrder(t *testing.T) {
	describeOrderRequest := &DescribeOrderRequest{
		Action:  util.PtrString(""),
		OrderId: util.PtrString(""),
	}
	result := &DescribeOrderResponse{}
	result, err := RAPIDFS_CLIENT.DescribeOrder(describeOrderRequest)
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
func TestClient_DescribePrice(t *testing.T) {
	describePriceRequest := &DescribePriceRequest{
		Action:      util.PtrString(""),
		Currency:    util.PtrString(""),
		ManagedMode: util.PtrString(""),
		MetaSpec:    util.PtrString(""),
		DataSpec:    util.PtrString(""),
		CapacityTiB: util.PtrInt32(int32(0)),
	}
	result := &DescribePriceResponse{}
	result, err := RAPIDFS_CLIENT.DescribePrice(describePriceRequest)
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
func TestClient_DescribeSpecs(t *testing.T) {
	describeSpecsRequest := &DescribeSpecsRequest{
		Action:  util.PtrString(""),
		Filters: []*Filter{},
	}
	result := &DescribeSpecsResponse{}
	result, err := RAPIDFS_CLIENT.DescribeSpecs(describeSpecsRequest)
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
func TestClient_DescribeToken(t *testing.T) {
	describeTokenRequest := &DescribeTokenRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		TokenId:    util.PtrString(""),
	}
	result := &DescribeTokenResponse{}
	result, err := RAPIDFS_CLIENT.DescribeToken(describeTokenRequest)
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
func TestClient_DescribeZones(t *testing.T) {
	describeZonesRequest := &DescribeZonesRequest{
		Action: util.PtrString(""),
	}
	result := &DescribeZonesResponse{}
	result, err := RAPIDFS_CLIENT.DescribeZones(describeZonesRequest)
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
func TestClient_DisableMetaSyncRule(t *testing.T) {
	disableMetaSyncRuleRequest := &DisableMetaSyncRuleRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.DisableMetaSyncRule(disableMetaSyncRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableMetaSyncRule(t *testing.T) {
	enableMetaSyncRuleRequest := &EnableMetaSyncRuleRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.EnableMetaSyncRule(enableMetaSyncRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ExecuteCacheRuleJob(t *testing.T) {
	executeCacheRuleJobRequest := &ExecuteCacheRuleJobRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		CacheRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.ExecuteCacheRuleJob(executeCacheRuleJobRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ExecuteMetaSyncJob(t *testing.T) {
	executeMetaSyncJobRequest := &ExecuteMetaSyncJobRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcId:      util.PtrString(""),
		MetaSyncRuleId: util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.ExecuteMetaSyncJob(executeMetaSyncJobRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ImportDataSrc(t *testing.T) {
	importDataSrcRequest := &ImportDataSrcRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcName:    util.PtrString(""),
		Bucket:         util.PtrString(""),
		OtherAccount:   util.PtrBool(false),
		BucketAK:       util.PtrString(""),
		BucketSK:       util.PtrString(""),
		BucketStsToken: util.PtrString(""),
		DirPrefix:      util.PtrString(""),
		KeepSymlink:    util.PtrBool(false),
		AuthGroupId:    util.PtrString(""),
		Description:    util.PtrString(""),
		QuotaMiB:       util.PtrInt32(int32(0)),
	}
	result := &ImportDataSrcResponse{}
	result, err := RAPIDFS_CLIENT.ImportDataSrc(importDataSrcRequest)
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
func TestClient_ModifyAuthGroup(t *testing.T) {
	modifyAuthGroupRequest := &ModifyAuthGroupRequest{
		Action:            util.PtrString(""),
		ClientToken:       util.PtrString(""),
		AuthGroupId:       util.PtrString(""),
		InstanceId:        util.PtrString(""),
		AuthGroupName:     util.PtrString(""),
		Description:       util.PtrString(""),
		AuthInfos:         []*AuthInfo{},
		OriginalAuthInfos: []*AuthInfo{},
	}
	err := RAPIDFS_CLIENT.ModifyAuthGroup(modifyAuthGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyDataSrc(t *testing.T) {
	modifyDataSrcRequest := &ModifyDataSrcRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		InstanceId:  util.PtrString(""),
		KeepSymlink: util.PtrBool(false),
		AuthGroupId: util.PtrString(""),
		Description: util.PtrString(""),
		QuotaMiB:    util.PtrInt32(int32(0)),
	}
	err := RAPIDFS_CLIENT.ModifyDataSrc(modifyDataSrcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyMetaSyncRule(t *testing.T) {
	modifyMetaSyncRuleRequest := &ModifyMetaSyncRuleRequest{
		Action:           util.PtrString(""),
		ClientToken:      util.PtrString(""),
		InstanceId:       util.PtrString(""),
		DataSrcId:        util.PtrString(""),
		MetaSyncRuleId:   util.PtrString(""),
		MetaSyncRuleName: util.PtrString(""),
		IntervalMinutes:  util.PtrInt32(int32(0)),
		Description:      util.PtrString(""),
	}
	err := RAPIDFS_CLIENT.ModifyMetaSyncRule(modifyMetaSyncRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyToken(t *testing.T) {
	modifyTokenRequest := &ModifyTokenRequest{
		Action:                      util.PtrString(""),
		ClientToken:                 util.PtrString(""),
		InstanceId:                  util.PtrString(""),
		TokenId:                     util.PtrString(""),
		TokenRefreshIntervalMinutes: util.PtrInt32(int32(0)),
	}
	err := RAPIDFS_CLIENT.ModifyToken(modifyTokenRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveCacheNodes(t *testing.T) {
	removeCacheNodesRequest := &RemoveCacheNodesRequest{
		Action:               util.PtrString(""),
		ClientToken:          util.PtrString(""),
		InstanceId:           util.PtrString(""),
		CacheNodeIds:         []*string{},
		ForceRemoveOnOffline: util.PtrBool(false),
	}
	err := RAPIDFS_CLIENT.RemoveCacheNodes(removeCacheNodesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveDataSrc(t *testing.T) {
	removeDataSrcRequest := &RemoveDataSrcRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		DataSrcId:   util.PtrString(""),
		InstanceId:  util.PtrString(""),
		Token:       util.PtrString(""),
	}
	result := &RemoveDataSrcResponse{}
	result, err := RAPIDFS_CLIENT.RemoveDataSrc(removeDataSrcRequest)
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
func TestClient_ResizeInstance(t *testing.T) {
	resizeInstanceRequest := &ResizeInstanceRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		CapacityTiB: util.PtrInt32(int32(0)),
	}
	result := &ResizeInstanceResponse{}
	result, err := RAPIDFS_CLIENT.ResizeInstance(resizeInstanceRequest)
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
func TestClient_RestartCacheNodes(t *testing.T) {
	restartCacheNodesRequest := &RestartCacheNodesRequest{
		Action:       util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceId:   util.PtrString(""),
		CacheNodeIds: []*string{},
	}
	err := RAPIDFS_CLIENT.RestartCacheNodes(restartCacheNodesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StartCacheNodes(t *testing.T) {
	startCacheNodesRequest := &StartCacheNodesRequest{
		Action:       util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceId:   util.PtrString(""),
		CacheNodeIds: []*string{},
	}
	err := RAPIDFS_CLIENT.StartCacheNodes(startCacheNodesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StopCacheNodes(t *testing.T) {
	stopCacheNodesRequest := &StopCacheNodesRequest{
		Action:                util.PtrString(""),
		ClientToken:           util.PtrString(""),
		InstanceId:            util.PtrString(""),
		CacheNodeIds:          []*string{},
		MigrateDataBeforeStop: util.PtrBool(false),
	}
	err := RAPIDFS_CLIENT.StopCacheNodes(stopCacheNodesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
