package cloudassistant

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
	CLOUDASSISTANT_CLIENT *Client
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

	CLOUDASSISTANT_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_ActionList(t *testing.T) {
	Action := &ActionFilter{
		CloudassistantType: util.PtrString(""),
		Command: &CommandFilter{
			Scope:              util.PtrString(""),
			Name:               util.PtrString(""),
			CloudassistantType: util.PtrString(""),
		},
		InstanceType: util.PtrString(""),
		Id:           util.PtrString(""),
		Name:         util.PtrString(""),
		FileUpload:   nil,
	}
	actionListRequest := &ActionListRequest{
		Locale:    util.PtrString(""),
		PageNo:    util.PtrInt32(int32(0)),
		PageSize:  util.PtrInt32(int32(0)),
		Sort:      util.PtrString(""),
		Ascending: util.PtrBool(false),
		Action:    Action,
	}
	result := &ActionListResponse{}
	result, err := CLOUDASSISTANT_CLIENT.ActionList(actionListRequest)
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
func TestClient_ActionLog(t *testing.T) {
	actionLogRequest := &ActionLogRequest{
		RunId:   util.PtrString(""),
		ChildId: util.PtrString(""),
		Cursor:  util.PtrInt32(int32(0)),
	}
	result := &ActionLogResponse{}
	result, err := CLOUDASSISTANT_CLIENT.ActionLog(actionLogRequest)
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
func TestClient_ActionRun(t *testing.T) {
	TargetSelector := &TargetSelector{
		InstanceType: util.PtrString(""),
		Tags:         []*Tag{},
		ImportInstances: &TargetImport{
			KeywordType: util.PtrString(""),
			Instances:   []*string{},
		},
	}
	actionRunRequest := &ActionRunRequest{
		Locale:             util.PtrString(""),
		Action:             nil,
		Parameters:         nil,
		TargetSelectorType: util.PtrString(""),
		Targets:            []*interface{}{},
		TargetSelector:     TargetSelector,
	}
	result := &ActionRunResponse{}
	result, err := CLOUDASSISTANT_CLIENT.ActionRun(actionRunRequest)
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
func TestClient_ActionRunList(t *testing.T) {
	Action := &ActionFilter{
		CloudassistantType: util.PtrString(""),
		Command: &CommandFilter{
			Scope:              util.PtrString(""),
			Name:               util.PtrString(""),
			CloudassistantType: util.PtrString(""),
		},
		InstanceType: util.PtrString(""),
		Id:           util.PtrString(""),
		Name:         util.PtrString(""),
		FileUpload:   nil,
	}
	actionRunListRequest := &ActionRunListRequest{
		Locale:    util.PtrString(""),
		PageNo:    util.PtrInt32(int32(0)),
		PageSize:  util.PtrInt32(int32(0)),
		Sort:      util.PtrString(""),
		Ascending: util.PtrBool(false),
		Action:    Action,
		State:     util.PtrString(""),
		RunId:     util.PtrString(""),
		StartTime: util.PtrInt32(int32(0)),
		EndTime:   util.PtrInt32(int32(0)),
	}
	result := &ActionRunListResponse{}
	result, err := CLOUDASSISTANT_CLIENT.ActionRunList(actionRunListRequest)
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
func TestClient_BatchGetAgent(t *testing.T) {
	batchGetAgentRequest := &BatchGetAgentRequest{
		Hosts: []*Host{},
	}
	result := &BatchGetAgentResponse{}
	result, err := CLOUDASSISTANT_CLIENT.BatchGetAgent(batchGetAgentRequest)
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
func TestClient_CreateAction(t *testing.T) {
	Action := &Action{
		Id:                 util.PtrString(""),
		Ref:                util.PtrString(""),
		CloudassistantType: util.PtrString(""),
		Name:               util.PtrString(""),
		Alias:              util.PtrString(""),
		Description:        util.PtrString(""),
		TimeoutSecond:      util.PtrInt32(int32(0)),
		Command: &Command{
			CloudassistantType: util.PtrString(""),
			Content:            util.PtrString(""),
			Scope:              util.PtrString(""),
			EnableParameter:    util.PtrBool(false),
			Parameters:         []*Parameter{},
			User:               util.PtrString(""),
			WorkDir:            util.PtrString(""),
			ExecParams:         nil,
			WaitOnAgentMilli:   util.PtrInt32(int32(0)),
		},
		FileUpload: &FileUpload{
			Os:            util.PtrString(""),
			Content:       util.PtrString(""),
			Filename:      util.PtrString(""),
			Filepath:      util.PtrString(""),
			BosBucketName: util.PtrString(""),
			BosFilePath:   util.PtrString(""),
			BosEtag:       util.PtrString(""),
			User:          util.PtrString(""),
			Group:         util.PtrString(""),
			Mode:          util.PtrString(""),
			Overwrite:     util.PtrBool(false),
		},
		SupportedInstanceTypes: []*string{},
		CreatedTimestamp:       util.PtrInt64(int64(0)),
		UpdatedTimestamp:       util.PtrInt64(int64(0)),
	}
	TargetSelector := &TargetSelector{
		InstanceType: util.PtrString(""),
		Tags:         []*Tag{},
		ImportInstances: &TargetImport{
			KeywordType: util.PtrString(""),
			Instances:   []*string{},
		},
	}
	createActionRequest := &CreateActionRequest{
		Execution:          util.PtrString(""),
		Action:             Action,
		Parameters:         nil,
		TargetSelectorType: util.PtrString(""),
		Targets:            []*Target{},
		TargetSelector:     TargetSelector,
	}
	result := &CreateActionResponse{}
	result, err := CLOUDASSISTANT_CLIENT.CreateAction(createActionRequest)
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
func TestClient_DeleteAction(t *testing.T) {
	deleteActionRequest := &DeleteActionRequest{
		Id: util.PtrString(""),
	}
	result := &DeleteActionResponse{}
	result, err := CLOUDASSISTANT_CLIENT.DeleteAction(deleteActionRequest)
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
func TestClient_GetAction(t *testing.T) {
	getActionRequest := &GetActionRequest{
		Id:     util.PtrString(""),
		Locale: util.PtrString(""),
	}
	result := &GetActionResponse{}
	result, err := CLOUDASSISTANT_CLIENT.GetAction(getActionRequest)
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
func TestClient_GetActionRun(t *testing.T) {
	getActionRunRequest := &GetActionRunRequest{
		Id:            util.PtrString(""),
		WithLog:       util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
		ChildRunState: util.PtrString(""),
		Locale:        util.PtrString(""),
	}
	result := &GetActionRunResponse{}
	result, err := CLOUDASSISTANT_CLIENT.GetActionRun(getActionRunRequest)
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
func TestClient_UpdateAction(t *testing.T) {
	Action := &Action{
		Id:                 util.PtrString(""),
		Ref:                util.PtrString(""),
		CloudassistantType: util.PtrString(""),
		Name:               util.PtrString(""),
		Alias:              util.PtrString(""),
		Description:        util.PtrString(""),
		TimeoutSecond:      util.PtrInt32(int32(0)),
		Command: &Command{
			CloudassistantType: util.PtrString(""),
			Content:            util.PtrString(""),
			Scope:              util.PtrString(""),
			EnableParameter:    util.PtrBool(false),
			Parameters:         []*Parameter{},
			User:               util.PtrString(""),
			WorkDir:            util.PtrString(""),
			ExecParams:         nil,
			WaitOnAgentMilli:   util.PtrInt32(int32(0)),
		},
		FileUpload: &FileUpload{
			Os:            util.PtrString(""),
			Content:       util.PtrString(""),
			Filename:      util.PtrString(""),
			Filepath:      util.PtrString(""),
			BosBucketName: util.PtrString(""),
			BosFilePath:   util.PtrString(""),
			BosEtag:       util.PtrString(""),
			User:          util.PtrString(""),
			Group:         util.PtrString(""),
			Mode:          util.PtrString(""),
			Overwrite:     util.PtrBool(false),
		},
		SupportedInstanceTypes: []*string{},
		CreatedTimestamp:       util.PtrInt64(int64(0)),
		UpdatedTimestamp:       util.PtrInt64(int64(0)),
	}
	TargetSelector := &TargetSelector{
		InstanceType: util.PtrString(""),
		Tags:         []*Tag{},
		ImportInstances: &TargetImport{
			KeywordType: util.PtrString(""),
			Instances:   []*string{},
		},
	}
	updateActionRequest := &UpdateActionRequest{
		Execution:          util.PtrString(""),
		Action:             Action,
		Parameters:         nil,
		TargetSelectorType: util.PtrString(""),
		Targets:            []*Target{},
		TargetSelector:     TargetSelector,
	}
	result := &UpdateActionResponse{}
	result, err := CLOUDASSISTANT_CLIENT.UpdateAction(updateActionRequest)
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
