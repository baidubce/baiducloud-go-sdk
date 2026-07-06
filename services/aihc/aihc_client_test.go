package aihc

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
	AIHC_CLIENT *Client
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

	AIHC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_BatchStopTrainingTasksV2(t *testing.T) {
	batchStopTrainingTasksV2Request := &BatchStopTrainingTasksV2Request{
		QueueID:        util.PtrString(""),
		ResourcePoolId: util.PtrString(""),
		JobList:        []*map[string]interface{}{},
		JobListJobId:   util.PtrString(""),
	}
	result := &BatchStopTrainingTasksV2Response{}
	result, err := AIHC_CLIENT.BatchStopTrainingTasksV2(batchStopTrainingTasksV2Request)
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
func TestClient_CreateDataset(t *testing.T) {
	InitVersionEntry := &DatasetVersionEntry{
		Id:          util.PtrString(""),
		Version:     util.PtrString(""),
		Description: util.PtrString(""),
		StoragePath: util.PtrString(""),
		MountPath:   util.PtrString(""),
		CreateUser:  util.PtrString(""),
	}
	createDatasetRequest := &CreateDatasetRequest{
		Name:             util.PtrString(""),
		StorageType:      util.PtrString(""),
		StorageInstance:  util.PtrString(""),
		ImportFormat:     util.PtrString(""),
		Description:      util.PtrString(""),
		Owner:            util.PtrString(""),
		VisibilityScope:  util.PtrString(""),
		VisibilityUser:   []*PermissionEntry{},
		VisibilityGroup:  []*PermissionEntry{},
		InitVersionEntry: InitVersionEntry,
	}
	result := &CreateDatasetResponse{}
	result, err := AIHC_CLIENT.CreateDataset(createDatasetRequest)
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
func TestClient_CreateDatasetVersion(t *testing.T) {
	createDatasetVersionRequest := &CreateDatasetVersionRequest{
		DatasetId:   util.PtrString(""),
		Description: util.PtrString(""),
		StoragePath: util.PtrString(""),
		MountPath:   util.PtrString(""),
	}
	result := &CreateDatasetVersionResponse{}
	result, err := AIHC_CLIENT.CreateDatasetVersion(createDatasetVersionRequest)
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
func TestClient_CreateJob(t *testing.T) {
	JobSpec := &JobSpec{
		Image: util.PtrString(""),
		ImageConfig: &ImageConfig{
			Username: util.PtrString(""),
			Password: util.PtrString(""),
		},
		Replicas:    util.PtrInt32(int32(0)),
		Resources:   []*Resource{},
		Envs:        []*Env{},
		EnableRDMA:  util.PtrBool(false),
		HostNetwork: util.PtrBool(false),
	}
	TensorboardConfig := &TensorboardConfig{
		Enable:  util.PtrBool(false),
		LogPath: util.PtrString(""),
	}
	AlertConfig := &AlertConfig{
		InstanceId:   util.PtrString(""),
		AlertItems:   []*string{},
		AihcFor:      util.PtrString(""),
		NotifyRuleId: util.PtrString(""),
	}
	AdvancedSettings := &AdvancedSettings{
		RuntimeEnv:            util.PtrString(""),
		SubmitterBackoffLimit: util.PtrInt32(int32(0)),
	}
	createJobRequest := &CreateJobRequest{
		ResourcePoolId:     util.PtrString(""),
		QueueID:            util.PtrString(""),
		Name:               util.PtrString(""),
		Queue:              util.PtrString(""),
		JobType:            util.PtrString(""),
		JobSpec:            JobSpec,
		Command:            util.PtrString(""),
		Labels:             []*Label{},
		Priority:           util.PtrString(""),
		Datasources:        []*DataSource{},
		EnableBccl:         util.PtrBool(false),
		FaultTolerance:     util.PtrBool(false),
		FaultToleranceArgs: util.PtrString(""),
		TensorboardConfig:  TensorboardConfig,
		AlertConfig:        AlertConfig,
		RetentionPeriod:    util.PtrString(""),
		AdvancedSettings:   AdvancedSettings,
	}
	result := &CreateJobResponse{}
	result, err := AIHC_CLIENT.CreateJob(createJobRequest)
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
func TestClient_CreateModel(t *testing.T) {
	InitVersionEntry := &ModelVersionEntry{
		Id:            util.PtrString(""),
		Version:       util.PtrString(""),
		Source:        util.PtrString(""),
		StorageBucket: util.PtrString(""),
		StoragePath:   util.PtrString(""),
		ModelMetrics:  util.PtrString(""),
		Description:   util.PtrString(""),
	}
	createModelRequest := &CreateModelRequest{
		Name:             util.PtrString(""),
		Description:      util.PtrString(""),
		ModelFormat:      util.PtrString(""),
		Owner:            util.PtrString(""),
		VisibilityScope:  util.PtrString(""),
		InitVersionEntry: InitVersionEntry,
	}
	result := &CreateModelResponse{}
	result, err := AIHC_CLIENT.CreateModel(createModelRequest)
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
func TestClient_CreateModelVersion(t *testing.T) {
	createModelVersionRequest := &CreateModelVersionRequest{
		ModelId:       util.PtrString(""),
		StorageBucket: util.PtrString(""),
		StoragePath:   util.PtrString(""),
		Description:   util.PtrString(""),
		Source:        util.PtrString(""),
		ModelMetrics:  util.PtrString(""),
	}
	result := &CreateModelVersionResponse{}
	result, err := AIHC_CLIENT.CreateModelVersion(createModelVersionRequest)
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
func TestClient_DeleteDataset(t *testing.T) {
	deleteDatasetRequest := &DeleteDatasetRequest{
		DatasetId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteDataset(deleteDatasetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDatasetVersion(t *testing.T) {
	deleteDatasetVersionRequest := &DeleteDatasetVersionRequest{
		DatasetId: util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteDatasetVersion(deleteDatasetVersionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteJob(t *testing.T) {
	deleteJobRequest := &DeleteJobRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
	}
	result := &DeleteJobResponse{}
	result, err := AIHC_CLIENT.DeleteJob(deleteJobRequest)
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
func TestClient_DeleteModel(t *testing.T) {
	deleteModelRequest := &DeleteModelRequest{
		ModelId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteModel(deleteModelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteModelVersion(t *testing.T) {
	deleteModelVersionRequest := &DeleteModelVersionRequest{
		ModelId:   util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteModelVersion(deleteModelVersionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeDataset(t *testing.T) {
	describeDatasetRequest := &DescribeDatasetRequest{
		DatasetId: util.PtrString(""),
	}
	result := &DescribeDatasetResponse{}
	result, err := AIHC_CLIENT.DescribeDataset(describeDatasetRequest)
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
func TestClient_DescribeDatasetVersion(t *testing.T) {
	describeDatasetVersionRequest := &DescribeDatasetVersionRequest{
		DatasetId: util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	result := &DescribeDatasetVersionResponse{}
	result, err := AIHC_CLIENT.DescribeDatasetVersion(describeDatasetVersionRequest)
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
func TestClient_DescribeDatasetVersions(t *testing.T) {
	describeDatasetVersionsRequest := &DescribeDatasetVersionsRequest{
		DatasetId:  util.PtrString(""),
		PageNumber: util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &DescribeDatasetVersionsResponse{}
	result, err := AIHC_CLIENT.DescribeDatasetVersions(describeDatasetVersionsRequest)
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
func TestClient_DescribeDatasets(t *testing.T) {
	describeDatasetsRequest := &DescribeDatasetsRequest{
		Keyword:          util.PtrString(""),
		StorageType:      util.PtrString(""),
		StorageInstances: util.PtrString(""),
		ImportFormat:     util.PtrString(""),
		PageNumber:       util.PtrInt32(int32(0)),
		PageSize:         util.PtrInt32(int32(0)),
	}
	result := &DescribeDatasetsResponse{}
	result, err := AIHC_CLIENT.DescribeDatasets(describeDatasetsRequest)
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
func TestClient_DescribeJob(t *testing.T) {
	describeJobRequest := &DescribeJobRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		NeedDetail:     util.PtrBool(false),
	}
	result := &DescribeJobResponse{}
	result, err := AIHC_CLIENT.DescribeJob(describeJobRequest)
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
func TestClient_DescribeJobEvents(t *testing.T) {
	describeJobEventsRequest := &DescribeJobEventsRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		StartTime:      util.PtrString(""),
		EndTime:        util.PtrString(""),
	}
	result := &DescribeJobEventsResponse{}
	result, err := AIHC_CLIENT.DescribeJobEvents(describeJobEventsRequest)
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
func TestClient_DescribeJobLogs(t *testing.T) {
	describeJobLogsRequest := &DescribeJobLogsRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		PodName:        util.PtrString(""),
		Keywords:       util.PtrString(""),
		StartTime:      util.PtrString(""),
		EndTime:        util.PtrString(""),
		MaxLines:       util.PtrString(""),
		ChunkSize:      util.PtrString(""),
		Marker:         util.PtrString(""),
	}
	result := &DescribeJobLogsResponse{}
	result, err := AIHC_CLIENT.DescribeJobLogs(describeJobLogsRequest)
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
func TestClient_DescribeJobMetrics(t *testing.T) {
	describeJobMetricsRequest := &DescribeJobMetricsRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		StartTime:      util.PtrString(""),
		EndTime:        util.PtrString(""),
		TimeStep:       util.PtrString(""),
		MetricType:     util.PtrString(""),
		RateInterval:   util.PtrString(""),
	}
	result := &DescribeJobMetricsResponse{}
	result, err := AIHC_CLIENT.DescribeJobMetrics(describeJobMetricsRequest)
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
func TestClient_DescribeJobNodes(t *testing.T) {
	describeJobNodesRequest := &DescribeJobNodesRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
	}
	result := &DescribeJobNodesResponse{}
	result, err := AIHC_CLIENT.DescribeJobNodes(describeJobNodesRequest)
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
func TestClient_DescribeJobWebterminal(t *testing.T) {
	describeJobWebterminalRequest := &DescribeJobWebterminalRequest{
		ResourcePoolId:         util.PtrString(""),
		QueueID:                util.PtrString(""),
		JobId:                  util.PtrString(""),
		PodName:                util.PtrString(""),
		HandshakeTimeoutSecond: util.PtrString(""),
		PingTimeoutSecond:      util.PtrString(""),
	}
	result := &DescribeJobWebterminalResponse{}
	result, err := AIHC_CLIENT.DescribeJobWebterminal(describeJobWebterminalRequest)
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
func TestClient_DescribeJobs(t *testing.T) {
	describeJobsRequest := &DescribeJobsRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		Queue:          util.PtrString(""),
		Status:         util.PtrString(""),
		KeywordType:    util.PtrString(""),
		Keyword:        util.PtrString(""),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
		PageNumber:     util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
	}
	result := &DescribeJobsResponse{}
	result, err := AIHC_CLIENT.DescribeJobs(describeJobsRequest)
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
func TestClient_DescribeModel(t *testing.T) {
	describeModelRequest := &DescribeModelRequest{
		ModelId: util.PtrString(""),
	}
	result := &DescribeModelResponse{}
	result, err := AIHC_CLIENT.DescribeModel(describeModelRequest)
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
func TestClient_DescribeModelVersion(t *testing.T) {
	describeModelVersionRequest := &DescribeModelVersionRequest{
		ModelId:   util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	result := &DescribeModelVersionResponse{}
	result, err := AIHC_CLIENT.DescribeModelVersion(describeModelVersionRequest)
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
func TestClient_DescribeModelVersions(t *testing.T) {
	describeModelVersionsRequest := &DescribeModelVersionsRequest{
		ModelId:    util.PtrString(""),
		PageNumber: util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &DescribeModelVersionsResponse{}
	result, err := AIHC_CLIENT.DescribeModelVersions(describeModelVersionsRequest)
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
func TestClient_DescribeModels(t *testing.T) {
	describeModelsRequest := &DescribeModelsRequest{
		Keyword:    util.PtrString(""),
		PageNumber: util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &DescribeModelsResponse{}
	result, err := AIHC_CLIENT.DescribeModels(describeModelsRequest)
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
func TestClient_DescribePodEvents(t *testing.T) {
	describePodEventsRequest := &DescribePodEventsRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		PodName:        util.PtrString(""),
		StartTime:      util.PtrString(""),
		EndTime:        util.PtrString(""),
	}
	result := &DescribePodEventsResponse{}
	result, err := AIHC_CLIENT.DescribePodEvents(describePodEventsRequest)
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
func TestClient_ModifyDataset(t *testing.T) {
	modifyDatasetRequest := &ModifyDatasetRequest{
		DatasetId:       util.PtrString(""),
		Name:            util.PtrString(""),
		Description:     util.PtrString(""),
		VisibilityScope: util.PtrString(""),
		VisibilityUser:  []*PermissionEntry{},
		VisibilityGroup: []*PermissionEntry{},
	}
	err := AIHC_CLIENT.ModifyDataset(modifyDatasetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyJob(t *testing.T) {
	modifyJobRequest := &ModifyJobRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		Priority:       util.PtrString(""),
	}
	result := &ModifyJobResponse{}
	result, err := AIHC_CLIENT.ModifyJob(modifyJobRequest)
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
func TestClient_ModifyModel(t *testing.T) {
	modifyModelRequest := &ModifyModelRequest{
		ModelId:     util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := AIHC_CLIENT.ModifyModel(modifyModelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StopJob(t *testing.T) {
	stopJobRequest := &StopJobRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
	}
	result := &StopJobResponse{}
	result, err := AIHC_CLIENT.StopJob(stopJobRequest)
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
