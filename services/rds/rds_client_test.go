package rds

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
	RDS_CLIENT *Client
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

	// ==== AK/SK 鉴权 ====
	RDS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_CreateAPrimaryInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation:   nil,
	}
	InitialDataReference := &InitialDataReference{
		InstanceId:    util.PtrString(""),
		ReferenceType: util.PtrString(""),
		Datetime:      util.PtrString(""),
		SnapshotId:    util.PtrString(""),
	}
	createAPrimaryInstanceRequest := &CreateAPrimaryInstanceRequest{
		ClientToken:          util.PtrString(""),
		Billing:              Billing,
		PurchaseCount:        util.PtrInt32(int32(0)),
		InstanceName:         util.PtrString(""),
		Engine:               util.PtrString(""),
		EngineVersion:        util.PtrString(""),
		Category:             util.PtrString(""),
		CpuCount:             util.PtrInt32(int32(0)),
		MemoryCapacity:       util.PtrInt32(int32(0)),
		VolumeCapacity:       util.PtrInt32(int32(0)),
		DiskIoType:           util.PtrString(""),
		ZoneNames:            []*string{},
		BcmGroupName:         util.PtrString(""),
		VpcId:                util.PtrString(""),
		IsDirectPay:          util.PtrBool(false),
		Subnets:              []*SubnetMap{},
		Tags:                 []*Tag{},
		BgwGroupId:           util.PtrString(""),
		BgwGroupExclusive:    util.PtrBool(false),
		CharacterSetName:     util.PtrString(""),
		AutoRenewTimeUnit:    util.PtrString(""),
		AutoRenewTime:        util.PtrInt32(int32(0)),
		LowerCaseTableNames:  util.PtrInt32(int32(0)),
		ParameterTemplateId:  util.PtrString(""),
		Ovip:                 util.PtrString(""),
		EntryPort:            util.PtrInt32(int32(0)),
		ReplicationType:      util.PtrString(""),
		ResourceGroupId:      util.PtrString(""),
		InitialDataReference: InitialDataReference,
		Data:                 []*RecoveryToSourceInstanceModel{},
		LeaderInstanceId:     util.PtrString(""),
		ResourceType:         util.PtrString(""),
	}
	result := &CreateAPrimaryInstanceResponse{}
	result, err := RDS_CLIENT.CreateAPrimaryInstance(createAPrimaryInstanceRequest)
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
func TestClient_CreateAProxyInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation:   nil,
	}
	createAProxyInstanceRequest := &CreateAProxyInstanceRequest{
		ClientToken:      util.PtrString(""),
		Billing:          Billing,
		SourceInstanceId: util.PtrString(""),
		InstanceName:     util.PtrString(""),
		NodeAmount:       util.PtrInt32(int32(0)),
		ZoneNames:        []*string{},
		VpcId:            util.PtrString(""),
		IsDirectPay:      util.PtrBool(false),
		Subnets:          []*SubnetMap{},
		Ovip:             util.PtrString(""),
		EntryPort:        util.PtrInt32(int32(0)),
		ResourceGroupId:  util.PtrString(""),
	}
	result := &CreateAProxyInstanceResponse{}
	result, err := RDS_CLIENT.CreateAProxyInstance(createAProxyInstanceRequest)
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
func TestClient_CreateReadOnlyInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation:   nil,
	}
	createReadOnlyInstanceRequest := &CreateReadOnlyInstanceRequest{
		ClientToken:           util.PtrString(""),
		Billing:               Billing,
		SourceInstanceId:      util.PtrString(""),
		InstanceName:          util.PtrString(""),
		CpuCount:              util.PtrInt32(int32(0)),
		MemoryCapacity:        util.PtrInt32(int32(0)),
		VolumeCapacity:        util.PtrInt32(int32(0)),
		DiskIoType:            util.PtrString(""),
		ZoneNames:             []*string{},
		VpcId:                 util.PtrString(""),
		BcmGroupName:          util.PtrString(""),
		IsDirectPay:           util.PtrBool(false),
		Subnets:               []*SubnetMap{},
		Ovip:                  util.PtrString(""),
		EntryPort:             util.PtrInt32(int32(0)),
		ResourceGroupId:       util.PtrString(""),
		ReplicaType:           util.PtrString(""),
		IsInheritMasterAuthip: util.PtrBool(false),
		PurchaseCount:         util.PtrInt32(int32(0)),
		ResourceType:          util.PtrString(""),
		AutoRenewTimeUnit:     util.PtrString(""),
		AutoRenewTime:         util.PtrInt32(int32(0)),
	}
	result := &CreateReadOnlyInstanceResponse{}
	result, err := RDS_CLIENT.CreateReadOnlyInstance(createReadOnlyInstanceRequest)
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
func TestClient_QueryTaskList(t *testing.T) {
	queryTaskListRequest := &QueryTaskListRequest{
		PageSize:     util.PtrString(""),
		PageNo:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		InstanceName: util.PtrString(""),
		TaskId:       util.PtrInt32(int32(0)),
		TaskType:     util.PtrString(""),
		TaskStatus:   util.PtrString(""),
		StartTime:    util.PtrString(""),
		EndTime:      util.PtrString(""),
	}
	result := &QueryTaskListResponse{}
	result, err := RDS_CLIENT.QueryTaskList(queryTaskListRequest)
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
