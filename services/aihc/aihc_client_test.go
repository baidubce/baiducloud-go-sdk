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
func TestClient_ModifyModel(t *testing.T) {
	modifyModelRequest := &ModifyModelRequest{
		ModelId:     util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := AIHC_CLIENT.ModifyModel(modifyModelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
