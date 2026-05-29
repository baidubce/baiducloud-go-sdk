package bls

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
	BLS_CLIENT *Client
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

	BLS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateProject(t *testing.T) {
	createProjectRequest := &CreateProjectRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateProjectResponse{}
	result, err := BLS_CLIENT.CreateProject(createProjectRequest)
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
func TestClient_DeleteProject(t *testing.T) {
	deleteProjectRequest := &DeleteProjectRequest{
		Uuid: util.PtrString(""),
	}
	result := &DeleteProjectResponse{}
	result, err := BLS_CLIENT.DeleteProject(deleteProjectRequest)
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
func TestClient_DescribeProject(t *testing.T) {
	describeProjectRequest := &DescribeProjectRequest{
		Uuid: util.PtrString(""),
	}
	result := &DescribeProjectResponse{}
	result, err := BLS_CLIENT.DescribeProject(describeProjectRequest)
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
func TestClient_ListProject(t *testing.T) {
	listProjectRequest := &ListProjectRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		OrderBy:     util.PtrString(""),
		Order:       util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &ListProjectResponse{}
	result, err := BLS_CLIENT.ListProject(listProjectRequest)
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
func TestClient_PullLogRecord(t *testing.T) {
	pullLogRecordRequest := &PullLogRecordRequest{
		LogStoreName:  util.PtrString(""),
		LogStreamName: util.PtrString(""),
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		Project:       util.PtrString(""),
		Limit:         util.PtrInt32(int32(0)),
		Marker:        util.PtrString(""),
	}
	result := &PullLogRecordResponse{}
	result, err := BLS_CLIENT.PullLogRecord(pullLogRecordRequest)
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
func TestClient_PushLogRecord(t *testing.T) {
	pushLogRecordRequest := &PushLogRecordRequest{
		LogStoreName:  util.PtrString(""),
		Project:       util.PtrString(""),
		LogStreamName: util.PtrString(""),
		BlsType:       util.PtrString(""),
		LogRecords:    []*LogRecord{},
		Tags:          []*LogTag{},
	}
	result := &PushLogRecordResponse{}
	result, err := BLS_CLIENT.PushLogRecord(pushLogRecordRequest)
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
func TestClient_QueryLogHistogram(t *testing.T) {
	queryLogHistogramRequest := &QueryLogHistogramRequest{
		LogStoreName:  util.PtrString(""),
		Query:         util.PtrString(""),
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		Project:       util.PtrString(""),
		LogStreamName: util.PtrString(""),
	}
	result := &QueryLogHistogramResponse{}
	result, err := BLS_CLIENT.QueryLogHistogram(queryLogHistogramRequest)
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
func TestClient_QueryLogRecord(t *testing.T) {
	queryLogRecordRequest := &QueryLogRecordRequest{
		LogStoreName:  util.PtrString(""),
		Query:         util.PtrString(""),
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		Project:       util.PtrString(""),
		LogStreamName: util.PtrString(""),
		Marker:        util.PtrString(""),
		Limit:         util.PtrInt32(int32(0)),
		Sort:          util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &QueryLogRecordResponse{}
	result, err := BLS_CLIENT.QueryLogRecord(queryLogRecordRequest)
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
func TestClient_UpdateProject(t *testing.T) {
	updateProjectRequest := &UpdateProjectRequest{
		Uuid:        util.PtrString(""),
		Description: util.PtrBool(false),
	}
	result := &UpdateProjectResponse{}
	result, err := BLS_CLIENT.UpdateProject(updateProjectRequest)
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
