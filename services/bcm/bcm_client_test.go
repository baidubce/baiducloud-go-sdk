package bcm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	BCM_CLIENT *Client
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
		BCM_CLIENT, _ = NewClient("ak", "sk", DEFAULT_ENDPOINT)
		log.SetLogLevel(log.WARN)
		return
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	BCM_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateAlarmTemplate(t *testing.T) {
	createAlarmTemplateRequest := &CreateAlarmTemplateRequest{
		Scope:           util.PtrString(""),
		ResourceType:    util.PtrString(""),
		SubResourceType: util.PtrString(""),
		Name:            util.PtrString(""),
		Comment:         util.PtrString(""),
		Rules:           []*AlarmRule{},
	}
	result := &CreateAlarmTemplateResponse{}
	result, err := BCM_CLIENT.CreateAlarmTemplate(createAlarmTemplateRequest)
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
func TestClient_CreateInstanceGroup(t *testing.T) {
	createInstanceGroupRequest := &CreateInstanceGroupRequest{
		Scope:        util.PtrString(""),
		ResourceType: util.PtrString(""),
		Name:         util.PtrString(""),
		Instances:    []*InstanceGroupInstance{},
	}
	result := &CreateInstanceGroupResponse{}
	result, err := BCM_CLIENT.CreateInstanceGroup(createInstanceGroupRequest)
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
func TestClient_CreateNotifyTemplate(t *testing.T) {
	createNotifyTemplateRequest := &CreateNotifyTemplateRequest{
		Name:           util.PtrString(""),
		SilencePeriods: []*SilencePeriod{},
		Receivers:      []*NotifyReceiver{},
		Callbacks:      []*Callback{},
	}
	result := &CreateNotifyTemplateResponse{}
	result, err := BCM_CLIENT.CreateNotifyTemplate(createNotifyTemplateRequest)
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
func TestClient_DeleteAlarmTemplates(t *testing.T) {
	deleteAlarmTemplatesRequest := &DeleteAlarmTemplatesRequest{
		Ids: []*string{},
	}
	result := &DeleteAlarmTemplatesResponse{}
	result, err := BCM_CLIENT.DeleteAlarmTemplates(deleteAlarmTemplatesRequest)
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
func TestClient_DeleteInstanceGroup(t *testing.T) {
	deleteInstanceGroupRequest := &DeleteInstanceGroupRequest{
		Id: util.PtrString(""),
	}
	result := &DeleteInstanceGroupResponse{}
	result, err := BCM_CLIENT.DeleteInstanceGroup(deleteInstanceGroupRequest)
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
func TestClient_DeleteInstanceGroupInstances(t *testing.T) {
	deleteInstanceGroupInstancesRequest := &DeleteInstanceGroupInstancesRequest{
		Id:        util.PtrString(""),
		Instances: []*InstanceGroupInstance{},
	}
	result := &DeleteInstanceGroupInstancesResponse{}
	result, err := BCM_CLIENT.DeleteInstanceGroupInstances(deleteInstanceGroupInstancesRequest)
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
func TestClient_DeleteNotifyTemplate(t *testing.T) {
	deleteNotifyTemplateRequest := &DeleteNotifyTemplateRequest{
		NotifyId: util.PtrString(""),
	}
	result := &DeleteNotifyTemplateResponse{}
	result, err := BCM_CLIENT.DeleteNotifyTemplate(deleteNotifyTemplateRequest)
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
func TestClient_DescribeAlarmTemplate(t *testing.T) {
	describeAlarmTemplateRequest := &DescribeAlarmTemplateRequest{
		Id: util.PtrString(""),
	}
	result := &DescribeAlarmTemplateResponse{}
	result, err := BCM_CLIENT.DescribeAlarmTemplate(describeAlarmTemplateRequest)
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
func TestClient_DescribeAlarmTemplates(t *testing.T) {
	describeAlarmTemplatesRequest := &DescribeAlarmTemplatesRequest{
		Name:            util.PtrString(""),
		Scope:           util.PtrString(""),
		ResourceType:    util.PtrString(""),
		SubResourceType: util.PtrString(""),
		Order:           util.PtrString(""),
		OrderBy:         util.PtrString(""),
		PageNo:          util.PtrInt32(int32(0)),
		PageSize:        util.PtrInt32(int32(0)),
	}
	result := &DescribeAlarmTemplatesResponse{}
	result, err := BCM_CLIENT.DescribeAlarmTemplates(describeAlarmTemplatesRequest)
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
func TestClient_DescribeInstanceGroup(t *testing.T) {
	describeInstanceGroupRequest := &DescribeInstanceGroupRequest{
		Id: util.PtrString(""),
	}
	result := &DescribeInstanceGroupResponse{}
	result, err := BCM_CLIENT.DescribeInstanceGroup(describeInstanceGroupRequest)
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
func TestClient_DescribeInstanceGroups(t *testing.T) {
	describeInstanceGroupsRequest := &DescribeInstanceGroupsRequest{
		Scope:        util.PtrString(""),
		ResourceType: util.PtrString(""),
		Name:         util.PtrString(""),
		Order:        util.PtrString(""),
		OrderBy:      util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &DescribeInstanceGroupsResponse{}
	result, err := BCM_CLIENT.DescribeInstanceGroups(describeInstanceGroupsRequest)
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
func TestClient_DescribeNotifyTemplate(t *testing.T) {
	describeNotifyTemplateRequest := &DescribeNotifyTemplateRequest{
		Id: util.PtrString(""),
	}
	result := &DescribeNotifyTemplateResponse{}
	result, err := BCM_CLIENT.DescribeNotifyTemplate(describeNotifyTemplateRequest)
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
func TestClient_DescribeNotifyTemplates(t *testing.T) {
	describeNotifyTemplatesRequest := &DescribeNotifyTemplatesRequest{
		Name:     util.PtrString(""),
		PolicyId: util.PtrString(""),
		Order:    util.PtrString(""),
		OrderBy:  util.PtrString(""),
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &DescribeNotifyTemplatesResponse{}
	result, err := BCM_CLIENT.DescribeNotifyTemplates(describeNotifyTemplatesRequest)
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
func TestClient_DescribeReceivers(t *testing.T) {
	describeReceiversRequest := &DescribeReceiversRequest{
		BcmType:  util.PtrString(""),
		Name:     util.PtrString(""),
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &DescribeReceiversResponse{}
	result, err := BCM_CLIENT.DescribeReceivers(describeReceiversRequest)
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
func TestClient_DescribeSystemTemplateRules(t *testing.T) {
	describeSystemTemplateRulesRequest := &DescribeSystemTemplateRulesRequest{
		Scope:           util.PtrString(""),
		ResourceType:    util.PtrString(""),
		SubResourceType: util.PtrString(""),
		Source:          util.PtrString(""),
	}
	result := &DescribeSystemTemplateRulesResponse{}
	result, err := BCM_CLIENT.DescribeSystemTemplateRules(describeSystemTemplateRulesRequest)
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
func TestClient_ExportAlarmTemplates(t *testing.T) {
	exportAlarmTemplatesRequest := &ExportAlarmTemplatesRequest{
		Names: []*string{},
	}
	result := &ExportAlarmTemplatesResponse{}
	result, err := BCM_CLIENT.ExportAlarmTemplates(exportAlarmTemplatesRequest)
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
func TestClient_ImportAlarmTemplates(t *testing.T) {
	importAlarmTemplatesRequest := &ImportAlarmTemplatesRequest{
		Overwrite: util.PtrBool(false),
		Templates: []*Template{},
	}
	result := &ImportAlarmTemplatesResponse{}
	result, err := BCM_CLIENT.ImportAlarmTemplates(importAlarmTemplatesRequest)
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
func TestClient_UpdateAlarmTemplate(t *testing.T) {
	updateAlarmTemplateRequest := &UpdateAlarmTemplateRequest{
		Id:              util.PtrString(""),
		Scope:           util.PtrString(""),
		ResourceType:    util.PtrString(""),
		SubResourceType: util.PtrString(""),
		Name:            util.PtrString(""),
		Comment:         util.PtrString(""),
		Rules:           []*AlarmRule{},
	}
	result := &UpdateAlarmTemplateResponse{}
	result, err := BCM_CLIENT.UpdateAlarmTemplate(updateAlarmTemplateRequest)
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
func TestClient_UpdateInstanceGroup(t *testing.T) {
	updateInstanceGroupRequest := &UpdateInstanceGroupRequest{
		Id:           util.PtrString(""),
		Scope:        util.PtrString(""),
		ResourceType: util.PtrString(""),
		Name:         util.PtrString(""),
		Instances:    []*InstanceGroupInstance{},
	}
	result := &UpdateInstanceGroupResponse{}
	result, err := BCM_CLIENT.UpdateInstanceGroup(updateInstanceGroupRequest)
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
func TestClient_UpdateNotifyTemplate(t *testing.T) {
	updateNotifyTemplateRequest := &UpdateNotifyTemplateRequest{
		Id:             util.PtrString(""),
		Name:           util.PtrString(""),
		SilencePeriods: []*SilencePeriod{},
		Receivers:      []*NotifyReceiver{},
		Callbacks:      []*Callback{},
	}
	result := &UpdateNotifyTemplateResponse{}
	result, err := BCM_CLIENT.UpdateNotifyTemplate(updateNotifyTemplateRequest)
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

func TestClient_BCMTemplateAndInstanceGroupAPIs(t *testing.T) {
	tests := []struct {
		name         string
		action       string
		request      string
		responseBody string
		call         func(*Client) (interface{}, error)
		assertResult func(*testing.T, interface{})
	}{
		{
			name:         "create alarm template sends request and parses success",
			action:       "CreateAlarmTemplate",
			request:      `"name":"cpu-high"`,
			responseBody: `{"success":true,"code":"OK","message":"created"}`,
			call: func(client *Client) (interface{}, error) {
				return client.CreateAlarmTemplate(&CreateAlarmTemplateRequest{
					Scope:           util.PtrString("scope-a"),
					ResourceType:    util.PtrString("BCC"),
					SubResourceType: util.PtrString("instance"),
					Name:            util.PtrString("cpu-high"),
					Comment:         util.PtrString("cpu alarm"),
					Rules: []*AlarmRule{{
						PendingCount:         util.PtrInt32(3),
						CheckIntervalSeconds: util.PtrInt32(60),
						Content:              util.PtrString("cpu too high"),
						Conditions: []*AlarmCondition{{
							MetricName:       util.PtrString("CpuPercent"),
							Operator:         util.PtrString(">"),
							Threshold:        util.PtrFloat32(80),
							AggregateType:    util.PtrString("average"),
							WindowSeconds:    util.PtrInt32(60),
							DisplayUnit:      util.PtrString("%"),
							DisplayThreshold: util.PtrString("80"),
						}},
					}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*CreateAlarmTemplateResponse)
				assertStringPtr(t, got.Code, "OK")
				assertStringPtr(t, got.Message, "created")
				assertBoolPtr(t, got.Success, true)
			},
		},
		{
			name:         "create instance group sends request and parses id",
			action:       "CreateInstanceGroup",
			request:      `"name":"prod-group"`,
			responseBody: `{"success":true,"code":"OK","message":"created","id":"group-001"}`,
			call: func(client *Client) (interface{}, error) {
				return client.CreateInstanceGroup(&CreateInstanceGroupRequest{
					Scope:        util.PtrString("scope-a"),
					ResourceType: util.PtrString("BCC"),
					Name:         util.PtrString("prod-group"),
					Instances: []*InstanceGroupInstance{{
						Region: util.PtrString("bj"),
					}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*CreateInstanceGroupResponse)
				assertStringPtr(t, got.Id, "group-001")
				assertBoolPtr(t, got.Success, true)
			},
		},
		{
			name:         "create notify template sends nested receiver data",
			action:       "CreateNotifyTemplate",
			request:      `"receiverType":"USER"`,
			responseBody: `{"success":true,"code":"OK","message":"notify-created","notifyId":"notify-001"}`,
			call: func(client *Client) (interface{}, error) {
				return client.CreateNotifyTemplate(&CreateNotifyTemplateRequest{
					Name: util.PtrString("notify-template"),
					SilencePeriods: []*SilencePeriod{{
						From: util.PtrString("00:00"),
						To:   util.PtrString("01:00"),
					}},
					Receivers: []*NotifyReceiver{{
						Id:           util.PtrString("receiver-001"),
						Name:         util.PtrString("owner"),
						ReceiverType: util.PtrString("USER"),
						Channels:     []*string{util.PtrString("EMAIL")},
					}},
					Callbacks: []*Callback{{
						Url: util.PtrString("https://example.com/callback"),
						Mention: &Mention{
							BcmType: util.PtrString("USER"),
							UserIds: []*string{util.PtrString("user-001")},
						},
					}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*CreateNotifyTemplateResponse)
				assertStringPtr(t, got.NotifyId, "notify-001")
				assertStringPtr(t, got.Message, "notify-created")
			},
		},
		{
			name:         "delete alarm templates sends ids",
			action:       "DeleteAlarmTemplates",
			request:      `"ids":["alarm-001","alarm-002"]`,
			responseBody: `{"success":true,"code":"OK","message":"deleted"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DeleteAlarmTemplates(&DeleteAlarmTemplatesRequest{
					Ids: []*string{util.PtrString("alarm-001"), util.PtrString("alarm-002")},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DeleteAlarmTemplatesResponse)
				assertStringPtr(t, got.Message, "deleted")
			},
		},
		{
			name:         "delete instance group sends id",
			action:       "DeleteInstanceGroup",
			request:      `"id":"group-001"`,
			responseBody: `{"success":true,"code":"OK","message":"deleted"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DeleteInstanceGroup(&DeleteInstanceGroupRequest{Id: util.PtrString("group-001")})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DeleteInstanceGroupResponse)
				assertBoolPtr(t, got.Success, true)
			},
		},
		{
			name:         "delete instance group instances sends instances",
			action:       "DeleteInstanceGroupInstances",
			request:      `"region":"bj"`,
			responseBody: `{"success":true,"code":"OK","message":"instances-deleted"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DeleteInstanceGroupInstances(&DeleteInstanceGroupInstancesRequest{
					Id:        util.PtrString("group-001"),
					Instances: []*InstanceGroupInstance{{Region: util.PtrString("bj")}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DeleteInstanceGroupInstancesResponse)
				assertStringPtr(t, got.Message, "instances-deleted")
			},
		},
		{
			name:         "delete notify template sends notify id",
			action:       "DeleteNotifyTemplate",
			request:      `"notifyId":"notify-001"`,
			responseBody: `{"success":true,"code":"OK","message":"notify-deleted"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DeleteNotifyTemplate(&DeleteNotifyTemplateRequest{NotifyId: util.PtrString("notify-001")})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DeleteNotifyTemplateResponse)
				assertStringPtr(t, got.Message, "notify-deleted")
			},
		},
		{
			name:         "describe alarm template sends id",
			action:       "DescribeAlarmTemplate",
			request:      `"id":"alarm-001"`,
			responseBody: `{"success":true,"code":"OK","message":"found","id":"alarm-001","name":"cpu-high"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeAlarmTemplate(&DescribeAlarmTemplateRequest{Id: util.PtrString("alarm-001")})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeAlarmTemplateResponse)
				assertStringPtr(t, got.Name, "cpu-high")
				assertStringPtr(t, got.Id, "alarm-001")
			},
		},
		{
			name:         "describe alarm templates sends pagination",
			action:       "DescribeAlarmTemplates",
			request:      `"pageSize":20`,
			responseBody: `{"success":true,"code":"OK","message":"listed","alarmTemplates":[{"name":"cpu-high"}],"pageNo":1,"pageSize":20,"totalCount":1}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeAlarmTemplates(&DescribeAlarmTemplatesRequest{
					Name:            util.PtrString("cpu"),
					Scope:           util.PtrString("scope-a"),
					ResourceType:    util.PtrString("BCC"),
					SubResourceType: util.PtrString("instance"),
					Order:           util.PtrString("desc"),
					OrderBy:         util.PtrString("name"),
					PageNo:          util.PtrInt32(1),
					PageSize:        util.PtrInt32(20),
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeAlarmTemplatesResponse)
				assertInt32Ptr(t, got.TotalCount, 1)
				assertStringPtr(t, got.AlarmTemplates[0].Name, "cpu-high")
			},
		},
		{
			name:         "describe instance group sends id",
			action:       "DescribeInstanceGroup",
			request:      `"id":"group-001"`,
			responseBody: `{"success":true,"code":"OK","message":"found","id":"group-001","name":"prod-group"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeInstanceGroup(&DescribeInstanceGroupRequest{Id: util.PtrString("group-001")})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeInstanceGroupResponse)
				assertStringPtr(t, got.Id, "group-001")
				assertStringPtr(t, got.Name, "prod-group")
			},
		},
		{
			name:         "describe instance groups sends filters",
			action:       "DescribeInstanceGroups",
			request:      `"resourceType":"BCC"`,
			responseBody: `{"success":true,"code":"OK","message":"listed","instanceGroups":[{"id":"group-001","name":"prod-group"}],"pageNo":1,"pageSize":10,"totalCount":1}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeInstanceGroups(&DescribeInstanceGroupsRequest{
					Scope:        util.PtrString("scope-a"),
					ResourceType: util.PtrString("BCC"),
					Name:         util.PtrString("prod"),
					Order:        util.PtrString("asc"),
					OrderBy:      util.PtrString("name"),
					PageNo:       util.PtrInt32(1),
					PageSize:     util.PtrInt32(10),
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeInstanceGroupsResponse)
				assertInt32Ptr(t, got.PageSize, 10)
				assertStringPtr(t, got.InstanceGroups[0].Name, "prod-group")
			},
		},
		{
			name:         "describe notify template sends id",
			action:       "DescribeNotifyTemplate",
			request:      `"id":"notify-001"`,
			responseBody: `{"success":true,"code":"OK","message":"found","id":"notify-001","name":"notify-template"}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeNotifyTemplate(&DescribeNotifyTemplateRequest{Id: util.PtrString("notify-001")})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeNotifyTemplateResponse)
				assertStringPtr(t, got.Id, "notify-001")
				assertStringPtr(t, got.Name, "notify-template")
			},
		},
		{
			name:         "describe notify templates sends policy filter",
			action:       "DescribeNotifyTemplates",
			request:      `"policyId":"policy-001"`,
			responseBody: `{"success":true,"code":"OK","message":"listed","notifyTemplates":[{"id":"notify-001","name":"notify-template"}]}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeNotifyTemplates(&DescribeNotifyTemplatesRequest{
					Name:     util.PtrString("notify"),
					PolicyId: util.PtrString("policy-001"),
					Order:    util.PtrString("desc"),
					OrderBy:  util.PtrString("name"),
					PageNo:   util.PtrInt32(1),
					PageSize: util.PtrInt32(5),
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeNotifyTemplatesResponse)
				assertStringPtr(t, got.NotifyTemplates[0].Id, "notify-001")
				assertStringPtr(t, got.NotifyTemplates[0].Name, "notify-template")
			},
		},
		{
			name:         "describe receivers sends receiver type",
			action:       "DescribeReceivers",
			request:      `"type":"USER"`,
			responseBody: `{"success":true,"code":"OK","message":"listed","receivers":[{"receivers.id":"receiver-001","receivers.name":"owner"}]}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeReceivers(&DescribeReceiversRequest{
					BcmType:  util.PtrString("USER"),
					Name:     util.PtrString("owner"),
					PageNo:   util.PtrInt32(2),
					PageSize: util.PtrInt32(10),
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeReceiversResponse)
				assertStringPtr(t, got.Receivers[0].ReceiversId, "receiver-001")
				assertStringPtr(t, got.Receivers[0].ReceiversName, "owner")
			},
		},
		{
			name:         "describe system template rules sends dimensions",
			action:       "DescribeSystemTemplateRules",
			request:      `"source":"system"`,
			responseBody: `{"success":true,"code":"OK","message":"listed","rules":[{"content":"default rule"}]}`,
			call: func(client *Client) (interface{}, error) {
				return client.DescribeSystemTemplateRules(&DescribeSystemTemplateRulesRequest{
					Scope:           util.PtrString("scope-a"),
					ResourceType:    util.PtrString("BCC"),
					SubResourceType: util.PtrString("instance"),
					Source:          util.PtrString("system"),
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*DescribeSystemTemplateRulesResponse)
				assertStringPtr(t, got.Rules[0].Content, "default rule")
			},
		},
		{
			name:         "export alarm templates sends names",
			action:       "ExportAlarmTemplates",
			request:      `"names":["cpu-high"]`,
			responseBody: `{"success":true,"code":"OK","message":"exported","templates":[{"name":"cpu-high"}]}`,
			call: func(client *Client) (interface{}, error) {
				return client.ExportAlarmTemplates(&ExportAlarmTemplatesRequest{Names: []*string{util.PtrString("cpu-high")}})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*ExportAlarmTemplatesResponse)
				assertStringPtr(t, got.Templates[0].Name, "cpu-high")
			},
		},
		{
			name:         "import alarm templates sends overwrite flag",
			action:       "ImportAlarmTemplates",
			request:      `"overwrite":true`,
			responseBody: `{"success":true,"code":"OK","message":"imported"}`,
			call: func(client *Client) (interface{}, error) {
				return client.ImportAlarmTemplates(&ImportAlarmTemplatesRequest{
					Overwrite: util.PtrBool(true),
					Templates: []*Template{{
						Scope:        util.PtrString("scope-a"),
						ResourceType: util.PtrString("BCC"),
						Name:         util.PtrString("cpu-high"),
					}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*ImportAlarmTemplatesResponse)
				assertStringPtr(t, got.Message, "imported")
			},
		},
		{
			name:         "update alarm template sends id and rules",
			action:       "UpdateAlarmTemplate",
			request:      `"id":"alarm-001"`,
			responseBody: `{"success":true,"code":"OK","message":"updated"}`,
			call: func(client *Client) (interface{}, error) {
				return client.UpdateAlarmTemplate(&UpdateAlarmTemplateRequest{
					Id:           util.PtrString("alarm-001"),
					Scope:        util.PtrString("scope-a"),
					ResourceType: util.PtrString("BCC"),
					Name:         util.PtrString("cpu-high-updated"),
					Rules:        []*AlarmRule{{Content: util.PtrString("updated rule")}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*UpdateAlarmTemplateResponse)
				assertStringPtr(t, got.Message, "updated")
			},
		},
		{
			name:         "update instance group sends id and instances",
			action:       "UpdateInstanceGroup",
			request:      `"id":"group-001"`,
			responseBody: `{"success":true,"code":"OK","message":"group-updated"}`,
			call: func(client *Client) (interface{}, error) {
				return client.UpdateInstanceGroup(&UpdateInstanceGroupRequest{
					Id:           util.PtrString("group-001"),
					Scope:        util.PtrString("scope-a"),
					ResourceType: util.PtrString("BCC"),
					Name:         util.PtrString("prod-group"),
					Instances:    []*InstanceGroupInstance{{Region: util.PtrString("gz")}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*UpdateInstanceGroupResponse)
				assertStringPtr(t, got.Message, "group-updated")
			},
		},
		{
			name:         "update notify template sends callback",
			action:       "UpdateNotifyTemplate",
			request:      `"url":"https://example.com/update"`,
			responseBody: `{"success":true,"code":"OK","message":"notify-updated"}`,
			call: func(client *Client) (interface{}, error) {
				return client.UpdateNotifyTemplate(&UpdateNotifyTemplateRequest{
					Id:        util.PtrString("notify-001"),
					Name:      util.PtrString("notify-updated"),
					Callbacks: []*Callback{{Url: util.PtrString("https://example.com/update")}},
				})
			},
			assertResult: func(t *testing.T, result interface{}) {
				got := result.(*UpdateNotifyTemplateResponse)
				assertStringPtr(t, got.Message, "notify-updated")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, receivedBody, closeServer := newBCMTestClient(t, tt.action, tt.responseBody)
			defer closeServer()

			result, err := tt.call(client)
			if err != nil {
				t.Fatalf("%s returned error = %v, want nil", tt.action, err)
			}
			if !strings.Contains(*receivedBody, tt.request) {
				t.Fatalf("%s request body = %s, want contains %s", tt.action, *receivedBody, tt.request)
			}
			tt.assertResult(t, result)
		})
	}
}

func TestClient_BCMTemplateAPIPropagatesServiceError(t *testing.T) {
	client, receivedBody, closeServer := newBCMErrorTestClient(t, "CreateAlarmTemplate")
	defer closeServer()

	result, err := client.CreateAlarmTemplate(&CreateAlarmTemplateRequest{Name: util.PtrString("invalid-template")})
	if err == nil {
		t.Fatal("CreateAlarmTemplate error = nil, want service error")
	}
	if result != nil {
		t.Fatalf("CreateAlarmTemplate result = %#v, want nil on service error", result)
	}
	if !strings.Contains(err.Error(), "InvalidTemplate") {
		t.Fatalf("CreateAlarmTemplate error = %v, want contains InvalidTemplate", err)
	}
	if !strings.Contains(*receivedBody, `"name":"invalid-template"`) {
		t.Fatalf("CreateAlarmTemplate request body = %s, want invalid-template", *receivedBody)
	}
}

func newBCMTestClient(t *testing.T, expectedAction string, responseBody string) (*Client, *string, func()) {
	t.Helper()
	var receivedBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertBCMHTTPRequest(t, r, expectedAction)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body error = %v, want nil", err)
		}
		receivedBody = string(body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	client, err := NewClient("ak", "sk", server.URL)
	if err != nil {
		server.Close()
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	return client, &receivedBody, server.Close
}

func newBCMErrorTestClient(t *testing.T, expectedAction string) (*Client, *string, func()) {
	t.Helper()
	var receivedBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertBCMHTTPRequest(t, r, expectedAction)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body error = %v, want nil", err)
		}
		receivedBody = string(body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"code":"InvalidTemplate","message":"template name is invalid"}`))
	}))
	client, err := NewClient("ak", "sk", server.URL)
	if err != nil {
		server.Close()
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	return client, &receivedBody, server.Close
}

func assertBCMHTTPRequest(t *testing.T, r *http.Request, expectedAction string) {
	t.Helper()
	if r.Method != http.MethodPost {
		t.Fatalf("method = %s, want %s", r.Method, http.MethodPost)
	}
	if r.URL.Path != "/v3/bcm" {
		t.Fatalf("path = %s, want /v3/bcm", r.URL.Path)
	}
	if got := r.URL.Query().Get("action"); got != expectedAction {
		t.Fatalf("action = %s, want %s", got, expectedAction)
	}
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json;charset=utf-8" {
		t.Fatalf("Content-Type = %s, want application/json;charset=utf-8", contentType)
	}
}

func assertStringPtr(t *testing.T, got *string, want string) {
	t.Helper()
	if got == nil {
		t.Fatalf("string pointer = nil, want %s", want)
	}
	if *got != want {
		t.Fatalf("string pointer = %s, want %s", *got, want)
	}
}

func assertBoolPtr(t *testing.T, got *bool, want bool) {
	t.Helper()
	if got == nil {
		t.Fatalf("bool pointer = nil, want %t", want)
	}
	if *got != want {
		t.Fatalf("bool pointer = %t, want %t", *got, want)
	}
}

func assertInt32Ptr(t *testing.T, got *int32, want int32) {
	t.Helper()
	if got == nil {
		t.Fatalf("int32 pointer = nil, want %d", want)
	}
	if *got != want {
		t.Fatalf("int32 pointer = %d, want %d", *got, want)
	}
}
