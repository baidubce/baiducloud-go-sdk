package as

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
	AS_CLIENT *Client
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

	AS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AdjustNumV2(t *testing.T) {
	adjustNumV2Request := &AdjustNumV2Request{
		GroupId:    util.PtrString(""),
		AdjustNode: util.PtrString(""),
		AdjustNum:  util.PtrInt32(int32(0)),
	}
	result := &AdjustNumV2Response{}
	result, err := AS_CLIENT.AdjustNumV2(adjustNumV2Request)
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
func TestClient_AttachNodeV2(t *testing.T) {
	attachNodeV2Request := &AttachNodeV2Request{
		GroupId:    util.PtrString(""),
		AttachNode: util.PtrString(""),
		Nodes:      []*string{},
	}
	result := &AttachNodeV2Response{}
	result, err := AS_CLIENT.AttachNodeV2(attachNodeV2Request)
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
func TestClient_CreateAsGroupV2(t *testing.T) {
	Config := &GroupConfig{
		MinNodeNum:    util.PtrInt32(int32(0)),
		MaxNodeNum:    util.PtrInt32(int32(0)),
		CooldownInSec: util.PtrInt32(int32(0)),
		ExpectNum:     util.PtrInt32(int32(0)),
		InitNum:       util.PtrInt32(int32(0)),
	}
	Eip := &EipInfo{
		IfBindEip:       util.PtrBool(false),
		BandwidthInMbps: util.PtrInt32(int32(0)),
		EipProductType:  util.PtrString(""),
		PurchaseType:    util.PtrString(""),
	}
	EipConfig := &EipConfig{
		EipGroupBindStrategy:   util.PtrString(""),
		EipGroupUnbindStrategy: util.PtrString(""),
		EipGroupIdList:         []*string{},
		Increase: &EipGroupIncrease{
			Enabled:  util.PtrBool(false),
			Strategy: util.PtrString(""),
		},
		Decrease: &EipGroupDecrease{
			Enabled: util.PtrBool(false),
		},
		Bandwidth: &EipGroupBandwidth{
			Max:      util.PtrInt32(int32(0)),
			Min:      util.PtrInt32(int32(0)),
			Standard: util.PtrInt32(int32(0)),
		},
	}
	Billing := &BillingInfo{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLengthInMonth: util.PtrInt32(int32(0)),
		},
	}
	HealthCheck := &HealthCheckConfig{
		HealthCheckInterval: util.PtrInt32(int32(0)),
		GraceTime:           util.PtrInt32(int32(0)),
	}
	AssignTagInfo := &AssignTagInfo{
		ResourceId:  util.PtrString(""),
		RelationTag: util.PtrBool(false),
		Tags:        []*TagInfo{},
	}
	CmdConfig := &CmdConfig{
		HasDecreaseCmd: util.PtrBool(false),
		DecCmdStrategy: util.PtrString(""),
		DecCmdData:     util.PtrString(""),
		DecCmdTimeout:  util.PtrInt32(int32(0)),
		DecCmdManual:   util.PtrBool(false),
		HasIncreaseCmd: util.PtrBool(false),
		IncCmdStrategy: util.PtrString(""),
		IncCmdData:     util.PtrString(""),
		IncCmdTimeout:  util.PtrInt32(int32(0)),
		IncCmdManual:   util.PtrBool(false),
	}
	BccNameConfig := &BccNameConfig{
		BccName:            util.PtrString(""),
		BccHostname:        util.PtrString(""),
		AutoSeqSuffix:      util.PtrBool(false),
		OpenHostnameDomain: util.PtrBool(false),
	}
	createAsGroupV2Request := &CreateAsGroupV2Request{
		GroupName:         util.PtrString(""),
		ZoneInfo:          []*ZoneInfo{},
		Config:            Config,
		KeypairId:         util.PtrString(""),
		KeypairName:       util.PtrString(""),
		KeepImageLogin:    util.PtrInt32(int32(0)),
		Blb:               []*BlbInfo{},
		BlbUnbindWaitTime: util.PtrInt32(int32(0)),
		Nodes:             []*NodeInfo{},
		Eip:               Eip,
		EipConfig:         EipConfig,
		Billing:           Billing,
		Rds:               []*string{},
		Scs:               []*string{},
		HealthCheck:       HealthCheck,
		ExpansionStrategy: util.PtrString(""),
		ShrinkageStrategy: util.PtrString(""),
		AssignTagInfo:     AssignTagInfo,
		CmdConfig:         CmdConfig,
		BccNameConfig:     BccNameConfig,
	}
	result := &CreateAsGroupV2Response{}
	result, err := AS_CLIENT.CreateAsGroupV2(createAsGroupV2Request)
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
func TestClient_CreateRuleV2(t *testing.T) {
	AsAlarmRule := &AsAlarmRule{
		Id:    util.PtrInt32(int32(0)),
		Scope: util.PtrString(""),
		MonitorObject: &MonitorObject{
			AsType:    util.PtrString(""),
			Names:     []*string{},
			Resources: []*PolicyResource{},
			TypeName:  util.PtrString(""),
		},
		Rules:               [][]AlarmRule{},
		AlarmName:           util.PtrString(""),
		AliasName:           util.PtrString(""),
		InsufficientCycle:   util.PtrInt32(int32(0)),
		PolicyEnabled:       util.PtrBool(false),
		RuleContents:        []*string{},
		RuleContentsEn:      []*string{},
		Source:              util.PtrString(""),
		ComponentType:       util.PtrString(""),
		AlarmActions:        []*string{},
		OkActions:           []*string{},
		InsufficientActions: []*string{},
	}
	createRuleV2Request := &CreateRuleV2Request{
		RuleName:        util.PtrString(""),
		GroupId:         util.PtrString(""),
		State:           util.PtrString(""),
		AsType:          util.PtrString(""),
		ActionType:      util.PtrString(""),
		ActionNum:       util.PtrInt32(int32(0)),
		CronTime:        util.PtrString(""),
		CooldownInSec:   util.PtrInt32(int32(0)),
		PeriodType:      util.PtrString(""),
		PeriodValue:     util.PtrInt32(int32(0)),
		PeriodStartTime: util.PtrString(""),
		PeriodEndTime:   util.PtrString(""),
		AsAlarmRule:     AsAlarmRule,
	}
	result := &CreateRuleV2Response{}
	result, err := AS_CLIENT.CreateRuleV2(createRuleV2Request)
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
func TestClient_DeleteAsGroupV2(t *testing.T) {
	deleteAsGroupV2Request := &DeleteAsGroupV2Request{
		GroupIds: []*string{},
	}
	err := AS_CLIENT.DeleteAsGroupV2(deleteAsGroupV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRuleV2(t *testing.T) {
	deleteRuleV2Request := &DeleteRuleV2Request{
		RuleId: util.PtrString(""),
	}
	err := AS_CLIENT.DeleteRuleV2(deleteRuleV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DetachNodeV2(t *testing.T) {
	detachNodeV2Request := &DetachNodeV2Request{
		GroupId:    util.PtrString(""),
		DetachNode: util.PtrString(""),
		Nodes:      []*string{},
	}
	result := &DetachNodeV2Response{}
	result, err := AS_CLIENT.DetachNodeV2(detachNodeV2Request)
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
func TestClient_ExecRuleV2(t *testing.T) {
	execRuleV2Request := &ExecRuleV2Request{
		GroupId:  util.PtrString(""),
		ExecRule: util.PtrString(""),
		RuleId:   util.PtrString(""),
	}
	result := &ExecRuleV2Response{}
	result, err := AS_CLIENT.ExecRuleV2(execRuleV2Request)
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
func TestClient_GetAsGroupV2(t *testing.T) {
	getAsGroupV2Request := &GetAsGroupV2Request{
		GroupId: util.PtrString(""),
	}
	result := &GetAsGroupV2Response{}
	result, err := AS_CLIENT.GetAsGroupV2(getAsGroupV2Request)
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
func TestClient_GetRuleV2(t *testing.T) {
	getRuleV2Request := &GetRuleV2Request{
		RuleId: util.PtrString(""),
	}
	result := &GetRuleV2Response{}
	result, err := AS_CLIENT.GetRuleV2(getRuleV2Request)
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
func TestClient_ListAsGroupV2(t *testing.T) {
	listAsGroupV2Request := &ListAsGroupV2Request{
		PageNo:         util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
		Keyword:        util.PtrString(""),
		KeywordType:    util.PtrString(""),
		SubKeywordType: util.PtrString(""),
		Order:          util.PtrString(""),
		OrderBy:        util.PtrString(""),
	}
	result := &ListAsGroupV2Response{}
	result, err := AS_CLIENT.ListAsGroupV2(listAsGroupV2Request)
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
func TestClient_ListAsNodeV2(t *testing.T) {
	listAsNodeV2Request := &ListAsNodeV2Request{
		Groupid:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
		Keyword:     util.PtrString(""),
		KeywordType: util.PtrString(""),
		Order:       util.PtrString(""),
		OrderBy:     util.PtrString(""),
	}
	result := &ListAsNodeV2Response{}
	result, err := AS_CLIENT.ListAsNodeV2(listAsNodeV2Request)
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
func TestClient_ListRuleV2(t *testing.T) {
	listRuleV2Request := &ListRuleV2Request{
		Groupid:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
		Keyword:     util.PtrString(""),
		KeywordType: util.PtrString(""),
		Order:       util.PtrString(""),
		OrderBy:     util.PtrString(""),
	}
	result := &ListRuleV2Response{}
	result, err := AS_CLIENT.ListRuleV2(listRuleV2Request)
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
func TestClient_ListTaskV2(t *testing.T) {
	listTaskV2Request := &ListTaskV2Request{
		Groupid:   util.PtrString(""),
		OrderBy:   util.PtrString(""),
		PageNo:    util.PtrInt32(int32(0)),
		Order:     util.PtrString(""),
		PageSize:  util.PtrInt32(int32(0)),
		StartTime: util.PtrString(""),
		EndTime:   util.PtrString(""),
	}
	result := &ListTaskV2Response{}
	result, err := AS_CLIENT.ListTaskV2(listTaskV2Request)
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
func TestClient_ScalingDownV2(t *testing.T) {
	scalingDownV2Request := &ScalingDownV2Request{
		GroupId:     util.PtrString(""),
		ScalingDown: util.PtrString(""),
		Nodes:       []*string{},
	}
	result := &ScalingDownV2Response{}
	result, err := AS_CLIENT.ScalingDownV2(scalingDownV2Request)
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
func TestClient_ScalingUpV2(t *testing.T) {
	scalingUpV2Request := &ScalingUpV2Request{
		GroupId:           util.PtrString(""),
		ScalingUp:         util.PtrString(""),
		NodeCount:         util.PtrInt32(int32(0)),
		Zone:              []*string{},
		ExpansionStrategy: util.PtrString(""),
	}
	result := &ScalingUpV2Response{}
	result, err := AS_CLIENT.ScalingUpV2(scalingUpV2Request)
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
func TestClient_UpdateIsManagedV2(t *testing.T) {
	updateIsManagedV2Request := &UpdateIsManagedV2Request{
		GroupId:           util.PtrString(""),
		UpdateIsManaged:   util.PtrString(""),
		AddManagedNodeIds: []*string{},
		DelManagedNodeIds: []*string{},
	}
	err := AS_CLIENT.UpdateIsManagedV2(updateIsManagedV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateProtectV2(t *testing.T) {
	updateProtectV2Request := &UpdateProtectV2Request{
		GroupId:       util.PtrString(""),
		UpdateProtect: util.PtrString(""),
		Nodes:         []*string{},
		IsProtected:   util.PtrBool(false),
	}
	err := AS_CLIENT.UpdateProtectV2(updateProtectV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
