package apm

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
	APM_CLIENT *Client
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

	APM_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_ApmCreateAlarmPolicy(t *testing.T) {
	Target := &AlarmTarget{
		ApmType:  util.PtrString(""),
		Tags:     []*Tag{},
		Services: []*string{},
	}
	Rule := &AlarmRule{
		Operator:        util.PtrString(""),
		Rules:           []*AlarmRule{},
		Metric:          util.PtrString(""),
		WindowInSeconds: util.PtrInt32(int32(0)),
		Aggregate:       util.PtrString(""),
		DisplayValue:    util.PtrFloat32(float32(0)),
		DisplayUnit:     util.PtrString(""),
	}
	apmCreateAlarmPolicyRequest := &ApmCreateAlarmPolicyRequest{
		Name:                                 util.PtrString(""),
		State:                                util.PtrString(""),
		Target:                               Target,
		MetricKind:                           util.PtrString(""),
		Rule:                                 Rule,
		Filters:                              []*AlarmFilter{},
		PendingCount:                         util.PtrInt32(int32(0)),
		RenotifyIntervalInMinutes:            util.PtrInt32(int32(0)),
		RenotifyCount:                        util.PtrInt32(int32(0)),
		NotifyRecovery:                       util.PtrBool(false),
		OnMissingData:                        util.PtrString(""),
		NoDataNotifyPendingIntervalInMinutes: util.PtrInt32(int32(0)),
		Level:                                util.PtrString(""),
		Actions:                              []*AlarmAction{},
	}
	result := &ApmCreateAlarmPolicyResponse{}
	result, err := APM_CLIENT.ApmCreateAlarmPolicy(apmCreateAlarmPolicyRequest)
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
func TestClient_ApmDeleteAlarmPolicy(t *testing.T) {
	apmDeleteAlarmPolicyRequest := &ApmDeleteAlarmPolicyRequest{
		Ids: []*string{},
	}
	result := &ApmDeleteAlarmPolicyResponse{}
	result, err := APM_CLIENT.ApmDeleteAlarmPolicy(apmDeleteAlarmPolicyRequest)
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
func TestClient_ApmDescribeAlarm(t *testing.T) {
	apmDescribeAlarmRequest := &ApmDescribeAlarmRequest{
		Id: util.PtrString(""),
	}
	result := &ApmDescribeAlarmResponse{}
	result, err := APM_CLIENT.ApmDescribeAlarm(apmDescribeAlarmRequest)
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
func TestClient_ApmDescribeAlarmPolicies(t *testing.T) {
	apmDescribeAlarmPoliciesRequest := &ApmDescribeAlarmPoliciesRequest{
		PolicyName: util.PtrString(""),
		PolicyId:   util.PtrString(""),
		State:      util.PtrString(""),
		Level:      util.PtrString(""),
		MetricKind: util.PtrString(""),
		OrderBy:    util.PtrString(""),
		Order:      util.PtrString(""),
		PageNo:     util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &ApmDescribeAlarmPoliciesResponse{}
	result, err := APM_CLIENT.ApmDescribeAlarmPolicies(apmDescribeAlarmPoliciesRequest)
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
func TestClient_ApmDescribeAlarmPolicy(t *testing.T) {
	apmDescribeAlarmPolicyRequest := &ApmDescribeAlarmPolicyRequest{
		Id: util.PtrString(""),
	}
	result := &ApmDescribeAlarmPolicyResponse{}
	result, err := APM_CLIENT.ApmDescribeAlarmPolicy(apmDescribeAlarmPolicyRequest)
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
func TestClient_ApmDescribeAlarms(t *testing.T) {
	apmDescribeAlarmsRequest := &ApmDescribeAlarmsRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		PolicyName:    util.PtrString(""),
		Level:         util.PtrString(""),
		MetricKind:    util.PtrString(""),
		State:         util.PtrString(""),
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &ApmDescribeAlarmsResponse{}
	result, err := APM_CLIENT.ApmDescribeAlarms(apmDescribeAlarmsRequest)
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
func TestClient_ApmUpdateAlarmPolicy(t *testing.T) {
	Target := &AlarmTarget{
		ApmType:  util.PtrString(""),
		Tags:     []*Tag{},
		Services: []*string{},
	}
	Rule := &AlarmRule{
		Operator:        util.PtrString(""),
		Rules:           []*AlarmRule{},
		Metric:          util.PtrString(""),
		WindowInSeconds: util.PtrInt32(int32(0)),
		Aggregate:       util.PtrString(""),
		DisplayValue:    util.PtrFloat32(float32(0)),
		DisplayUnit:     util.PtrString(""),
	}
	apmUpdateAlarmPolicyRequest := &ApmUpdateAlarmPolicyRequest{
		Id:                                   util.PtrString(""),
		Name:                                 util.PtrString(""),
		Target:                               Target,
		MetricKind:                           util.PtrString(""),
		Rule:                                 Rule,
		Filters:                              []*AlarmFilter{},
		PendingCount:                         util.PtrInt32(int32(0)),
		RenotifyIntervalInMinutes:            util.PtrInt32(int32(0)),
		RenotifyCount:                        util.PtrInt32(int32(0)),
		NotifyRecovery:                       util.PtrBool(false),
		OnMissingData:                        util.PtrString(""),
		NoDataNotifyPendingIntervalInMinutes: util.PtrInt32(int32(0)),
		Level:                                util.PtrString(""),
		Actions:                              []*AlarmAction{},
	}
	result := &ApmUpdateAlarmPolicyResponse{}
	result, err := APM_CLIENT.ApmUpdateAlarmPolicy(apmUpdateAlarmPolicyRequest)
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
func TestClient_ApmUpdateAlarmPolicyAction(t *testing.T) {
	apmUpdateAlarmPolicyActionRequest := &ApmUpdateAlarmPolicyActionRequest{
		Id:      util.PtrString(""),
		Actions: []*AlarmAction{},
	}
	result := &ApmUpdateAlarmPolicyActionResponse{}
	result, err := APM_CLIENT.ApmUpdateAlarmPolicyAction(apmUpdateAlarmPolicyActionRequest)
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
func TestClient_ApmUpdateAlarmPolicyState(t *testing.T) {
	apmUpdateAlarmPolicyStateRequest := &ApmUpdateAlarmPolicyStateRequest{
		Ids:   []*string{},
		State: util.PtrString(""),
	}
	result := &ApmUpdateAlarmPolicyStateResponse{}
	result, err := APM_CLIENT.ApmUpdateAlarmPolicyState(apmUpdateAlarmPolicyStateRequest)
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
func TestClient_BindServiceTag(t *testing.T) {
	bindServiceTagRequest := &BindServiceTagRequest{
		ServiceName: util.PtrString(""),
		ServiceId:   util.PtrString(""),
		Tags:        []*Tag{},
	}
	result := &BindServiceTagResponse{}
	result, err := APM_CLIENT.BindServiceTag(bindServiceTagRequest)
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
func TestClient_DeleteServices(t *testing.T) {
	deleteServicesRequest := &DeleteServicesRequest{
		ServiceNames: []*string{},
	}
	result := &DeleteServicesResponse{}
	result, err := APM_CLIENT.DeleteServices(deleteServicesRequest)
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
func TestClient_DescribeDbStatement(t *testing.T) {
	describeDbStatementRequest := &DescribeDbStatementRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Service:       util.PtrString(""),
		Statements:    []*StatementQuery{},
	}
	result := &DescribeDbStatementResponse{}
	result, err := APM_CLIENT.DescribeDbStatement(describeDbStatementRequest)
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
func TestClient_DescribeDefaultConfig(t *testing.T) {
	result := &DescribeDefaultConfigResponse{}
	result, err := APM_CLIENT.DescribeDefaultConfig()
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
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		DimensionKey:  util.PtrString(""),
		Filters:       []*Filter{},
	}
	result := &DescribeDimensionValuesResponse{}
	result, err := APM_CLIENT.DescribeDimensionValues(describeDimensionValuesRequest)
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
func TestClient_DescribeEnv(t *testing.T) {
	result := &DescribeEnvResponse{}
	result, err := APM_CLIENT.DescribeEnv()
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
func TestClient_DescribeExceptions(t *testing.T) {
	describeExceptionsRequest := &DescribeExceptionsRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Service:       util.PtrString(""),
		Exceptions:    []*ExceptionQuery{},
	}
	result := &DescribeExceptionsResponse{}
	result, err := APM_CLIENT.DescribeExceptions(describeExceptionsRequest)
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
func TestClient_DescribeLLMDimensionValues(t *testing.T) {
	describeLLMDimensionValuesRequest := &DescribeLLMDimensionValuesRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Key:           util.PtrString(""),
		Filters:       []*Filter{},
	}
	result := &DescribeLLMDimensionValuesResponse{}
	result, err := APM_CLIENT.DescribeLLMDimensionValues(describeLLMDimensionValuesRequest)
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
func TestClient_DescribeLLMMetricData(t *testing.T) {
	describeLLMMetricDataRequest := &DescribeLLMMetricDataRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Metrics:       []*MetricQuery{},
		Filters:       []*Filter{},
		GroupBy:       []*string{},
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		Limit:         util.PtrInt32(int32(0)),
		PeriodSeconds: util.PtrInt32(int32(0)),
		Aggregate:     []*string{},
	}
	result := &DescribeLLMMetricDataResponse{}
	result, err := APM_CLIENT.DescribeLLMMetricData(describeLLMMetricDataRequest)
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
func TestClient_DescribeLLMServices(t *testing.T) {
	Tag := &Tag{
		Key:   util.PtrString(""),
		Value: util.PtrString(""),
	}
	describeLLMServicesRequest := &DescribeLLMServicesRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		ServiceName:   util.PtrString(""),
		ServiceId:     util.PtrString(""),
		Env:           util.PtrString(""),
		Tag:           Tag,
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
	}
	result := &DescribeLLMServicesResponse{}
	result, err := APM_CLIENT.DescribeLLMServices(describeLLMServicesRequest)
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
func TestClient_DescribeLLMSession(t *testing.T) {
	describeLLMSessionRequest := &DescribeLLMSessionRequest{
		SessionID:     util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
	}
	result := &DescribeLLMSessionResponse{}
	result, err := APM_CLIENT.DescribeLLMSession(describeLLMSessionRequest)
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
func TestClient_DescribeLLMSessions(t *testing.T) {
	describeLLMSessionsRequest := &DescribeLLMSessionsRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Filters:       []*Filter{},
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &DescribeLLMSessionsResponse{}
	result, err := APM_CLIENT.DescribeLLMSessions(describeLLMSessionsRequest)
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
func TestClient_DescribeLLMSessionsStatistics(t *testing.T) {
	describeLLMSessionsStatisticsRequest := &DescribeLLMSessionsStatisticsRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Filters:       []*Filter{},
	}
	result := &DescribeLLMSessionsStatisticsResponse{}
	result, err := APM_CLIENT.DescribeLLMSessionsStatistics(describeLLMSessionsStatisticsRequest)
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
func TestClient_DescribeLLMSpans(t *testing.T) {
	describeLLMSpansRequest := &DescribeLLMSpansRequest{
		BeginDatetime:       util.PtrString(""),
		EndDatetime:         util.PtrString(""),
		ParseLLMInputOutput: util.PtrBool(false),
		Filters:             []*Filter{},
		OrderBy:             util.PtrString(""),
		Order:               util.PtrString(""),
		Marker:              util.PtrString(""),
	}
	result := &DescribeLLMSpansResponse{}
	result, err := APM_CLIENT.DescribeLLMSpans(describeLLMSpansRequest)
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
func TestClient_DescribeLLMTrace(t *testing.T) {
	describeLLMTraceRequest := &DescribeLLMTraceRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		TraceId:       util.PtrString(""),
		Filters:       []*Filter{},
		ReturnHeight:  util.PtrBool(false),
	}
	result := &DescribeLLMTraceResponse{}
	result, err := APM_CLIENT.DescribeLLMTrace(describeLLMTraceRequest)
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
func TestClient_DescribeLLMTraces(t *testing.T) {
	describeLLMTracesRequest := &DescribeLLMTracesRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Filters:       []*Filter{},
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &DescribeLLMTracesResponse{}
	result, err := APM_CLIENT.DescribeLLMTraces(describeLLMTracesRequest)
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
func TestClient_DescribeLLMTracesStatistics(t *testing.T) {
	describeLLMTracesStatisticsRequest := &DescribeLLMTracesStatisticsRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Filters:       []*Filter{},
	}
	result := &DescribeLLMTracesStatisticsResponse{}
	result, err := APM_CLIENT.DescribeLLMTracesStatistics(describeLLMTracesStatisticsRequest)
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
		Action:                 util.PtrString(""),
		Metrics:                []*MetricQuery{},
		MetricsName:            util.PtrString(""),
		MetricsCompareTo:       []*string{},
		MetricsFilters:         []*Filter{},
		BeginDatetime:          util.PtrString(""),
		EndDatetime:            util.PtrString(""),
		Filters:                []*Filter{},
		GroupBy:                []*string{},
		OrderBy:                util.PtrString(""),
		Order:                  util.PtrString(""),
		Limit:                  util.PtrInt32(int32(0)),
		PeriodSeconds:          util.PtrInt32(int32(0)),
		ReserveEmptyDimensions: util.PtrBool(false),
	}
	result := &DescribeMetricDataResponse{}
	result, err := APM_CLIENT.DescribeMetricData(describeMetricDataRequest)
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
func TestClient_DescribeRetentionLimit(t *testing.T) {
	result := &DescribeRetentionLimitResponse{}
	result, err := APM_CLIENT.DescribeRetentionLimit()
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
func TestClient_DescribeServiceConfig(t *testing.T) {
	describeServiceConfigRequest := &DescribeServiceConfigRequest{
		ServiceName: util.PtrString(""),
	}
	result := &DescribeServiceConfigResponse{}
	result, err := APM_CLIENT.DescribeServiceConfig(describeServiceConfigRequest)
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
func TestClient_DescribeServicesMetrics(t *testing.T) {
	describeServicesMetricsRequest := &DescribeServicesMetricsRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Services:      []*string{},
		Metrics:       []*string{},
		MetricFilters: []*MetricFilter{},
	}
	result := &DescribeServicesMetricsResponse{}
	result, err := APM_CLIENT.DescribeServicesMetrics(describeServicesMetricsRequest)
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
func TestClient_DescribeServicesNames(t *testing.T) {
	Tag := &Tag{
		Key:   util.PtrString(""),
		Value: util.PtrString(""),
	}
	describeServicesNamesRequest := &DescribeServicesNamesRequest{
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		ServiceName:   util.PtrString(""),
		ServiceId:     util.PtrString(""),
		Language:      util.PtrString(""),
		Env:           util.PtrString(""),
		Source:        util.PtrString(""),
		Tag:           Tag,
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		MetricFilters: []*MetricFilter{},
	}
	result := &DescribeServicesNamesResponse{}
	result, err := APM_CLIENT.DescribeServicesNames(describeServicesNamesRequest)
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
func TestClient_DescribeSpanFieldValues(t *testing.T) {
	describeSpanFieldValuesRequest := &DescribeSpanFieldValuesRequest{
		Action:        util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Key:           util.PtrString(""),
		Filters:       []*Filter{},
	}
	result := &DescribeSpanFieldValuesResponse{}
	result, err := APM_CLIENT.DescribeSpanFieldValues(describeSpanFieldValuesRequest)
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
func TestClient_DescribeSpans(t *testing.T) {
	describeSpansRequest := &DescribeSpansRequest{
		Action:        util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Filters:       []*Filter{},
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		Marker:        util.PtrString(""),
	}
	result := &DescribeSpansResponse{}
	result, err := APM_CLIENT.DescribeSpans(describeSpansRequest)
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
func TestClient_DescribeTopology(t *testing.T) {
	describeTopologyRequest := &DescribeTopologyRequest{
		Action:        util.PtrString(""),
		ServiceName:   util.PtrString(""),
		Env:           util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
	}
	result := &DescribeTopologyResponse{}
	result, err := APM_CLIENT.DescribeTopology(describeTopologyRequest)
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
func TestClient_DescribeTrace(t *testing.T) {
	describeTraceRequest := &DescribeTraceRequest{
		Action:        util.PtrString(""),
		SpanDatetime:  util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		TraceId:       util.PtrString(""),
		Filters:       []*Filter{},
		ReturnHeight:  util.PtrBool(false),
	}
	result := &DescribeTraceResponse{}
	result, err := APM_CLIENT.DescribeTrace(describeTraceRequest)
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
func TestClient_DescribeTraceMetricData(t *testing.T) {
	describeTraceMetricDataRequest := &DescribeTraceMetricDataRequest{
		Action:        util.PtrString(""),
		BeginDatetime: util.PtrString(""),
		EndDatetime:   util.PtrString(""),
		Metrics:       []*TraceMetricQuery{},
		MetricsName:   util.PtrString(""),
		Filters:       []*Filter{},
		GroupBy:       []*string{},
		PeriodSeconds: util.PtrInt32(int32(0)),
		Aggregate:     []*string{},
	}
	result := &DescribeTraceMetricDataResponse{}
	result, err := APM_CLIENT.DescribeTraceMetricData(describeTraceMetricDataRequest)
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
func TestClient_UpdateServiceConfig(t *testing.T) {
	SampleConfig := &SampleConfig{
		Enabled:    util.PtrBool(false),
		Processors: []*SampleProcessor{},
	}
	LoggingConfig := &LoggingConfig{
		Enabled:      util.PtrBool(false),
		Region:       util.PtrString(""),
		Project:      util.PtrString(""),
		LogStoreName: util.PtrString(""),
		TraceIdIndex: util.PtrString(""),
		TraceIdKey:   util.PtrString(""),
		SpanIdIndex:  util.PtrString(""),
		SpanIdKey:    util.PtrString(""),
	}
	RequestConfig := &ServiceRequestConfig{
		Global:                            util.PtrBool(false),
		ServerSlowRequestThresholdSeconds: util.PtrFloat64(float64(0)),
		DbSlowRequestThresholdSeconds:     util.PtrFloat64(float64(0)),
		OkHttpStatus:                      []*int32{},
	}
	TopoConfig := &ServiceTopoConfig{
		Global:                  util.PtrBool(false),
		RequestSecondsThreshold: util.PtrFloat64(float64(0)),
		ErrorRateThreshold:      util.PtrFloat64(float64(0)),
	}
	MllmResourceDumpConfig := &MllmResourceDumpConfig{
		RetentionDays: util.PtrInt32(int32(0)),
		Bucket:        util.PtrString(""),
	}
	updateServiceConfigRequest := &UpdateServiceConfigRequest{
		ServiceNames:           []*string{},
		SampleConfig:           SampleConfig,
		LoggingConfig:          LoggingConfig,
		RequestConfig:          RequestConfig,
		TopoConfig:             TopoConfig,
		MllmResourceDumpConfig: MllmResourceDumpConfig,
	}
	result := &UpdateServiceConfigResponse{}
	result, err := APM_CLIENT.UpdateServiceConfig(updateServiceConfigRequest)
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
