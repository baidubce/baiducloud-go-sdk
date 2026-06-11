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

func TestClient_CreateADatasetV2(t *testing.T) {
	InitVersionEntry := &DatasetVersionEntry{
		Id:          util.PtrString(""),
		Version:     util.PtrString(""),
		Description: util.PtrString(""),
		StoragePath: util.PtrString(""),
		MountPath:   util.PtrString(""),
		CreateUser:  util.PtrString(""),
	}
	createADatasetV2Request := &CreateADatasetV2Request{
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
	result := &CreateADatasetV2Response{}
	result, err := AIHC_CLIENT.CreateADatasetV2(createADatasetV2Request)
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
func TestClient_CreateAModelV2(t *testing.T) {
	InitVersionEntry := &ModelVersionEntry{
		Id:            util.PtrString(""),
		Version:       util.PtrString(""),
		Source:        util.PtrString(""),
		StorageBucket: util.PtrString(""),
		StoragePath:   util.PtrString(""),
		ModelMetrics:  util.PtrString(""),
		Description:   util.PtrString(""),
	}
	createAModelV2Request := &CreateAModelV2Request{
		Name:             util.PtrString(""),
		Description:      util.PtrString(""),
		ModelFormat:      util.PtrString(""),
		Owner:            util.PtrString(""),
		VisibilityScope:  util.PtrString(""),
		InitVersionEntry: InitVersionEntry,
	}
	result := &CreateAModelV2Response{}
	result, err := AIHC_CLIENT.CreateAModelV2(createAModelV2Request)
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
func TestClient_CreateDatasetVersionV2(t *testing.T) {
	createDatasetVersionV2Request := &CreateDatasetVersionV2Request{
		DatasetId:   util.PtrString(""),
		Description: util.PtrString(""),
		StoragePath: util.PtrString(""),
		MountPath:   util.PtrString(""),
	}
	result := &CreateDatasetVersionV2Response{}
	result, err := AIHC_CLIENT.CreateDatasetVersionV2(createDatasetVersionV2Request)
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
func TestClient_DeleteDatasetV2(t *testing.T) {
	deleteDatasetV2Request := &DeleteDatasetV2Request{
		DatasetId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteDatasetV2(deleteDatasetV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDatasetVersionV2(t *testing.T) {
	deleteDatasetVersionV2Request := &DeleteDatasetVersionV2Request{
		DatasetId: util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteDatasetVersionV2(deleteDatasetVersionV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteModelV2(t *testing.T) {
	deleteModelV2Request := &DeleteModelV2Request{
		ModelId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteModelV2(deleteModelV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteModelVersionV2(t *testing.T) {
	deleteModelVersionV2Request := &DeleteModelVersionV2Request{
		ModelId:   util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	err := AIHC_CLIENT.DeleteModelVersionV2(deleteModelVersionV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetAListOfModelVersionsV2(t *testing.T) {
	getAListOfModelVersionsV2Request := &GetAListOfModelVersionsV2Request{
		ModelId:    util.PtrString(""),
		PageNumber: util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &GetAListOfModelVersionsV2Response{}
	result, err := AIHC_CLIENT.GetAListOfModelVersionsV2(getAListOfModelVersionsV2Request)
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
func TestClient_GetDatasetDetailsV2(t *testing.T) {
	getDatasetDetailsV2Request := &GetDatasetDetailsV2Request{
		DatasetId: util.PtrString(""),
	}
	result := &GetDatasetDetailsV2Response{}
	result, err := AIHC_CLIENT.GetDatasetDetailsV2(getDatasetDetailsV2Request)
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
func TestClient_GetDatasetVersionDetailsV2(t *testing.T) {
	getDatasetVersionDetailsV2Request := &GetDatasetVersionDetailsV2Request{
		DatasetId: util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	result := &GetDatasetVersionDetailsV2Response{}
	result, err := AIHC_CLIENT.GetDatasetVersionDetailsV2(getDatasetVersionDetailsV2Request)
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
func TestClient_GetModelDetailsV2(t *testing.T) {
	getModelDetailsV2Request := &GetModelDetailsV2Request{
		ModelId: util.PtrString(""),
	}
	result := &GetModelDetailsV2Response{}
	result, err := AIHC_CLIENT.GetModelDetailsV2(getModelDetailsV2Request)
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
func TestClient_GetModelListV2(t *testing.T) {
	getModelListV2Request := &GetModelListV2Request{
		Keyword:    util.PtrString(""),
		PageNumber: util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &GetModelListV2Response{}
	result, err := AIHC_CLIENT.GetModelListV2(getModelListV2Request)
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
func TestClient_GetModelVersionDetailsV2(t *testing.T) {
	getModelVersionDetailsV2Request := &GetModelVersionDetailsV2Request{
		ModelId:   util.PtrString(""),
		VersionId: util.PtrString(""),
	}
	result := &GetModelVersionDetailsV2Response{}
	result, err := AIHC_CLIENT.GetModelVersionDetailsV2(getModelVersionDetailsV2Request)
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
func TestClient_ModifyDatasetV2(t *testing.T) {
	modifyDatasetV2Request := &ModifyDatasetV2Request{
		DatasetId:       util.PtrString(""),
		Name:            util.PtrString(""),
		Description:     util.PtrString(""),
		VisibilityScope: util.PtrString(""),
		VisibilityUser:  []*PermissionEntry{},
		VisibilityGroup: []*PermissionEntry{},
	}
	err := AIHC_CLIENT.ModifyDatasetV2(modifyDatasetV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTheModelV2(t *testing.T) {
	modifyTheModelV2Request := &ModifyTheModelV2Request{
		ModelId:     util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := AIHC_CLIENT.ModifyTheModelV2(modifyTheModelV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_NewModelVersionV2(t *testing.T) {
	newModelVersionV2Request := &NewModelVersionV2Request{
		ModelId:       util.PtrString(""),
		StorageBucket: util.PtrString(""),
		StoragePath:   util.PtrString(""),
		Description:   util.PtrString(""),
		Source:        util.PtrString(""),
		ModelMetrics:  util.PtrString(""),
	}
	result := &NewModelVersionV2Response{}
	result, err := AIHC_CLIENT.NewModelVersionV2(newModelVersionV2Request)
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
func TestClient_RetrieveTheDatasetListV2(t *testing.T) {
	retrieveTheDatasetListV2Request := &RetrieveTheDatasetListV2Request{
		Keyword:          util.PtrString(""),
		StorageType:      util.PtrString(""),
		StorageInstances: util.PtrString(""),
		ImportFormat:     util.PtrString(""),
		PageNumber:       util.PtrInt32(int32(0)),
		PageSize:         util.PtrInt32(int32(0)),
	}
	result := &RetrieveTheDatasetListV2Response{}
	result, err := AIHC_CLIENT.RetrieveTheDatasetListV2(retrieveTheDatasetListV2Request)
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
func TestClient_RetrieveTheDatasetVersionListV2(t *testing.T) {
	retrieveTheDatasetVersionListV2Request := &RetrieveTheDatasetVersionListV2Request{
		DatasetId:  util.PtrString(""),
		PageNumber: util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &RetrieveTheDatasetVersionListV2Response{}
	result, err := AIHC_CLIENT.RetrieveTheDatasetVersionListV2(retrieveTheDatasetVersionListV2Request)
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
