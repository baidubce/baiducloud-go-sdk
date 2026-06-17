package bcm

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
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
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

func TestClient_AddAlarmPolicyActions(t *testing.T) {
	addAlarmPolicyActionsRequest := &AddAlarmPolicyActionsRequest{
		Id:      util.PtrString(""),
		Actions: []*PolicyAction{},
	}
	result := &AddAlarmPolicyActionsResponse{}
	result, err := BCM_CLIENT.AddAlarmPolicyActions(addAlarmPolicyActionsRequest)
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
func TestClient_CreateAlarmMasking(t *testing.T) {
	createAlarmMaskingRequest := &CreateAlarmMaskingRequest{
		Name:                util.PtrString(""),
		Scope:               util.PtrString(""),
		ResourceType:        util.PtrString(""),
		PolicyId:            util.PtrString(""),
		Instances:           []*TargetInstance{},
		Region:              util.PtrString(""),
		MetricNames:         []*string{},
		PeriodType:          util.PtrString(""),
		BeginTime:           util.PtrString(""),
		EndTime:             util.PtrString(""),
		Tz:                  util.PtrString(""),
		DailyBeginTimestamp: util.PtrInt32(int32(0)),
		DailyEndTimestamp:   util.PtrInt32(int32(0)),
	}
	result := &CreateAlarmMaskingResponse{}
	result, err := BCM_CLIENT.CreateAlarmMasking(createAlarmMaskingRequest)
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
func TestClient_CreateAlarmPolicy(t *testing.T) {
	Target := &AlarmTarget{
		BcmType:             util.PtrString(""),
		Instances:           []*TargetInstance{},
		Region:              util.PtrString(""),
		Tags:                []*Dimension{},
		InstanceGroups:      []*string{},
		IncludingDimensions: []*string{},
		ExcludingDimensions: []*string{},
	}
	createAlarmPolicyRequest := &CreateAlarmPolicyRequest{
		Name:                       util.PtrString(""),
		Scope:                      util.PtrString(""),
		ResourceType:               util.PtrString(""),
		Target:                     Target,
		Rules:                      []*AlarmRule{},
		PendingCount:               util.PtrInt32(int32(0)),
		OnMissingData:              util.PtrString(""),
		NoDataNotifyPendingMinutes: util.PtrInt32(int32(0)),
		BcmType:                    util.PtrString(""),
		Level:                      util.PtrString(""),
		Actions:                    []*PolicyAction{},
		NotifyEnabled:              util.PtrBool(false),
		Callbacks:                  []*Callback{},
		RenotifyCount:              util.PtrInt32(int32(0)),
		RenotifyIntervalMinutes:    util.PtrInt32(int32(0)),
		NotifyMergeWindowSeconds:   util.PtrInt32(int32(0)),
	}
	result := &CreateAlarmPolicyResponse{}
	result, err := BCM_CLIENT.CreateAlarmPolicy(createAlarmPolicyRequest)
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
func TestClient_DeleteAlarmMaskings(t *testing.T) {
	deleteAlarmMaskingsRequest := &DeleteAlarmMaskingsRequest{
		Ids: []*string{},
	}
	result := &DeleteAlarmMaskingsResponse{}
	result, err := BCM_CLIENT.DeleteAlarmMaskings(deleteAlarmMaskingsRequest)
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
func TestClient_DeleteAlarmPolicies(t *testing.T) {
	deleteAlarmPoliciesRequest := &DeleteAlarmPoliciesRequest{
		Ids: []*string{},
	}
	result := &DeleteAlarmPoliciesResponse{}
	result, err := BCM_CLIENT.DeleteAlarmPolicies(deleteAlarmPoliciesRequest)
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
func TestClient_DeleteAlarmPolicyActions(t *testing.T) {
	deleteAlarmPolicyActionsRequest := &DeleteAlarmPolicyActionsRequest{
		Id:      util.PtrString(""),
		Actions: []*PolicyAction{},
	}
	result := &DeleteAlarmPolicyActionsResponse{}
	result, err := BCM_CLIENT.DeleteAlarmPolicyActions(deleteAlarmPolicyActionsRequest)
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
func TestClient_DescribeAlarm(t *testing.T) {
	describeAlarmRequest := &DescribeAlarmRequest{
		Id: util.PtrString(""),
	}
	result := &DescribeAlarmResponse{}
	result, err := BCM_CLIENT.DescribeAlarm(describeAlarmRequest)
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
func TestClient_DescribeAlarmMasking(t *testing.T) {
	describeAlarmMaskingRequest := &DescribeAlarmMaskingRequest{
		Id: util.PtrString(""),
	}
	result := &DescribeAlarmMaskingResponse{}
	result, err := BCM_CLIENT.DescribeAlarmMasking(describeAlarmMaskingRequest)
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
func TestClient_DescribeAlarmMaskings(t *testing.T) {
	describeAlarmMaskingsRequest := &DescribeAlarmMaskingsRequest{
		MaskingName: util.PtrString(""),
		MaskingId:   util.PtrString(""),
		Order:       util.PtrString(""),
		OrderBy:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &DescribeAlarmMaskingsResponse{}
	result, err := BCM_CLIENT.DescribeAlarmMaskings(describeAlarmMaskingsRequest)
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
func TestClient_DescribeAlarmPolicies(t *testing.T) {
	describeAlarmPoliciesRequest := &DescribeAlarmPoliciesRequest{
		PolicyName:      util.PtrString(""),
		PolicyId:        util.PtrString(""),
		Scope:           util.PtrString(""),
		ResourceType:    util.PtrString(""),
		Recursive:       util.PtrBool(false),
		SubResourceType: util.PtrString(""),
		NotifyEnabled:   util.PtrBool(false),
		BcmType:         util.PtrString(""),
		Order:           util.PtrString(""),
		OrderBy:         util.PtrString(""),
		PageNo:          util.PtrInt32(int32(0)),
		PageSize:        util.PtrInt32(int32(0)),
	}
	result := &DescribeAlarmPoliciesResponse{}
	result, err := BCM_CLIENT.DescribeAlarmPolicies(describeAlarmPoliciesRequest)
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
func TestClient_DescribeAlarmPolicy(t *testing.T) {
	describeAlarmPolicyRequest := &DescribeAlarmPolicyRequest{
		Id:                     util.PtrString(""),
		RequireSubResourceType: util.PtrBool(false),
	}
	result := &DescribeAlarmPolicyResponse{}
	result, err := BCM_CLIENT.DescribeAlarmPolicy(describeAlarmPolicyRequest)
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
func TestClient_DescribeAlarms(t *testing.T) {
	describeAlarmsRequest := &DescribeAlarmsRequest{
		StartTime:    util.PtrString(""),
		EndTime:      util.PtrString(""),
		PolicyName:   util.PtrString(""),
		Scope:        util.PtrString(""),
		ResourceType: util.PtrString(""),
		State:        util.PtrString(""),
		BcmType:      util.PtrString(""),
		Order:        util.PtrString(""),
		OrderBy:      util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &DescribeAlarmsResponse{}
	result, err := BCM_CLIENT.DescribeAlarms(describeAlarmsRequest)
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
func TestClient_DescribeDimensionValues(t *testing.T) {
	describeDimensionValuesRequest := &DescribeDimensionValuesRequest{
		Action:        util.PtrString(""),
		Scope:         util.PtrString(""),
		ResourceType:  util.PtrString(""),
		Region:        util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		MetricName:    util.PtrString(""),
		DimensionKey:  util.PtrString(""),
		Filters:       []*Filter{},
	}
	result := &DescribeDimensionValuesResponse{}
	result, err := BCM_CLIENT.DescribeDimensionValues(describeDimensionValuesRequest)
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
func TestClient_DescribeMetricData(t *testing.T) {
	describeMetricDataRequest := &DescribeMetricDataRequest{
		Action:              util.PtrString(""),
		Scope:               util.PtrString(""),
		ResourceType:        util.PtrString(""),
		Region:              util.PtrString(""),
		BeginDatetime:       util.PtrString(""),
		EndDatetime:         util.PtrString(""),
		MetricName:          util.PtrString(""),
		Filters:             []*Filter{},
		Limit:               util.PtrInt32(int32(0)),
		Offset:              util.PtrInt32(int32(0)),
		PeriodSeconds:       util.PtrInt32(int32(0)),
		AggregationOverTime: []*string{},
	}
	result := &DescribeMetricDataResponse{}
	result, err := BCM_CLIENT.DescribeMetricData(describeMetricDataRequest)
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
func TestClient_DescribeMetricDataLatest(t *testing.T) {
	describeMetricDataLatestRequest := &DescribeMetricDataLatestRequest{
		Action:              util.PtrString(""),
		Scope:               util.PtrString(""),
		ResourceType:        util.PtrString(""),
		Region:              util.PtrString(""),
		EndDatetime:         util.PtrString(""),
		MetricName:          util.PtrString(""),
		Filters:             []*Filter{},
		Limit:               util.PtrInt32(int32(0)),
		Offset:              util.PtrInt32(int32(0)),
		PeriodSeconds:       util.PtrInt32(int32(0)),
		AggregationOverTime: []*string{},
	}
	result := &DescribeMetricDataLatestResponse{}
	result, err := BCM_CLIENT.DescribeMetricDataLatest(describeMetricDataLatestRequest)
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
func TestClient_DescribeMetricDataLatestTop(t *testing.T) {
	describeMetricDataLatestTopRequest := &DescribeMetricDataLatestTopRequest{
		Action:              util.PtrString(""),
		Scope:               util.PtrString(""),
		Region:              util.PtrString(""),
		EndDatetime:         util.PtrString(""),
		MetricName:          util.PtrString(""),
		Filters:             []*Filter{},
		Limit:               util.PtrInt32(int32(0)),
		Asc:                 util.PtrBool(false),
		PeriodSeconds:       util.PtrInt32(int32(0)),
		AggregationOverTime: util.PtrString(""),
	}
	result := &DescribeMetricDataLatestTopResponse{}
	result, err := BCM_CLIENT.DescribeMetricDataLatestTop(describeMetricDataLatestTopRequest)
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
func TestClient_UpdateAlarmMasking(t *testing.T) {
	updateAlarmMaskingRequest := &UpdateAlarmMaskingRequest{
		Id:                  util.PtrString(""),
		State:               util.PtrString(""),
		Name:                util.PtrString(""),
		Scope:               util.PtrString(""),
		ResourceType:        util.PtrString(""),
		PolicyId:            util.PtrString(""),
		Instances:           []*TargetInstance{},
		Region:              util.PtrString(""),
		MetricNames:         []*string{},
		PeriodType:          util.PtrString(""),
		BeginTime:           util.PtrString(""),
		EndTime:             util.PtrString(""),
		Tz:                  util.PtrString(""),
		DailyBeginTimestamp: util.PtrInt32(int32(0)),
		DailyEndTimestamp:   util.PtrInt32(int32(0)),
	}
	result := &UpdateAlarmMaskingResponse{}
	result, err := BCM_CLIENT.UpdateAlarmMasking(updateAlarmMaskingRequest)
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
func TestClient_UpdateAlarmMaskingStates(t *testing.T) {
	updateAlarmMaskingStatesRequest := &UpdateAlarmMaskingStatesRequest{
		Ids:   []*string{},
		State: util.PtrString(""),
	}
	result := &UpdateAlarmMaskingStatesResponse{}
	result, err := BCM_CLIENT.UpdateAlarmMaskingStates(updateAlarmMaskingStatesRequest)
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
func TestClient_UpdateAlarmPolicy(t *testing.T) {
	Target := &AlarmTarget{
		BcmType:             util.PtrString(""),
		Instances:           []*TargetInstance{},
		Region:              util.PtrString(""),
		Tags:                []*Dimension{},
		InstanceGroups:      []*string{},
		IncludingDimensions: []*string{},
		ExcludingDimensions: []*string{},
	}
	updateAlarmPolicyRequest := &UpdateAlarmPolicyRequest{
		Id:                         util.PtrString(""),
		State:                      util.PtrString(""),
		Name:                       util.PtrString(""),
		Scope:                      util.PtrString(""),
		ResourceType:               util.PtrString(""),
		Target:                     Target,
		Rules:                      []*AlarmRule{},
		PendingCount:               util.PtrInt32(int32(0)),
		OnMissingData:              util.PtrString(""),
		NoDataNotifyPendingMinutes: util.PtrInt32(int32(0)),
		BcmType:                    util.PtrString(""),
		Level:                      util.PtrString(""),
		Actions:                    []*PolicyAction{},
		NotifyEnabled:              util.PtrBool(false),
		Callbacks:                  []*Callback{},
		RenotifyCount:              util.PtrInt32(int32(0)),
		RenotifyIntervalMinutes:    util.PtrInt32(int32(0)),
		NotifyMergeWindowSeconds:   util.PtrInt32(int32(0)),
	}
	result := &UpdateAlarmPolicyResponse{}
	result, err := BCM_CLIENT.UpdateAlarmPolicy(updateAlarmPolicyRequest)
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
func TestClient_UpdateAlarmPolicyNotifyEnabled(t *testing.T) {
	updateAlarmPolicyNotifyEnabledRequest := &UpdateAlarmPolicyNotifyEnabledRequest{
		Ids:           []*string{},
		NotifyEnabled: util.PtrBool(false),
	}
	result := &UpdateAlarmPolicyNotifyEnabledResponse{}
	result, err := BCM_CLIENT.UpdateAlarmPolicyNotifyEnabled(updateAlarmPolicyNotifyEnabledRequest)
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
func TestClient_UpdateAlarmPolicyState(t *testing.T) {
	updateAlarmPolicyStateRequest := &UpdateAlarmPolicyStateRequest{
		Ids:   []*string{},
		State: util.PtrString(""),
	}
	result := &UpdateAlarmPolicyStateResponse{}
	result, err := BCM_CLIENT.UpdateAlarmPolicyState(updateAlarmPolicyStateRequest)
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
