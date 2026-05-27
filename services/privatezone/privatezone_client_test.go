package privatezone

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
	PRIVATEZONE_CLIENT *Client
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

	PRIVATEZONE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddRecord(t *testing.T) {
	addRecordRequest := &AddRecordRequest{
		ZoneId:          util.PtrString(""),
		ClientToken:     util.PtrString(""),
		Rr:              util.PtrString(""),
		Value:           util.PtrString(""),
		PrivatezoneType: util.PtrString(""),
		Priority:        util.PtrInt32(int32(0)),
		Ttl:             util.PtrInt32(int32(0)),
		Description:     util.PtrString(""),
	}
	result := &AddRecordResponse{}
	result, err := PRIVATEZONE_CLIENT.AddRecord(addRecordRequest)
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
func TestClient_BindVpc(t *testing.T) {
	bindVpcRequest := &BindVpcRequest{
		ZoneId:      util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Region:      util.PtrString(""),
		VpcIds:      []*string{},
	}
	err := PRIVATEZONE_CLIENT.BindVpc(bindVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreatePrivateZone(t *testing.T) {
	createPrivateZoneRequest := &CreatePrivateZoneRequest{
		ClientToken: util.PtrString(""),
		ZoneName:    util.PtrString(""),
	}
	result := &CreatePrivateZoneResponse{}
	result, err := PRIVATEZONE_CLIENT.CreatePrivateZone(createPrivateZoneRequest)
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
func TestClient_DeletePrivateZone(t *testing.T) {
	deletePrivateZoneRequest := &DeletePrivateZoneRequest{
		ZoneId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := PRIVATEZONE_CLIENT.DeletePrivateZone(deletePrivateZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRecord(t *testing.T) {
	deleteRecordRequest := &DeleteRecordRequest{
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := PRIVATEZONE_CLIENT.DeleteRecord(deleteRecordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableRecord(t *testing.T) {
	disableRecordRequest := &DisableRecordRequest{
		RecordId:    util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	result, err := PRIVATEZONE_CLIENT.DisableRecord(disableRecordRequest)
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
func TestClient_EnableRecord(t *testing.T) {
	enableRecordRequest := &EnableRecordRequest{
		RecordId:    util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	result, err := PRIVATEZONE_CLIENT.EnableRecord(enableRecordRequest)
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
func TestClient_GetPrivateZone(t *testing.T) {
	getPrivateZoneRequest := &GetPrivateZoneRequest{
		ZoneId: util.PtrString(""),
	}
	result := &GetPrivateZoneResponse{}
	result, err := PRIVATEZONE_CLIENT.GetPrivateZone(getPrivateZoneRequest)
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
func TestClient_ListPrivateZone(t *testing.T) {
	listPrivateZoneRequest := &ListPrivateZoneRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListPrivateZoneResponse{}
	result, err := PRIVATEZONE_CLIENT.ListPrivateZone(listPrivateZoneRequest)
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
		ZoneId:     util.PtrString(""),
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Rr:         util.PtrString(""),
		SearchMode: util.PtrString(""),
		Type:       util.PtrString(""),
		Value:      util.PtrString(""),
	}
	result := &ListRecordResponse{}
	result, err := PRIVATEZONE_CLIENT.ListRecord(listRecordRequest)
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
func TestClient_UnbindVpc(t *testing.T) {
	unbindVpcRequest := &UnbindVpcRequest{
		ZoneId:      util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Region:      util.PtrString(""),
		VpcIds:      []*string{},
	}
	err := PRIVATEZONE_CLIENT.UnbindVpc(unbindVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRecord(t *testing.T) {
	updateRecordRequest := &UpdateRecordRequest{
		RecordId:        util.PtrString(""),
		ClientToken:     util.PtrString(""),
		Rr:              util.PtrString(""),
		Value:           util.PtrString(""),
		PrivatezoneType: util.PtrString(""),
		Ttl:             util.PtrInt32(int32(0)),
		Priority:        util.PtrInt32(int32(0)),
		Description:     util.PtrString(""),
	}
	err := PRIVATEZONE_CLIENT.UpdateRecord(updateRecordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
