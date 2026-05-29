package cfs

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
	CFS_CLIENT *Client
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

	CFS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_BatchCreationOfPermissionGroupRules(t *testing.T) {
	batchCreationOfPermissionGroupRulesRequest := &BatchCreationOfPermissionGroupRulesRequest{
		AgName:      util.PtrString(""),
		AccessRules: []*RuleInfo{},
	}
	result := &BatchCreationOfPermissionGroupRulesResponse{}
	result, err := CFS_CLIENT.BatchCreationOfPermissionGroupRules(batchCreationOfPermissionGroupRulesRequest)
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
func TestClient_CreateFileSystem(t *testing.T) {
	createFileSystemRequest := &CreateFileSystemRequest{
		FsName:        util.PtrString(""),
		Zone:          util.PtrString(""),
		CfsType:       util.PtrString(""),
		Protocol:      util.PtrString(""),
		Tags:          []*Tag{},
		CapacityQuota: util.PtrInt64(int64(0)),
	}
	result := &CreateFileSystemResponse{}
	result, err := CFS_CLIENT.CreateFileSystem(createFileSystemRequest)
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
func TestClient_CreateMountingTarget(t *testing.T) {
	createMountingTargetRequest := &CreateMountingTargetRequest{
		FsId:            util.PtrString(""),
		SubnetId:        util.PtrString(""),
		VpcId:           util.PtrString(""),
		AccessGroupName: util.PtrString(""),
		Address:         util.PtrString(""),
	}
	result := &CreateMountingTargetResponse{}
	result, err := CFS_CLIENT.CreateMountingTarget(createMountingTargetRequest)
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
func TestClient_CreatePermissionGroup(t *testing.T) {
	createPermissionGroupRequest := &CreatePermissionGroupRequest{
		AccessGroupName: util.PtrString(""),
		Description:     util.PtrString(""),
	}
	err := CFS_CLIENT.CreatePermissionGroup(createPermissionGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreatePermissionGroupRules(t *testing.T) {
	createPermissionGroupRulesRequest := &CreatePermissionGroupRulesRequest{
		AgName:      util.PtrString(""),
		Ip:          util.PtrString(""),
		Mask:        util.PtrInt32(int32(0)),
		WriteAccess: util.PtrString(""),
		UserAccess:  util.PtrString(""),
		Priority:    util.PtrInt32(int32(0)),
	}
	result := &CreatePermissionGroupRulesResponse{}
	result, err := CFS_CLIENT.CreatePermissionGroupRules(createPermissionGroupRulesRequest)
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
func TestClient_DeletePermissionGroup(t *testing.T) {
	deletePermissionGroupRequest := &DeletePermissionGroupRequest{
		AgName: util.PtrString(""),
	}
	err := CFS_CLIENT.DeletePermissionGroup(deletePermissionGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeletePermissionGroupRule(t *testing.T) {
	deletePermissionGroupRuleRequest := &DeletePermissionGroupRuleRequest{
		AgName: util.PtrString(""),
		ArId:   util.PtrString(""),
	}
	err := CFS_CLIENT.DeletePermissionGroupRule(deletePermissionGroupRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DropFileSystem(t *testing.T) {
	dropFileSystemRequest := &DropFileSystemRequest{
		FsId: util.PtrString(""),
	}
	err := CFS_CLIENT.DropFileSystem(dropFileSystemRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DropMountTarget(t *testing.T) {
	dropMountTargetRequest := &DropMountTargetRequest{
		FsId:    util.PtrString(""),
		MountId: util.PtrString(""),
	}
	err := CFS_CLIENT.DropMountTarget(dropMountTargetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTheMountTargetPermissionGroup(t *testing.T) {
	modifyTheMountTargetPermissionGroupRequest := &ModifyTheMountTargetPermissionGroupRequest{
		FsId:            util.PtrString(""),
		MountID:         util.PtrString(""),
		AccessGroupName: util.PtrString(""),
	}
	err := CFS_CLIENT.ModifyTheMountTargetPermissionGroup(modifyTheMountTargetPermissionGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryFileSystem(t *testing.T) {
	queryFileSystemRequest := &QueryFileSystemRequest{
		UserId:    util.PtrString(""),
		FsId:      util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		FilterTag: util.PtrString(""),
	}
	result := &QueryFileSystemResponse{}
	result, err := CFS_CLIENT.QueryFileSystem(queryFileSystemRequest)
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
func TestClient_QueryMountedClient(t *testing.T) {
	queryMountedClientRequest := &QueryMountedClientRequest{
		FsId:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryMountedClientResponse{}
	result, err := CFS_CLIENT.QueryMountedClient(queryMountedClientRequest)
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
func TestClient_QueryMountingTarget(t *testing.T) {
	queryMountingTargetRequest := &QueryMountingTargetRequest{
		FId:     util.PtrString(""),
		MountId: util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryMountingTargetResponse{}
	result, err := CFS_CLIENT.QueryMountingTarget(queryMountingTargetRequest)
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
func TestClient_QueryPermissionGroup(t *testing.T) {
	queryPermissionGroupRequest := &QueryPermissionGroupRequest{
		AgName:  util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryPermissionGroupResponse{}
	result, err := CFS_CLIENT.QueryPermissionGroup(queryPermissionGroupRequest)
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
func TestClient_QueryPermissionGroupRules(t *testing.T) {
	queryPermissionGroupRulesRequest := &QueryPermissionGroupRulesRequest{
		AgName:  util.PtrString(""),
		ArId:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryPermissionGroupRulesResponse{}
	result, err := CFS_CLIENT.QueryPermissionGroupRules(queryPermissionGroupRulesRequest)
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
func TestClient_UpdateFileSystem(t *testing.T) {
	updateFileSystemRequest := &UpdateFileSystemRequest{
		FsId:          util.PtrString(""),
		FsName:        util.PtrString(""),
		CapacityQuota: util.PtrInt32(int32(0)),
	}
	err := CFS_CLIENT.UpdateFileSystem(updateFileSystemRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateFileSystemLabels(t *testing.T) {
	updateFileSystemLabelsRequest := &UpdateFileSystemLabelsRequest{
		FsId: []*string{},
		Tags: []*Tag{},
	}
	err := CFS_CLIENT.UpdateFileSystemLabels(updateFileSystemLabelsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePermissionGroup(t *testing.T) {
	updatePermissionGroupRequest := &UpdatePermissionGroupRequest{
		AgName:      util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := CFS_CLIENT.UpdatePermissionGroup(updatePermissionGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePermissionGroupRules(t *testing.T) {
	updatePermissionGroupRulesRequest := &UpdatePermissionGroupRulesRequest{
		AgName:      util.PtrString(""),
		ArId:        util.PtrString(""),
		Ip:          util.PtrString(""),
		Mask:        util.PtrInt32(int32(0)),
		WriteAccess: util.PtrString(""),
		UserAccess:  util.PtrString(""),
		Priority:    util.PtrInt32(int32(0)),
	}
	err := CFS_CLIENT.UpdatePermissionGroupRules(updatePermissionGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
