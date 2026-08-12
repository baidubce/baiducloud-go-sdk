package ax

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
	AX_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
	ApiKey   string
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
	// AX_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	// ==== API Key 鉴权 ====
	AX_CLIENT, _ = NewClientWithApiKey(confObj.ApiKey, confObj.Endpoint)

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

func TestClient_BatchReleaseSandboxes(t *testing.T) {
	batchReleaseSandboxesRequest := &BatchReleaseSandboxesRequest{
		SandboxIds: []*string{},
	}
	result := &BatchReleaseSandboxesResponse{}
	result, err := AX_CLIENT.BatchReleaseSandboxes(batchReleaseSandboxesRequest)
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
func TestClient_ConnectSandbox(t *testing.T) {
	connectSandboxRequest := &ConnectSandboxRequest{
		SandboxID:  util.PtrString(""),
		Timeout:    util.PtrInt32(int32(0)),
		SnapshotID: util.PtrString(""),
	}
	result := &ConnectSandboxResponse{}
	result, err := AX_CLIENT.ConnectSandbox(connectSandboxRequest)
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
func TestClient_CreateSandbox(t *testing.T) {
	Metadata := make(map[string]string)
	EnvVars := make(map[string]string)
	AutoResume := make(map[string]interface{})
	Mcp := make(map[string]interface{})
	createSandboxRequest := &CreateSandboxRequest{
		TemplateID:          util.PtrString(""),
		Timeout:             util.PtrInt32(int32(0)),
		Metadata:            nil,
		EnvVars:             nil,
		Secure:              util.PtrBool(false),
		AllowInternetAccess: util.PtrBool(false),
		AutoPause:           util.PtrBool(false),
		AutoResume:          nil,
		RuntimeType:         util.PtrString(""),
		Mcp:                 nil,
		VolumeMounts:        []*map[string]interface{}{},
	}
	result := &CreateSandboxResponse{}
	result, err := AX_CLIENT.CreateSandbox(createSandboxRequest)
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
func TestClient_CreateSandboxSnapshot(t *testing.T) {
	createSandboxSnapshotRequest := &CreateSandboxSnapshotRequest{
		SandboxID: util.PtrString(""),
		Name:      util.PtrString(""),
	}
	result := &CreateSandboxSnapshotResponse{}
	result, err := AX_CLIENT.CreateSandboxSnapshot(createSandboxSnapshotRequest)
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
func TestClient_DeleteSandbox(t *testing.T) {
	deleteSandboxRequest := &DeleteSandboxRequest{
		SandboxID: util.PtrString(""),
	}
	err := AX_CLIENT.DeleteSandbox(deleteSandboxRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ForkSandbox(t *testing.T) {
	forkSandboxRequest := &ForkSandboxRequest{
		SandboxID: util.PtrString(""),
	}
	result := &ForkSandboxResponse{}
	result, err := AX_CLIENT.ForkSandbox(forkSandboxRequest)
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
func TestClient_GetSandbox(t *testing.T) {
	getSandboxRequest := &GetSandboxRequest{
		SandboxID: util.PtrString(""),
	}
	result := &GetSandboxResponse{}
	result, err := AX_CLIENT.GetSandbox(getSandboxRequest)
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
func TestClient_GetSandboxResources(t *testing.T) {
	getSandboxResourcesRequest := &GetSandboxResourcesRequest{
		SandboxID: util.PtrString(""),
	}
	result := &GetSandboxResourcesResponse{}
	result, err := AX_CLIENT.GetSandboxResources(getSandboxResourcesRequest)
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
func TestClient_GetSandboxSnapshot(t *testing.T) {
	getSandboxSnapshotRequest := &GetSandboxSnapshotRequest{
		SandboxID:  util.PtrString(""),
		SnapshotID: util.PtrString(""),
	}
	result := &GetSandboxSnapshotResponse{}
	result, err := AX_CLIENT.GetSandboxSnapshot(getSandboxSnapshotRequest)
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
func TestClient_ListSandboxSnapshots(t *testing.T) {
	listSandboxSnapshotsRequest := &ListSandboxSnapshotsRequest{
		SandboxID: util.PtrString(""),
	}
	result := &ListSandboxSnapshotsResponse{}
	result, err := AX_CLIENT.ListSandboxSnapshots(listSandboxSnapshotsRequest)
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
func TestClient_ListSandboxes(t *testing.T) {
	listSandboxesRequest := &ListSandboxesRequest{
		Metadata: util.PtrString(""),
	}
	result := &ListSandboxesResponse{}
	result, err := AX_CLIENT.ListSandboxes(listSandboxesRequest)
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
func TestClient_ListSandboxesV2(t *testing.T) {
	listSandboxesV2Request := &ListSandboxesV2Request{
		Limit:     util.PtrInt32(int32(0)),
		NextToken: util.PtrString(""),
		Metadata:  util.PtrString(""),
		State:     util.PtrString(""),
	}
	result := &ListSandboxesV2Response{}
	result, err := AX_CLIENT.ListSandboxesV2(listSandboxesV2Request)
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
func TestClient_ListSandboxesV2ByPath(t *testing.T) {
	listSandboxesV2ByPathRequest := &ListSandboxesV2ByPathRequest{
		Limit:     util.PtrInt32(int32(0)),
		NextToken: util.PtrString(""),
		Metadata:  util.PtrString(""),
		State:     util.PtrString(""),
	}
	result := &ListSandboxesV2ByPathResponse{}
	result, err := AX_CLIENT.ListSandboxesV2ByPath(listSandboxesV2ByPathRequest)
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
func TestClient_PauseSandbox(t *testing.T) {
	pauseSandboxRequest := &PauseSandboxRequest{
		SandboxID:     util.PtrString(""),
		HibernateMode: util.PtrString(""),
	}
	err := AX_CLIENT.PauseSandbox(pauseSandboxRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySandboxes(t *testing.T) {
	Metadata := make(map[string]string)
	querySandboxesRequest := &QuerySandboxesRequest{
		Limit:      util.PtrInt32(int32(0)),
		NextToken:  util.PtrString(""),
		SandboxIds: []*string{},
		ImagePaths: []*string{},
		Metadata:   nil,
		State:      []*string{},
	}
	result := &QuerySandboxesResponse{}
	result, err := AX_CLIENT.QuerySandboxes(querySandboxesRequest)
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
func TestClient_ResumeSandbox(t *testing.T) {
	resumeSandboxRequest := &ResumeSandboxRequest{
		SandboxID: util.PtrString(""),
		Timeout:   util.PtrInt32(int32(0)),
		AutoPause: util.PtrBool(false),
	}
	result := &ResumeSandboxResponse{}
	result, err := AX_CLIENT.ResumeSandbox(resumeSandboxRequest)
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
func TestClient_SetSandboxTimeout(t *testing.T) {
	setSandboxTimeoutRequest := &SetSandboxTimeoutRequest{
		SandboxID: util.PtrString(""),
		Timeout:   util.PtrInt32(int32(0)),
	}
	err := AX_CLIENT.SetSandboxTimeout(setSandboxTimeoutRequest)
	ExpectEqual(t.Errorf, nil, err)
}
