package oos

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
	OOS_CLIENT *Client
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

	OOS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CheckTemplateV2(t *testing.T) {
	checkTemplateV2Request := &CheckTemplateV2Request{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Tags:        []*KeyValuePair{},
		Linear:      util.PtrBool(false),
		Parallelism: util.PtrInt32(int32(0)),
		Operators:   []*Operator{},
		Links:       []*LinkModel{},
		Properties:  []*Property{},
	}
	result := &CheckTemplateV2Response{}
	result, err := OOS_CLIENT.CheckTemplateV2(checkTemplateV2Request)
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
func TestClient_CreateExecutionV2(t *testing.T) {
	Template := &Template{
		Id:                     util.PtrString(""),
		Ref:                    util.PtrString(""),
		Name:                   util.PtrString(""),
		OosType:                util.PtrString(""),
		Description:            util.PtrString(""),
		Tags:                   []*KeyValuePair{},
		Linear:                 util.PtrBool(false),
		Parallelism:            util.PtrInt32(int32(0)),
		Operators:              []*Operator{},
		Links:                  []*LinkModel{},
		Properties:             []*Property{},
		UpdatedTime:            util.PtrString(""),
		SupportedInstanceTypes: []*string{},
	}
	Properties := make(map[string]interface{})
	createExecutionV2Request := &CreateExecutionV2Request{
		Locale:      util.PtrString(""),
		Description: util.PtrString(""),
		Template:    Template,
		Parallelism: util.PtrInt32(int32(0)),
		Manually:    util.PtrBool(false),
		Properties:  nil,
		Tags:        []*Tag{},
	}
	result := &CreateExecutionV2Response{}
	result, err := OOS_CLIENT.CreateExecutionV2(createExecutionV2Request)
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
func TestClient_CreateTemplateV2(t *testing.T) {
	createTemplateV2Request := &CreateTemplateV2Request{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Tags:        []*KeyValuePair{},
		Linear:      util.PtrBool(false),
		Parallelism: util.PtrInt32(int32(0)),
		Operators:   []*Operator{},
		Links:       []*LinkModel{},
		Properties:  []*Property{},
	}
	result := &CreateTemplateV2Response{}
	result, err := OOS_CLIENT.CreateTemplateV2(createTemplateV2Request)
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
func TestClient_DeleteTemplateV2(t *testing.T) {
	deleteTemplateV2Request := &DeleteTemplateV2Request{
		Id: util.PtrString(""),
	}
	result := &DeleteTemplateV2Response{}
	result, err := OOS_CLIENT.DeleteTemplateV2(deleteTemplateV2Request)
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
func TestClient_GetExecutionDetailV2(t *testing.T) {
	getExecutionDetailV2Request := &GetExecutionDetailV2Request{
		Id:      util.PtrString(""),
		WithLog: util.PtrString(""),
		Locale:  util.PtrString(""),
	}
	result := &GetExecutionDetailV2Response{}
	result, err := OOS_CLIENT.GetExecutionDetailV2(getExecutionDetailV2Request)
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
func TestClient_GetExecutionListV2(t *testing.T) {
	Template := &TemplateFilter{
		Name: util.PtrString(""),
	}
	getExecutionListV2Request := &GetExecutionListV2Request{
		Locale:             util.PtrString(""),
		Template:           Template,
		State:              util.PtrString(""),
		Trigger:            util.PtrString(""),
		CronExecutionName:  util.PtrString(""),
		EventExecutionName: util.PtrString(""),
		StartTime:          util.PtrInt32(int32(0)),
		EndTime:            util.PtrInt32(int32(0)),
		Sort:               util.PtrString(""),
		Ascending:          util.PtrBool(false),
		PageNo:             util.PtrInt32(int32(0)),
		PageSize:           util.PtrInt32(int32(0)),
	}
	result := &GetExecutionListV2Response{}
	result, err := OOS_CLIENT.GetExecutionListV2(getExecutionListV2Request)
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
func TestClient_GetOperatorListV2(t *testing.T) {
	Operator := &OperatorFilter{
		Name: util.PtrString(""),
	}
	getOperatorListV2Request := &GetOperatorListV2Request{
		Locale:    util.PtrString(""),
		Operator:  Operator,
		Sort:      util.PtrString(""),
		Ascending: util.PtrBool(false),
		PageNo:    util.PtrInt32(int32(0)),
		PageSize:  util.PtrInt32(int32(0)),
	}
	result := &GetOperatorListV2Response{}
	result, err := OOS_CLIENT.GetOperatorListV2(getOperatorListV2Request)
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
func TestClient_GetTaskChildrenListV2(t *testing.T) {
	getTaskChildrenListV2Request := &GetTaskChildrenListV2Request{
		Locale:      util.PtrString(""),
		ExecutionId: util.PtrString(""),
		TaskId:      util.PtrString(""),
		States:      []*string{},
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &GetTaskChildrenListV2Response{}
	result, err := OOS_CLIENT.GetTaskChildrenListV2(getTaskChildrenListV2Request)
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
func TestClient_GetTaskDetailV2(t *testing.T) {
	getTaskDetailV2Request := &GetTaskDetailV2Request{
		DagId:          util.PtrString(""),
		TaskId:         util.PtrString(""),
		IgnoreChildren: util.PtrString(""),
		Locale:         util.PtrString(""),
	}
	result := &GetTaskDetailV2Response{}
	result, err := OOS_CLIENT.GetTaskDetailV2(getTaskDetailV2Request)
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
func TestClient_GetTemplateDetailV2(t *testing.T) {
	getTemplateDetailV2Request := &GetTemplateDetailV2Request{
		Id:     util.PtrString(""),
		Name:   util.PtrString(""),
		Type:   util.PtrString(""),
		Locale: util.PtrString(""),
	}
	result := &GetTemplateDetailV2Response{}
	result, err := OOS_CLIENT.GetTemplateDetailV2(getTemplateDetailV2Request)
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
func TestClient_GetTemplateListV2(t *testing.T) {
	getTemplateListV2Request := &GetTemplateListV2Request{
		Locale:                util.PtrString(""),
		Name:                  util.PtrString(""),
		Id:                    util.PtrString(""),
		OosType:               util.PtrString(""),
		Sort:                  util.PtrString(""),
		Ascending:             util.PtrBool(false),
		PageNo:                util.PtrInt32(int32(0)),
		PageSize:              util.PtrInt32(int32(0)),
		SupportedInstanceType: util.PtrString(""),
	}
	result := &GetTemplateListV2Response{}
	result, err := OOS_CLIENT.GetTemplateListV2(getTemplateListV2Request)
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
func TestClient_UpdateTemplateV2(t *testing.T) {
	updateTemplateV2Request := &UpdateTemplateV2Request{
		Id:          util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Tags:        []*KeyValuePair{},
		Linear:      util.PtrBool(false),
		Parallelism: util.PtrInt32(int32(0)),
		Operators:   []*Operator{},
		Links:       []*LinkModel{},
		Properties:  []*Property{},
	}
	result := &UpdateTemplateV2Response{}
	result, err := OOS_CLIENT.UpdateTemplateV2(updateTemplateV2Request)
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
