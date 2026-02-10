package eip

import (
	"bytes"
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
	EIP_CLIENT *Client
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

	EIP_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddTbspAreaBlocking(t *testing.T) {
	addTbspAreaBlockingRequest := &AddTbspAreaBlockingRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		Ip:          util.PtrString(""),
		BlockTime:   util.PtrInt32(int32(0)),
		BlockType:   util.PtrString(""),
	}
	err := EIP_CLIENT.AddTbspAreaBlocking(addTbspAreaBlockingRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddTbspIpWhitelist(t *testing.T) {
	addTbspIpWhitelistRequest := &AddTbspIpWhitelistRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		Ip:          util.PtrString(""),
		IpCidrList:  []*string{},
	}
	err := EIP_CLIENT.AddTbspIpWhitelist(addTbspIpWhitelistRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddTbspProtocolBlocking(t *testing.T) {
	addTbspProtocolBlockingRequest := &AddTbspProtocolBlockingRequest{
		Id:               util.PtrString(""),
		ClientToken:      util.PtrString(""),
		Ip:               util.PtrString(""),
		ProtocolPortList: []*TbspProtocolPortModel{},
	}
	err := EIP_CLIENT.AddTbspProtocolBlocking(addTbspProtocolBlockingRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BandwidthPackageInquiry(t *testing.T) {
	bandwidthPackageInquiryRequest := &BandwidthPackageInquiryRequest{
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Count:           util.PtrInt32(int32(0)),
		Type:            util.PtrString(""),
	}
	result := &BandwidthPackageInquiryResponse{}
	result, err := EIP_CLIENT.BandwidthPackageInquiry(bandwidthPackageInquiryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindTbspProtectionObject(t *testing.T) {
	bindTbspProtectionObjectRequest := &BindTbspProtectionObjectRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpList:      []*string{},
	}
	err := EIP_CLIENT.BindTbspProtectionObject(bindTbspProtectionObjectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateTbsp(t *testing.T) {
	createTbspRequest := &CreateTbspRequest{
		ClientToken:         util.PtrString(""),
		Name:                util.PtrString(""),
		LineType:            util.PtrString(""),
		IpCapacity:          util.PtrInt32(int32(0)),
		ReservationLength:   util.PtrInt32(int32(0)),
		ReservationTimeUnit: util.PtrString(""),
		AutoRenewTime:       util.PtrInt32(int32(0)),
		AutoRenewTimeUnit:   util.PtrString(""),
	}
	result := &CreateTbspResponse{}
	result, err := EIP_CLIENT.CreateTbsp(createTbspRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DetailTbsp(t *testing.T) {
	detailTbspRequest := &DetailTbspRequest{
		Id: util.PtrString(""),
	}
	result := &DetailTbspResponse{}
	result, err := EIP_CLIENT.DetailTbsp(detailTbspRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableTbspIpClean(t *testing.T) {
	disableTbspIpCleanRequest := &DisableTbspIpCleanRequest{
		Id:              util.PtrString(""),
		ClientToken:     util.PtrString(""),
		Ip:              util.PtrString(""),
		TurnOffDuration: util.PtrInt32(int32(0)),
	}
	err := EIP_CLIENT.DisableTbspIpClean(disableTbspIpCleanRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EipInquiry(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	eipInquiryRequest := &EipInquiryRequest{
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Count:           util.PtrInt32(int32(0)),
		PurchaseType:    util.PtrString(""),
		Billing:         Billing,
	}
	result := &EipInquiryResponse{}
	result, err := EIP_CLIENT.EipInquiry(eipInquiryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableTbspIpClean(t *testing.T) {
	enableTbspIpCleanRequest := &EnableTbspIpCleanRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		Ip:          util.PtrString(""),
	}
	err := EIP_CLIENT.EnableTbspIpClean(enableTbspIpCleanRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTbsp(t *testing.T) {
	listTbspRequest := &ListTbspRequest{
		Id:      util.PtrString(""),
		Name:    util.PtrString(""),
		Status:  util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListTbspResponse{}
	result, err := EIP_CLIENT.ListTbsp(listTbspRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTbspAreaBlocking(t *testing.T) {
	listTbspAreaBlockingRequest := &ListTbspAreaBlockingRequest{
		Id: util.PtrString(""),
		Ip: util.PtrString(""),
	}
	result := &ListTbspAreaBlockingResponse{}
	result, err := EIP_CLIENT.ListTbspAreaBlocking(listTbspAreaBlockingRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTbspIpClean(t *testing.T) {
	listTbspIpCleanRequest := &ListTbspIpCleanRequest{
		Id:      util.PtrString(""),
		Ip:      util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListTbspIpCleanResponse{}
	result, err := EIP_CLIENT.ListTbspIpClean(listTbspIpCleanRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTbspIpWhitelist(t *testing.T) {
	listTbspIpWhitelistRequest := &ListTbspIpWhitelistRequest{
		Id:      util.PtrString(""),
		Ip:      util.PtrString(""),
		IpCidr:  util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListTbspIpWhitelistResponse{}
	result, err := EIP_CLIENT.ListTbspIpWhitelist(listTbspIpWhitelistRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTbspProtocolBlocking(t *testing.T) {
	listTbspProtocolBlockingRequest := &ListTbspProtocolBlockingRequest{
		Id: util.PtrString(""),
		Ip: util.PtrString(""),
	}
	result := &ListTbspProtocolBlockingResponse{}
	result, err := EIP_CLIENT.ListTbspProtocolBlocking(listTbspProtocolBlockingRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTbspIpCleanThreshold(t *testing.T) {
	modifyTbspIpCleanThresholdRequest := &ModifyTbspIpCleanThresholdRequest{
		Id:            util.PtrString(""),
		ClientToken:   util.PtrString(""),
		Ip:            util.PtrString(""),
		ThresholdType: util.PtrString(""),
		CleanMbps:     util.PtrInt32(int32(0)),
		CleanPps:      util.PtrInt32(int32(0)),
	}
	err := EIP_CLIENT.ModifyTbspIpCleanThreshold(modifyTbspIpCleanThresholdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTbspIpProtectLevel(t *testing.T) {
	modifyTbspIpProtectLevelRequest := &ModifyTbspIpProtectLevelRequest{
		Id:           util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Ip:           util.PtrString(""),
		ProtectLevel: util.PtrInt32(int32(0)),
	}
	err := EIP_CLIENT.ModifyTbspIpProtectLevel(modifyTbspIpProtectLevelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveTbspAreaBlocking(t *testing.T) {
	removeTbspAreaBlockingRequest := &RemoveTbspAreaBlockingRequest{
		Id:          util.PtrString(""),
		Ip:          util.PtrString(""),
		BlockType:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := EIP_CLIENT.RemoveTbspAreaBlocking(removeTbspAreaBlockingRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveTbspIpWhitelist(t *testing.T) {
	removeTbspIpWhitelistRequest := &RemoveTbspIpWhitelistRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		Ip:          util.PtrString(""),
		WhitelistId: util.PtrString(""),
	}
	err := EIP_CLIENT.RemoveTbspIpWhitelist(removeTbspIpWhitelistRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveTbspProtocolBlocking(t *testing.T) {
	removeTbspProtocolBlockingRequest := &RemoveTbspProtocolBlockingRequest{
		Id:          util.PtrString(""),
		Ip:          util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := EIP_CLIENT.RemoveTbspProtocolBlocking(removeTbspProtocolBlockingRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RenewTbsp(t *testing.T) {
	renewTbspRequest := &RenewTbspRequest{
		Id:            util.PtrString(""),
		ClientToken:   util.PtrString(""),
		RenewTime:     util.PtrInt32(int32(0)),
		RenewTimeUnit: util.PtrString(""),
	}
	err := EIP_CLIENT.RenewTbsp(renewTbspRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeTbsp(t *testing.T) {
	resizeTbspRequest := &ResizeTbspRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpCapacity:  util.PtrInt32(int32(0)),
	}
	err := EIP_CLIENT.ResizeTbsp(resizeTbspRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SharedBandwidthInquiry(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	sharedBandwidthInquiryRequest := &SharedBandwidthInquiryRequest{
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Count:           util.PtrInt32(int32(0)),
		IpNum:           util.PtrInt32(int32(0)),
		Billing:         Billing,
	}
	result := &SharedBandwidthInquiryResponse{}
	result, err := EIP_CLIENT.SharedBandwidthInquiry(sharedBandwidthInquiryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SharedDataPackageInquiry(t *testing.T) {
	sharedDataPackageInquiryRequest := &SharedDataPackageInquiryRequest{
		ClientToken:       util.PtrString(""),
		ReservationLength: util.PtrInt32(int32(0)),
		Capacity:          util.PtrString(""),
		DeductPolicy:      util.PtrString(""),
		PackageType:       util.PtrString(""),
	}
	result := &SharedDataPackageInquiryResponse{}
	result, err := EIP_CLIENT.SharedDataPackageInquiry(sharedDataPackageInquiryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindTbspProtectionObject(t *testing.T) {
	unbindTbspProtectionObjectRequest := &UnbindTbspProtectionObjectRequest{
		Id:          util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpList:      []*string{},
	}
	err := EIP_CLIENT.UnbindTbspProtectionObject(unbindTbspProtectionObjectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
