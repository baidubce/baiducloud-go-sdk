package cprom

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
	CPROM_CLIENT *Client
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

	CPROM_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_BindCluster(t *testing.T) {
	bindClusterRequest := &BindClusterRequest{
		InstanceId: util.PtrString(""),
		Action:     util.PtrString(""),
		ClusterId:  util.PtrString(""),
	}
	result := &BindClusterResponse{}
	result, err := CPROM_CLIENT.BindCluster(bindClusterRequest)
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
func TestClient_ClaimAlertEvent(t *testing.T) {
	claimAlertEventRequest := &ClaimAlertEventRequest{
		EventIds:    []*string{},
		ClaimReason: util.PtrString(""),
	}
	result := &ClaimAlertEventResponse{}
	result, err := CPROM_CLIENT.ClaimAlertEvent(claimAlertEventRequest)
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
func TestClient_CreateAlert(t *testing.T) {
	Labels := make(map[string]string)
	Annotations := make(map[string]string)
	createAlertRequest := &CreateAlertRequest{
		InstanceId:   util.PtrString(""),
		AlertName:    util.PtrString(""),
		Expr:         util.PtrString(""),
		CpromFor:     util.PtrString(""),
		Description:  util.PtrString(""),
		NotifyRuleId: util.PtrString(""),
		Enable:       util.PtrBool(false),
		Severity:     util.PtrString(""),
		Labels:       nil,
		Annotations:  nil,
	}
	result := &CreateAlertResponse{}
	result, err := CPROM_CLIENT.CreateAlert(createAlertRequest)
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
func TestClient_CreateCustomScrapeTask(t *testing.T) {
	createCustomScrapeTaskRequest := &CreateCustomScrapeTaskRequest{
		InstanceId: util.PtrString(""),
		AgentId:    util.PtrString(""),
		Config:     util.PtrString(""),
	}
	result := &CreateCustomScrapeTaskResponse{}
	result, err := CPROM_CLIENT.CreateCustomScrapeTask(createCustomScrapeTaskRequest)
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
func TestClient_CreateInstance(t *testing.T) {
	createInstanceRequest := &CreateInstanceRequest{
		InstanceName:         util.PtrString(""),
		InstanceType:         util.PtrString(""),
		InstanceSpec:         util.PtrString(""),
		RetentionPeriod:      util.PtrString(""),
		NeedGrafana:          util.PtrBool(false),
		GrafanaName:          util.PtrString(""),
		GrafanaAdminPassword: util.PtrString(""),
	}
	err := CPROM_CLIENT.CreateInstance(createInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateNotificationPolicy(t *testing.T) {
	RepeatNotifyConfig := &RepeatNotifyConfig{
		Enabled:      util.PtrBool(false),
		IntervalHour: util.PtrInt32(int32(0)),
		IntervalMin:  util.PtrInt32(int32(0)),
		MaxCount:     util.PtrInt32(int32(0)),
		Strategy:     util.PtrString(""),
	}
	createNotificationPolicyRequest := &CreateNotificationPolicyRequest{
		NotifyRuleName:     util.PtrString(""),
		StartTime:          util.PtrString(""),
		EndTime:            util.PtrString(""),
		Channel:            []*string{},
		ReceiverType:       util.PtrString(""),
		Users:              []*User{},
		UserGroups:         []*UserGroup{},
		WebhookConfigList:  []*CallbackConfig{},
		EscalateConfigList: []*EscalateParam{},
		RepeatNotifyConfig: RepeatNotifyConfig,
	}
	result := &CreateNotificationPolicyResponse{}
	result, err := CPROM_CLIENT.CreateNotificationPolicy(createNotificationPolicyRequest)
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
func TestClient_CreatePodmonitor(t *testing.T) {
	Metadata := &ObjectMeta{
		Name:      util.PtrString(""),
		Namespace: util.PtrString(""),
	}
	Spec := &PodMonitorSpec{
		NamespaceSelector: &NamespaceSelector{
			MatchNames: []*string{},
		},
		PodMetricsEndpoints: []*PodMetricsEndpoint{},
		Selector: &LabelSelector{
			MatchLabels: nil,
		},
	}
	createPodmonitorRequest := &CreatePodmonitorRequest{
		InstanceId: util.PtrString(""),
		AgentId:    util.PtrString(""),
		ApiVersion: util.PtrString(""),
		Kind:       util.PtrString(""),
		Metadata:   Metadata,
		Spec:       Spec,
	}
	result := &CreatePodmonitorResponse{}
	result, err := CPROM_CLIENT.CreatePodmonitor(createPodmonitorRequest)
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
func TestClient_CreateServiceMonitor(t *testing.T) {
	Metadata := &ObjectMeta{
		Name:      util.PtrString(""),
		Namespace: util.PtrString(""),
	}
	Spec := &ServiceMonitorSpec{
		Endpoints: []*ServiceMonitorEndpoint{},
		NamespaceSelector: &NamespaceSelector{
			MatchNames: []*string{},
		},
		Selector: &LabelSelector{
			MatchLabels: nil,
		},
	}
	createServiceMonitorRequest := &CreateServiceMonitorRequest{
		InstanceId: util.PtrString(""),
		AgentId:    util.PtrString(""),
		ApiVersion: util.PtrString(""),
		Kind:       util.PtrString(""),
		Metadata:   Metadata,
		Spec:       Spec,
	}
	result := &CreateServiceMonitorResponse{}
	result, err := CPROM_CLIENT.CreateServiceMonitor(createServiceMonitorRequest)
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
func TestClient_DeleteAlert(t *testing.T) {
	deleteAlertRequest := &DeleteAlertRequest{
		AlertingRuleId: util.PtrString(""),
		InstanceId:     util.PtrString(""),
	}
	err := CPROM_CLIENT.DeleteAlert(deleteAlertRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCustomScrapeTask(t *testing.T) {
	deleteCustomScrapeTaskRequest := &DeleteCustomScrapeTaskRequest{
		ScrapeJobId: util.PtrString(""),
		InstanceId:  util.PtrString(""),
		AgentId:     util.PtrString(""),
	}
	err := CPROM_CLIENT.DeleteCustomScrapeTask(deleteCustomScrapeTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstance(t *testing.T) {
	deleteInstanceRequest := &DeleteInstanceRequest{
		InstanceId:    util.PtrString(""),
		DeleteGrafana: util.PtrBool(false),
	}
	err := CPROM_CLIENT.DeleteInstance(deleteInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteNotificationPolicy(t *testing.T) {
	deleteNotificationPolicyRequest := &DeleteNotificationPolicyRequest{
		NotifyRuleId: util.PtrString(""),
	}
	err := CPROM_CLIENT.DeleteNotificationPolicy(deleteNotificationPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeletePodmonitor(t *testing.T) {
	deletePodmonitorRequest := &DeletePodmonitorRequest{
		PodMonitorName: util.PtrString(""),
		InstanceId:     util.PtrString(""),
		AgentId:        util.PtrString(""),
	}
	err := CPROM_CLIENT.DeletePodmonitor(deletePodmonitorRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteServiceMonitor(t *testing.T) {
	deleteServiceMonitorRequest := &DeleteServiceMonitorRequest{
		ServiceMonitorName: util.PtrString(""),
		InstanceId:         util.PtrString(""),
		AgentId:            util.PtrString(""),
	}
	err := CPROM_CLIENT.DeleteServiceMonitor(deleteServiceMonitorRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GenerateInstanceToken(t *testing.T) {
	generateInstanceTokenRequest := &GenerateInstanceTokenRequest{
		InstanceId: util.PtrString(""),
		Action:     util.PtrString(""),
		Token:      util.PtrString(""),
	}
	err := CPROM_CLIENT.GenerateInstanceToken(generateInstanceTokenRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetAlertDetail(t *testing.T) {
	getAlertDetailRequest := &GetAlertDetailRequest{
		AlertingRuleId: util.PtrString(""),
		InstanceId:     util.PtrString(""),
	}
	result := &GetAlertDetailResponse{}
	result, err := CPROM_CLIENT.GetAlertDetail(getAlertDetailRequest)
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
func TestClient_GetAlertEventDetail(t *testing.T) {
	getAlertEventDetailRequest := &GetAlertEventDetailRequest{
		EventId: util.PtrString(""),
	}
	result := &GetAlertEventDetailResponse{}
	result, err := CPROM_CLIENT.GetAlertEventDetail(getAlertEventDetailRequest)
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
func TestClient_GetClusterBindStatus(t *testing.T) {
	getClusterBindStatusRequest := &GetClusterBindStatusRequest{
		ClusterId: util.PtrString(""),
	}
	result := &GetClusterBindStatusResponse{}
	result, err := CPROM_CLIENT.GetClusterBindStatus(getClusterBindStatusRequest)
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
func TestClient_GetNotificationPolicy(t *testing.T) {
	getNotificationPolicyRequest := &GetNotificationPolicyRequest{
		NotifyRuleId: util.PtrString(""),
	}
	result := &GetNotificationPolicyResponse{}
	result, err := CPROM_CLIENT.GetNotificationPolicy(getNotificationPolicyRequest)
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
func TestClient_GetPodMonitorDetail(t *testing.T) {
	getPodMonitorDetailRequest := &GetPodMonitorDetailRequest{
		PodMonitorName: util.PtrString(""),
		InstanceId:     util.PtrString(""),
		AgentId:        util.PtrString(""),
	}
	result := &GetPodMonitorDetailResponse{}
	result, err := CPROM_CLIENT.GetPodMonitorDetail(getPodMonitorDetailRequest)
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
func TestClient_GetServiceMonitorDetail(t *testing.T) {
	getServiceMonitorDetailRequest := &GetServiceMonitorDetailRequest{
		ServiceMonitorName: util.PtrString(""),
		InstanceId:         util.PtrString(""),
		AgentId:            util.PtrString(""),
	}
	result := &GetServiceMonitorDetailResponse{}
	result, err := CPROM_CLIENT.GetServiceMonitorDetail(getServiceMonitorDetailRequest)
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
func TestClient_ListAlertEvents(t *testing.T) {
	listAlertEventsRequest := &ListAlertEventsRequest{
		StartTime:         util.PtrInt32(int32(0)),
		EndTime:           util.PtrInt32(int32(0)),
		PageNo:            util.PtrInt32(int32(0)),
		PageSize:          util.PtrInt32(int32(0)),
		MonitorInstanceId: util.PtrString(""),
		AlertingRuleId:    util.PtrString(""),
		AlertingRuleName:  util.PtrString(""),
		NotifyRuleId:      util.PtrString(""),
		NotifyRuleName:    util.PtrString(""),
		Severity:          util.PtrString(""),
		Status:            util.PtrString(""),
		Expr:              util.PtrString(""),
		OrderBy:           util.PtrString(""),
		Order:             util.PtrString(""),
		AlarmTags:         util.PtrString(""),
	}
	result := &ListAlertEventsResponse{}
	result, err := CPROM_CLIENT.ListAlertEvents(listAlertEventsRequest)
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
func TestClient_ListAlertTemplates(t *testing.T) {
	result := &ListAlertTemplatesResponse{}
	result, err := CPROM_CLIENT.ListAlertTemplates()
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
func TestClient_ListAlerts(t *testing.T) {
	listAlertsRequest := &ListAlertsRequest{
		InstanceId:  util.PtrString(""),
		PageSize:    util.PtrInt32(int32(0)),
		PageNo:      util.PtrInt32(int32(0)),
		KeywordType: util.PtrString(""),
		Keyword:     util.PtrString(""),
	}
	result := &ListAlertsResponse{}
	result, err := CPROM_CLIENT.ListAlerts(listAlertsRequest)
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
func TestClient_ListBindableCloudProducts(t *testing.T) {
	result := &ListBindableCloudProductsResponse{}
	result, err := CPROM_CLIENT.ListBindableCloudProducts()
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
func TestClient_ListInstances(t *testing.T) {
	listInstancesRequest := &ListInstancesRequest{
		PageSize: util.PtrInt32(int32(0)),
		PageNo:   util.PtrInt32(int32(0)),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
		Phase:    util.PtrString(""),
	}
	result := &ListInstancesResponse{}
	result, err := CPROM_CLIENT.ListInstances(listInstancesRequest)
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
func TestClient_ListNotificationPolicies(t *testing.T) {
	listNotificationPoliciesRequest := &ListNotificationPoliciesRequest{
		PageSize:    util.PtrInt32(int32(0)),
		PageNo:      util.PtrInt32(int32(0)),
		KeywordType: util.PtrString(""),
		Keyword:     util.PtrString(""),
	}
	result := &ListNotificationPoliciesResponse{}
	result, err := CPROM_CLIENT.ListNotificationPolicies(listNotificationPoliciesRequest)
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
func TestClient_ListPodMonitors(t *testing.T) {
	listPodMonitorsRequest := &ListPodMonitorsRequest{
		InstanceId:  util.PtrString(""),
		AgentId:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
		KeywordType: util.PtrString(""),
		Keyword:     util.PtrString(""),
		OrderBy:     util.PtrString(""),
		Order:       util.PtrString(""),
	}
	result := &ListPodMonitorsResponse{}
	result, err := CPROM_CLIENT.ListPodMonitors(listPodMonitorsRequest)
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
func TestClient_ListRelatedCloudProducts(t *testing.T) {
	listRelatedCloudProductsRequest := &ListRelatedCloudProductsRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ListRelatedCloudProductsResponse{}
	result, err := CPROM_CLIENT.ListRelatedCloudProducts(listRelatedCloudProductsRequest)
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
func TestClient_ListServiceMonitors(t *testing.T) {
	listServiceMonitorsRequest := &ListServiceMonitorsRequest{
		InstanceId:  util.PtrString(""),
		AgentId:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
		KeywordType: util.PtrString(""),
		Keyword:     util.PtrString(""),
		OrderBy:     util.PtrString(""),
		Order:       util.PtrString(""),
	}
	result := &ListServiceMonitorsResponse{}
	result, err := CPROM_CLIENT.ListServiceMonitors(listServiceMonitorsRequest)
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
func TestClient_RemoteRead(t *testing.T) {
	remoteReadRequest := &RemoteReadRequest{
		RemoteReadUrl: util.PtrString(""),
		Query:         util.PtrString(""),
		Step:          util.PtrInt32(int32(0)),
		Start:         util.PtrInt64(int64(0)),
		End:           util.PtrInt64(int64(0)),
	}
	result := &RemoteReadResponse{}
	result, err := CPROM_CLIENT.RemoteRead(remoteReadRequest)
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
func TestClient_RemoteWrite(t *testing.T) {
	remoteWriteRequest := &RemoteWriteRequest{
		RemoteWriteUrl:  util.PtrString(""),
		ContentType:     util.PtrString(""),
		ContentEncoding: util.PtrString(""),
		InstanceId:      util.PtrString(""),
		Authorization:   util.PtrString(""),
		Body:            nil,
	}
	err := CPROM_CLIENT.RemoteWrite(remoteWriteRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TogglePodMonitorService(t *testing.T) {
	togglePodMonitorServiceRequest := &TogglePodMonitorServiceRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		AgentId:    util.PtrString(""),
	}
	err := CPROM_CLIENT.TogglePodMonitorService(togglePodMonitorServiceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ToggleServiceMonitorService(t *testing.T) {
	toggleServiceMonitorServiceRequest := &ToggleServiceMonitorServiceRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		AgentId:    util.PtrString(""),
	}
	err := CPROM_CLIENT.ToggleServiceMonitorService(toggleServiceMonitorServiceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAlert(t *testing.T) {
	Labels := make(map[string]string)
	Annotations := make(map[string]string)
	updateAlertRequest := &UpdateAlertRequest{
		AlertingRuleId: util.PtrString(""),
		InstanceId:     util.PtrString(""),
		AlertName:      util.PtrString(""),
		Expr:           util.PtrString(""),
		CpromFor:       util.PtrString(""),
		Description:    util.PtrString(""),
		NotifyRuleId:   util.PtrString(""),
		Severity:       util.PtrString(""),
		Enable:         util.PtrBool(false),
		Labels:         nil,
		Annotations:    nil,
	}
	err := CPROM_CLIENT.UpdateAlert(updateAlertRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateNotificationPolicy(t *testing.T) {
	RepeatNotifyConfig := &RepeatNotifyConfig{
		Enabled:      util.PtrBool(false),
		IntervalHour: util.PtrInt32(int32(0)),
		IntervalMin:  util.PtrInt32(int32(0)),
		MaxCount:     util.PtrInt32(int32(0)),
		Strategy:     util.PtrString(""),
	}
	updateNotificationPolicyRequest := &UpdateNotificationPolicyRequest{
		NotifyRuleId:       util.PtrString(""),
		NotifyRuleName:     util.PtrString(""),
		StartTime:          util.PtrString(""),
		EndTime:            util.PtrString(""),
		Channel:            []*string{},
		ReceiverType:       util.PtrString(""),
		Users:              []*User{},
		UserGroups:         []*UserGroup{},
		WebhookConfigList:  []*CallbackConfig{},
		EscalateConfigList: []*EscalateParam{},
		RepeatNotifyConfig: RepeatNotifyConfig,
	}
	err := CPROM_CLIENT.UpdateNotificationPolicy(updateNotificationPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePodMonitor(t *testing.T) {
	PodMonitor := &PodMonitor{
		ApiVersion: util.PtrString(""),
		Kind:       util.PtrString(""),
		Metadata: &ObjectMeta{
			Name:      util.PtrString(""),
			Namespace: util.PtrString(""),
		},
		Spec: &PodMonitorSpec{
			NamespaceSelector: &NamespaceSelector{
				MatchNames: []*string{},
			},
			PodMetricsEndpoints: []*PodMetricsEndpoint{},
			Selector: &LabelSelector{
				MatchLabels: nil,
			},
		},
	}
	updatePodMonitorRequest := &UpdatePodMonitorRequest{
		PodMonitorName: util.PtrString(""),
		InstanceId:     util.PtrString(""),
		AgentId:        util.PtrString(""),
		Enable:         util.PtrString(""),
		PodMonitor:     PodMonitor,
	}
	err := CPROM_CLIENT.UpdatePodMonitor(updatePodMonitorRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRelatedCloudProducts(t *testing.T) {
	updateRelatedCloudProductsRequest := &UpdateRelatedCloudProductsRequest{
		InstanceId: util.PtrString(""),
		Scopes:     []*string{},
	}
	err := CPROM_CLIENT.UpdateRelatedCloudProducts(updateRelatedCloudProductsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateServiceMonitor(t *testing.T) {
	ServiceMonitor := &ServiceMonitor{
		ApiVersion: util.PtrString(""),
		Kind:       util.PtrString(""),
		Metadata: &ObjectMeta{
			Name:      util.PtrString(""),
			Namespace: util.PtrString(""),
		},
		Spec: &ServiceMonitorSpec{
			Endpoints: []*ServiceMonitorEndpoint{},
			NamespaceSelector: &NamespaceSelector{
				MatchNames: []*string{},
			},
			Selector: &LabelSelector{
				MatchLabels: nil,
			},
		},
	}
	updateServiceMonitorRequest := &UpdateServiceMonitorRequest{
		ServiceMonitorName: util.PtrString(""),
		InstanceId:         util.PtrString(""),
		AgentId:            util.PtrString(""),
		Enable:             util.PtrString(""),
		ServiceMonitor:     ServiceMonitor,
	}
	err := CPROM_CLIENT.UpdateServiceMonitor(updateServiceMonitorRequest)
	ExpectEqual(t.Errorf, nil, err)
}
