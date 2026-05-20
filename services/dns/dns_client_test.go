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

func TestClient_AddDomainName(t *testing.T) {
	addDomainNameRequest := &AddDomainNameRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
	}
	err := DNS_CLIENT.AddDomainName(addDomainNameRequest)
	ExpectEqual(t.Errorf, nil, err)
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
func TestClient_AddParsingRecords(t *testing.T) {
	addParsingRecordsRequest := &AddParsingRecordsRequest{
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
	err := DNS_CLIENT.AddParsingRecords(addParsingRecordsRequest)
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
func TestClient_DeleteParsingRecords(t *testing.T) {
	deleteParsingRecordsRequest := &DeleteParsingRecordsRequest{
		ZoneName:    util.PtrString(""),
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.DeleteParsingRecords(deleteParsingRecordsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DomainNameRenewal(t *testing.T) {
	domainNameRenewalRequest := &DomainNameRenewalRequest{
		Name:        util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     []*BillingForRenew{},
	}
	err := DNS_CLIENT.DomainNameRenewal(domainNameRenewalRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyParsingRecords(t *testing.T) {
	modifyParsingRecordsRequest := &ModifyParsingRecordsRequest{
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
	err := DNS_CLIENT.ModifyParsingRecords(modifyParsingRecordsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTheParsingRecordStatus(t *testing.T) {
	modifyTheParsingRecordStatusRequest := &ModifyTheParsingRecordStatusRequest{
		ZoneName:    util.PtrString(""),
		RecordId:    util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.ModifyTheParsingRecordStatus(modifyTheParsingRecordStatusRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PurchaseAPaidDomainName(t *testing.T) {
	purchaseAPaidDomainNameRequest := &PurchaseAPaidDomainNameRequest{
		ClientToken:    util.PtrString(""),
		Names:          []*string{},
		ProductVersion: util.PtrString(""),
		Billing:        []*Billing{},
	}
	err := DNS_CLIENT.PurchaseAPaidDomainName(purchaseAPaidDomainNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryAndParseRecordList(t *testing.T) {
	queryAndParseRecordListRequest := &QueryAndParseRecordListRequest{
		ZoneName: util.PtrString(""),
		Rr:       util.PtrString(""),
		Id:       util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
	}
	result := &QueryAndParseRecordListResponse{}
	result, err := DNS_CLIENT.QueryAndParseRecordList(queryAndParseRecordListRequest)
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
func TestClient_QueryDomainNameList(t *testing.T) {
	queryDomainNameListRequest := &QueryDomainNameListRequest{
		Name:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryDomainNameListResponse{}
	result, err := DNS_CLIENT.QueryDomainNameList(queryDomainNameListRequest)
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
func TestClient_QueryTheListOfLineGroups(t *testing.T) {
	queryTheListOfLineGroupsRequest := &QueryTheListOfLineGroupsRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfLineGroupsResponse{}
	result, err := DNS_CLIENT.QueryTheListOfLineGroups(queryTheListOfLineGroupsRequest)
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
func TestClient_RemoveDomainName(t *testing.T) {
	removeDomainNameRequest := &RemoveDomainNameRequest{
		ZoneName:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := DNS_CLIENT.RemoveDomainName(removeDomainNameRequest)
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
func TestClient_UpgradeTheFreeDomainNameToTheUniversalVersion(t *testing.T) {
	upgradeTheFreeDomainNameToTheUniversalVersionRequest := &UpgradeTheFreeDomainNameToTheUniversalVersionRequest{
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Names:       []*string{},
		Billing:     []*Billing{},
	}
	err := DNS_CLIENT.UpgradeTheFreeDomainNameToTheUniversalVersion(upgradeTheFreeDomainNameToTheUniversalVersionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
