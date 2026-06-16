package cfw

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
	CFW_CLIENT *Client
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

	CFW_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_BindCfw(t *testing.T) {
	bindCfwRequest := &BindCfwRequest{
		CfwId:        util.PtrString(""),
		InstanceType: util.PtrString(""),
		Instances:    []*CfwBind{},
	}
	err := CFW_CLIENT.BindCfw(bindCfwRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateCfw(t *testing.T) {
	createCfwRequest := &CreateCfwRequest{
		Name:        util.PtrString(""),
		CfwType:     util.PtrInt32(int32(0)),
		Border:      util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
		CfwRules:    []*CreateRule{},
	}
	result := &CreateCfwResponse{}
	result, err := CFW_CLIENT.CreateCfw(createCfwRequest)
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
func TestClient_CreateCfwRule(t *testing.T) {
	createCfwRuleRequest := &CreateCfwRuleRequest{
		CfwId:    util.PtrString(""),
		CfwRules: []*CreateRule{},
	}
	err := CFW_CLIENT.CreateCfwRule(createCfwRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateStatelessCfw(t *testing.T) {
	createStatelessCfwRequest := &CreateStatelessCfwRequest{
		Name:          util.PtrString(""),
		Description:   util.PtrString(""),
		DefaultAction: util.PtrString(""),
		Protocol:      util.PtrString(""),
		IpList:        []*string{},
	}
	result := &CreateStatelessCfwResponse{}
	result, err := CFW_CLIENT.CreateStatelessCfw(createStatelessCfwRequest)
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
func TestClient_DeleteCfw(t *testing.T) {
	deleteCfwRequest := &DeleteCfwRequest{
		CfwId: util.PtrString(""),
	}
	err := CFW_CLIENT.DeleteCfw(deleteCfwRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCfwRule(t *testing.T) {
	deleteCfwRuleRequest := &DeleteCfwRuleRequest{
		CfwId:      util.PtrString(""),
		CfwRuleIds: []*string{},
	}
	err := CFW_CLIENT.DeleteCfwRule(deleteCfwRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableCfwProtect(t *testing.T) {
	disableCfwProtectRequest := &DisableCfwProtectRequest{
		CfwId:      util.PtrString(""),
		InstanceId: util.PtrString(""),
		Role:       util.PtrString(""),
		MemberId:   util.PtrString(""),
	}
	err := CFW_CLIENT.DisableCfwProtect(disableCfwProtectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableCfwProtect(t *testing.T) {
	enableCfwProtectRequest := &EnableCfwProtectRequest{
		CfwId:      util.PtrString(""),
		InstanceId: util.PtrString(""),
		Role:       util.PtrString(""),
		MemberId:   util.PtrString(""),
	}
	err := CFW_CLIENT.EnableCfwProtect(enableCfwProtectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetCfw(t *testing.T) {
	getCfwRequest := &GetCfwRequest{
		CfwId: util.PtrString(""),
	}
	result := &GetCfwResponse{}
	result, err := CFW_CLIENT.GetCfw(getCfwRequest)
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
func TestClient_GetStatelessCfw(t *testing.T) {
	getStatelessCfwRequest := &GetStatelessCfwRequest{
		CfwId: util.PtrString(""),
	}
	result := &GetStatelessCfwResponse{}
	result, err := CFW_CLIENT.GetStatelessCfw(getStatelessCfwRequest)
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
func TestClient_ListCfw(t *testing.T) {
	listCfwRequest := &ListCfwRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListCfwResponse{}
	result, err := CFW_CLIENT.ListCfw(listCfwRequest)
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
func TestClient_ListProtectInstances(t *testing.T) {
	listProtectInstancesRequest := &ListProtectInstancesRequest{
		InstanceType: util.PtrString(""),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
		Status:       util.PtrString(""),
		Region:       util.PtrString(""),
	}
	result := &ListProtectInstancesResponse{}
	result, err := CFW_CLIENT.ListProtectInstances(listProtectInstancesRequest)
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
func TestClient_ListStatelessCfw(t *testing.T) {
	listStatelessCfwRequest := &ListStatelessCfwRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListStatelessCfwResponse{}
	result, err := CFW_CLIENT.ListStatelessCfw(listStatelessCfwRequest)
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
func TestClient_UnbindCfw(t *testing.T) {
	unbindCfwRequest := &UnbindCfwRequest{
		CfwId:        util.PtrString(""),
		InstanceType: util.PtrString(""),
		Instances:    []*CfwBind{},
	}
	err := CFW_CLIENT.UnbindCfw(unbindCfwRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateCfw(t *testing.T) {
	updateCfwRequest := &UpdateCfwRequest{
		CfwId:       util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := CFW_CLIENT.UpdateCfw(updateCfwRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateCfwRule(t *testing.T) {
	updateCfwRuleRequest := &UpdateCfwRuleRequest{
		CfwId:         util.PtrString(""),
		CfwRuleId:     util.PtrString(""),
		IpVersion:     util.PtrInt32(int32(0)),
		Priority:      util.PtrInt32(int32(0)),
		Protocol:      util.PtrString(""),
		Direction:     util.PtrString(""),
		SourceAddress: util.PtrString(""),
		DestAddress:   util.PtrString(""),
		SourcePort:    util.PtrString(""),
		DestPort:      util.PtrString(""),
		Action:        util.PtrString(""),
		Description:   util.PtrString(""),
	}
	err := CFW_CLIENT.UpdateCfwRule(updateCfwRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateStatelessCfw(t *testing.T) {
	updateStatelessCfwRequest := &UpdateStatelessCfwRequest{
		CfwId:       util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Protocol:    util.PtrString(""),
		IpList:      []*string{},
	}
	err := CFW_CLIENT.UpdateStatelessCfw(updateStatelessCfwRequest)
	ExpectEqual(t.Errorf, nil, err)
}
