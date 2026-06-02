package pfs

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
	PFS_CLIENT *Client
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

	PFS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreatePfs(t *testing.T) {
	createPfsRequest := &CreatePfsRequest{
		Name:         util.PtrString(""),
		InstanceType: util.PtrString(""),
		Capacity:     util.PtrInt32(int32(0)),
		SubnetId:     util.PtrString(""),
		Description:  util.PtrString(""),
		Tags:         []*Tag{},
	}
	result := &CreatePfsResponse{}
	result, err := PFS_CLIENT.CreatePfs(createPfsRequest)
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
func TestClient_DeletePfs(t *testing.T) {
	deletePfsRequest := &DeletePfsRequest{
		InstanceId: util.PtrString(""),
	}
	err := PFS_CLIENT.DeletePfs(deletePfsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescPfs(t *testing.T) {
	descPfsRequest := &DescPfsRequest{
		InstanceId: util.PtrString(""),
	}
	result := &DescPfsResponse{}
	result, err := PFS_CLIENT.DescPfs(descPfsRequest)
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
func TestClient_InstanceListClients(t *testing.T) {
	instanceListClientsRequest := &InstanceListClientsRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Manner:     util.PtrString(""),
		Marker:     util.PtrString(""),
	}
	result := &InstanceListClientsResponse{}
	result, err := PFS_CLIENT.InstanceListClients(instanceListClientsRequest)
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
func TestClient_ListPfs(t *testing.T) {
	listPfsRequest := &ListPfsRequest{
		Manner:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		Marker:    util.PtrString(""),
		FilterTag: util.PtrString(""),
	}
	result := &ListPfsResponse{}
	result, err := PFS_CLIENT.ListPfs(listPfsRequest)
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
func TestClient_MountTargetListClients(t *testing.T) {
	mountTargetListClientsRequest := &MountTargetListClientsRequest{
		Action:        util.PtrString(""),
		MountTargetId: util.PtrString(""),
		MaxKeys:       util.PtrInt32(int32(0)),
		Manner:        util.PtrString(""),
		Marker:        util.PtrString(""),
	}
	result := &MountTargetListClientsResponse{}
	result, err := PFS_CLIENT.MountTargetListClients(mountTargetListClientsRequest)
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
func TestClient_UpdatePFSTag(t *testing.T) {
	updatePFSTagRequest := &UpdatePFSTagRequest{
		InstanceId: []*string{},
		Tags:       []*Tag{},
	}
	err := PFS_CLIENT.UpdatePFSTag(updatePFSTagRequest)
	ExpectEqual(t.Errorf, nil, err)
}
