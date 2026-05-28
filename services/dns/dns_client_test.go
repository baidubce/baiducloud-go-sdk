package dns

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
	DNS_CLIENT *Client
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

	DNS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddLineGroup(t *testing.T) {
	addLineGroupRequest := &AddLineGroupRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Lines:       []*string{},
	}
	err := DNS_CLIENT.AddLineGroup(addLineGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreatePaidZone(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	createPaidZoneRequest := &CreatePaidZoneRequest{
		ClientToken:    util.PtrString(""),
		Names:          []*string{},
		ProductVersion: util.PtrString(""),
		Billing:        Billing,
	}
	err := DNS_CLIENT.CreatePaidZone(createPaidZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateRecord(t *testing.T) {
	createRecordRequest := &CreateRecordRequest{
		ZoneName:    util.PtrString(""),
		ClientToken: util.PtrString(""),
		Rr:          util.PtrString(""),
		DnsType:     util.PtrString(""),
		Value:       util.PtrString(""),
		Ttl:         util.PtrInt32(int32(0)),
		Line:        util.PtrString(""),
		Description: util.PtrString(""),
		Priority:    util.PtrInt32(int32(0)),
	}
	err := DNS_CLIENT.CreateRecord(createRecordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateZone(t *testing.T) {
	createZoneRequest := &CreateZoneRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
	}
	err := DNS_CLIENT.CreateZone(createZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteLineGroup(t *testing.T) {
	deleteLineGroupRequest := &DeleteLineGroupRequest{
		LineId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.DeleteLineGroup(deleteLineGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRecord(t *testing.T) {
	deleteRecordRequest := &DeleteRecordRequest{
		ZoneName:    util.PtrString(""),
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.DeleteRecord(deleteRecordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteZone(t *testing.T) {
	deleteZoneRequest := &DeleteZoneRequest{
		ZoneName:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.DeleteZone(deleteZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListLineGroup(t *testing.T) {
	listLineGroupRequest := &ListLineGroupRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListLineGroupResponse{}
	result, err := DNS_CLIENT.ListLineGroup(listLineGroupRequest)
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
func TestClient_ListRecord(t *testing.T) {
	listRecordRequest := &ListRecordRequest{
		ZoneName: util.PtrString(""),
		Rr:       util.PtrString(""),
		Id:       util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
	}
	result := &ListRecordResponse{}
	result, err := DNS_CLIENT.ListRecord(listRecordRequest)
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
func TestClient_ListZone(t *testing.T) {
	listZoneRequest := &ListZoneRequest{
		Name:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListZoneResponse{}
	result, err := DNS_CLIENT.ListZone(listZoneRequest)
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
func TestClient_RenewZone(t *testing.T) {
	Billing := &BillingForRenew{
		Reservation: &Reservation{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	renewZoneRequest := &RenewZoneRequest{
		Name:        util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     Billing,
	}
	err := DNS_CLIENT.RenewZone(renewZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateLineGroup(t *testing.T) {
	updateLineGroupRequest := &UpdateLineGroupRequest{
		LineId:      util.PtrInt32(int32(0)),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Lines:       []*string{},
	}
	err := DNS_CLIENT.UpdateLineGroup(updateLineGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRecord(t *testing.T) {
	updateRecordRequest := &UpdateRecordRequest{
		ZoneName:    util.PtrString(""),
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
		Rr:          util.PtrString(""),
		DnsType:     util.PtrString(""),
		Value:       util.PtrString(""),
		Ttl:         util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
		Priority:    util.PtrInt32(int32(0)),
	}
	err := DNS_CLIENT.UpdateRecord(updateRecordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRecordDisable(t *testing.T) {
	updateRecordDisableRequest := &UpdateRecordDisableRequest{
		ZoneName:    util.PtrString(""),
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.UpdateRecordDisable(updateRecordDisableRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRecordEnable(t *testing.T) {
	updateRecordEnableRequest := &UpdateRecordEnableRequest{
		ZoneName:    util.PtrString(""),
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.UpdateRecordEnable(updateRecordEnableRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpgradeZone(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	upgradeZoneRequest := &UpgradeZoneRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Names:       []*string{},
		Billing:     Billing,
	}
	err := DNS_CLIENT.UpgradeZone(upgradeZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
