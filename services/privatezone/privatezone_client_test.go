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

func TestClient_AddParsingRecords(t *testing.T) {
	addParsingRecordsRequest := &AddParsingRecordsRequest{
		ZoneId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Rr:          util.PtrString(""),
		Value:       util.PtrString(""),
		Type:        util.PtrString(""),
		Priority:    util.PtrInt32(int32(0)),
		Ttl:         util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
	}
	result := &AddParsingRecordsResponse{}
	result, err := PRIVATEZONE_CLIENT.AddParsingRecords(addParsingRecordsRequest)
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
func TestClient_AssociateVpc(t *testing.T) {
	associateVpcRequest := &AssociateVpcRequest{
		ZoneId:      util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Region:      util.PtrString(""),
		VpcIds:      []*string{},
	}
	err := PRIVATEZONE_CLIENT.AssociateVpc(associateVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAPrivateZone(t *testing.T) {
	createAPrivateZoneRequest := &CreateAPrivateZoneRequest{
		ClientToken: util.PtrString(""),
		ZoneName:    util.PtrString(""),
	}
	result := &CreateAPrivateZoneResponse{}
	result, err := PRIVATEZONE_CLIENT.CreateAPrivateZone(createAPrivateZoneRequest)
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
func TestClient_DeleteParsingRecords(t *testing.T) {
	deleteParsingRecordsRequest := &DeleteParsingRecordsRequest{
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := PRIVATEZONE_CLIENT.DeleteParsingRecords(deleteParsingRecordsRequest)
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
func TestClient_DisassociateVpc(t *testing.T) {
	disassociateVpcRequest := &DisassociateVpcRequest{
		ZoneId:      util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Region:      util.PtrString(""),
		VpcIds:      []*string{},
	}
	err := PRIVATEZONE_CLIENT.DisassociateVpc(disassociateVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyParsingRecords(t *testing.T) {
	modifyParsingRecordsRequest := &ModifyParsingRecordsRequest{
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
		Rr:          util.PtrString(""),
		Value:       util.PtrString(""),
		Type:        util.PtrString(""),
		Ttl:         util.PtrInt32(int32(0)),
		Priority:    util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
	}
	err := PRIVATEZONE_CLIENT.ModifyParsingRecords(modifyParsingRecordsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryAndParseRecordList(t *testing.T) {
	queryAndParseRecordListRequest := &QueryAndParseRecordListRequest{
		ZoneId:     util.PtrString(""),
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Rr:         util.PtrString(""),
		SearchMode: util.PtrString(""),
		Type:       util.PtrString(""),
		Value:      util.PtrString(""),
	}
	result := &QueryAndParseRecordListResponse{}
	result, err := PRIVATEZONE_CLIENT.QueryAndParseRecordList(queryAndParseRecordListRequest)
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
func TestClient_QueryTheListOfPrivateZones(t *testing.T) {
	queryTheListOfPrivateZonesRequest := &QueryTheListOfPrivateZonesRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfPrivateZonesResponse{}
	result, err := PRIVATEZONE_CLIENT.QueryTheListOfPrivateZones(queryTheListOfPrivateZonesRequest)
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
func TestClient_SearchForDetailsOfPrivatzone(t *testing.T) {
	searchForDetailsOfPrivatzoneRequest := &SearchForDetailsOfPrivatzoneRequest{
		ZoneId: util.PtrString(""),
	}
	result := &SearchForDetailsOfPrivatzoneResponse{}
	result, err := PRIVATEZONE_CLIENT.SearchForDetailsOfPrivatzone(searchForDetailsOfPrivatzoneRequest)
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
func TestClient_SetParsingRecordStatus(t *testing.T) {
	setParsingRecordStatusRequest := &SetParsingRecordStatusRequest{
		RecordId:    util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := PRIVATEZONE_CLIENT.SetParsingRecordStatus(setParsingRecordStatusRequest)
	ExpectEqual(t.Errorf, nil, err)
}
