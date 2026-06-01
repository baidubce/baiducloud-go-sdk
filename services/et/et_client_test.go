package et

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
	ET_CLIENT *Client
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

	ET_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_ApplyPhysicalDedicatedLine(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	AutoRenew := &Reservation{
		ReservationLength:   util.PtrInt32(int32(0)),
		ReservationTimeUnit: util.PtrString(""),
	}
	applyPhysicalDedicatedLineRequest := &ApplyPhysicalDedicatedLineRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Isp:         util.PtrString(""),
		IntfType:    util.PtrString(""),
		ApType:      util.PtrString(""),
		ApAddr:      util.PtrString(""),
		LinkDelay:   util.PtrInt32(int32(0)),
		UserName:    util.PtrString(""),
		UserPhone:   util.PtrString(""),
		UserEmail:   util.PtrString(""),
		UserIdc:     util.PtrString(""),
		Billing:     Billing,
		AutoRenew:   AutoRenew,
		Tags:        []*TagModel{},
	}
	result := &ApplyPhysicalDedicatedLineResponse{}
	result, err := ET_CLIENT.ApplyPhysicalDedicatedLine(applyPhysicalDedicatedLineRequest)
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
func TestClient_AssociatedDedicatedChannel(t *testing.T) {
	associatedDedicatedChannelRequest := &AssociatedDedicatedChannelRequest{
		EtId:           util.PtrString(""),
		EtChannelId:    util.PtrString(""),
		ClientToken:    util.PtrString(""),
		ExtraChannelId: util.PtrString(""),
	}
	err := ET_CLIENT.AssociatedDedicatedChannel(associatedDedicatedChannelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateDedicatedChannel(t *testing.T) {
	createDedicatedChannelRequest := &CreateDedicatedChannelRequest{
		EtId:                util.PtrString(""),
		ClientToken:         util.PtrString(""),
		AuthorizedUsers:     []*string{},
		Description:         util.PtrString(""),
		BaiduAddress:        util.PtrString(""),
		Name:                util.PtrString(""),
		Networks:            []*string{},
		CustomerAddress:     util.PtrString(""),
		RouteType:           util.PtrString(""),
		VlanId:              util.PtrInt32(int32(0)),
		BgpAsn:              util.PtrString(""),
		BgpKey:              util.PtrString(""),
		EnableIpv6:          util.PtrInt32(int32(0)),
		BaiduIpv6Address:    util.PtrString(""),
		CustomerIpv6Address: util.PtrString(""),
		Ipv6Networks:        []*string{},
		Tags:                []*TagModel{},
	}
	result := &CreateDedicatedChannelResponse{}
	result, err := ET_CLIENT.CreateDedicatedChannel(createDedicatedChannelRequest)
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
func TestClient_CreateDedicatedChannelBfd(t *testing.T) {
	createDedicatedChannelBfdRequest := &CreateDedicatedChannelBfdRequest{
		EtId:             util.PtrString(""),
		EtChannelId:      util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SendInterval:     util.PtrInt32(int32(0)),
		ReceivInterval:   util.PtrInt32(int32(0)),
		DetectMultiplier: util.PtrInt32(int32(0)),
	}
	err := ET_CLIENT.CreateDedicatedChannelBfd(createDedicatedChannelBfdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateDedicatedChannelRouteParameters(t *testing.T) {
	createDedicatedChannelRouteParametersRequest := &CreateDedicatedChannelRouteParametersRequest{
		EtId:         util.PtrString(""),
		EtChannelId:  util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Networks:     []*string{},
		Ipv6Networks: []*string{},
		RouteType:    util.PtrString(""),
	}
	err := ET_CLIENT.CreateDedicatedChannelRouteParameters(createDedicatedChannelRouteParametersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateDedicatedChannelRouteRules(t *testing.T) {
	createDedicatedChannelRouteRulesRequest := &CreateDedicatedChannelRouteRulesRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpVersion:   util.PtrInt32(int32(0)),
		DestAddress: util.PtrString(""),
		NexthopType: util.PtrString(""),
		NexthopId:   util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateDedicatedChannelRouteRulesResponse{}
	result, err := ET_CLIENT.CreateDedicatedChannelRouteRules(createDedicatedChannelRouteRulesRequest)
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
func TestClient_CreateDedicatedChannelUserObject(t *testing.T) {
	createDedicatedChannelUserObjectRequest := &CreateDedicatedChannelUserObjectRequest{
		EtId:            util.PtrString(""),
		EtChannelId:     util.PtrString(""),
		ClientToken:     util.PtrString(""),
		AuthorizedUsers: []*string{},
	}
	err := ET_CLIENT.CreateDedicatedChannelUserObject(createDedicatedChannelUserObjectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDedicatedChannel(t *testing.T) {
	deleteDedicatedChannelRequest := &DeleteDedicatedChannelRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := ET_CLIENT.DeleteDedicatedChannel(deleteDedicatedChannelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDedicatedChannelBfd(t *testing.T) {
	deleteDedicatedChannelBfdRequest := &DeleteDedicatedChannelBfdRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := ET_CLIENT.DeleteDedicatedChannelBfd(deleteDedicatedChannelBfdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDedicatedChannelRouteRules(t *testing.T) {
	deleteDedicatedChannelRouteRulesRequest := &DeleteDedicatedChannelRouteRulesRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		RouteRuleId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := ET_CLIENT.DeleteDedicatedChannelRouteRules(deleteDedicatedChannelRouteRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableDedicatedChannelIpv6(t *testing.T) {
	disableDedicatedChannelIpv6Request := &DisableDedicatedChannelIpv6Request{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := ET_CLIENT.DisableDedicatedChannelIpv6(disableDedicatedChannelIpv6Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableDedicatedChannelIpv6(t *testing.T) {
	enableDedicatedChannelIpv6Request := &EnableDedicatedChannelIpv6Request{
		EtId:                util.PtrString(""),
		EtChannelId:         util.PtrString(""),
		ClientToken:         util.PtrString(""),
		BaiduIpv6Address:    util.PtrString(""),
		CustomerIpv6Address: util.PtrString(""),
		Ipv6Networks:        []*string{},
	}
	err := ET_CLIENT.EnableDedicatedChannelIpv6(enableDedicatedChannelIpv6Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryDedicatedChannel(t *testing.T) {
	queryDedicatedChannelRequest := &QueryDedicatedChannelRequest{
		EtId:        util.PtrString(""),
		ClientToken: util.PtrString(""),
		EtChannelId: util.PtrString(""),
	}
	result := &QueryDedicatedChannelResponse{}
	result, err := ET_CLIENT.QueryDedicatedChannel(queryDedicatedChannelRequest)
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
func TestClient_QueryDedicatedChannelRouteRules(t *testing.T) {
	queryDedicatedChannelRouteRulesRequest := &QueryDedicatedChannelRouteRulesRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
		DestAddress: util.PtrString(""),
	}
	result := &QueryDedicatedChannelRouteRulesResponse{}
	result, err := ET_CLIENT.QueryDedicatedChannelRouteRules(queryDedicatedChannelRouteRulesRequest)
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
func TestClient_QueryDedicatedLineDetail(t *testing.T) {
	queryDedicatedLineDetailRequest := &QueryDedicatedLineDetailRequest{
		DcphyId: util.PtrString(""),
	}
	result := &QueryDedicatedLineDetailResponse{}
	result, err := ET_CLIENT.QueryDedicatedLineDetail(queryDedicatedLineDetailRequest)
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
func TestClient_QueryDedicatedLines(t *testing.T) {
	queryDedicatedLinesRequest := &QueryDedicatedLinesRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		Status:  util.PtrString(""),
	}
	result := &QueryDedicatedLinesResponse{}
	result, err := ET_CLIENT.QueryDedicatedLines(queryDedicatedLinesRequest)
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
func TestClient_RemoveDedicatedChannelRouteParameters(t *testing.T) {
	removeDedicatedChannelRouteParametersRequest := &RemoveDedicatedChannelRouteParametersRequest{
		EtId:         util.PtrString(""),
		EtChannelId:  util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Networks:     []*string{},
		Ipv6Networks: []*string{},
		RouteType:    util.PtrString(""),
	}
	err := ET_CLIENT.RemoveDedicatedChannelRouteParameters(removeDedicatedChannelRouteParametersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveDedicatedChannelUserObject(t *testing.T) {
	removeDedicatedChannelUserObjectRequest := &RemoveDedicatedChannelUserObjectRequest{
		EtId:            util.PtrString(""),
		EtChannelId:     util.PtrString(""),
		ClientToken:     util.PtrString(""),
		AuthorizedUsers: []*string{},
	}
	err := ET_CLIENT.RemoveDedicatedChannelUserObject(removeDedicatedChannelUserObjectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResubmitDedicatedChannel(t *testing.T) {
	resubmitDedicatedChannelRequest := &ResubmitDedicatedChannelRequest{
		EtId:                util.PtrString(""),
		EtChannelId:         util.PtrString(""),
		ClientToken:         util.PtrString(""),
		AuthorizedUsers:     []*string{},
		Description:         util.PtrString(""),
		BaiduAddress:        util.PtrString(""),
		Name:                util.PtrString(""),
		Networks:            []*string{},
		CustomerAddress:     util.PtrString(""),
		RouteType:           util.PtrString(""),
		VlanId:              util.PtrInt32(int32(0)),
		EnableIpv6:          util.PtrInt32(int32(0)),
		BaiduIpv6Address:    util.PtrString(""),
		CustomerIpv6Address: util.PtrString(""),
		Ipv6Networks:        []*string{},
	}
	err := ET_CLIENT.ResubmitDedicatedChannel(resubmitDedicatedChannelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnrelatedDedicatedLineChannel(t *testing.T) {
	unrelatedDedicatedLineChannelRequest := &UnrelatedDedicatedLineChannelRequest{
		EtId:           util.PtrString(""),
		EtChannelId:    util.PtrString(""),
		ClientToken:    util.PtrString(""),
		ExtraChannelId: util.PtrString(""),
	}
	err := ET_CLIENT.UnrelatedDedicatedLineChannel(unrelatedDedicatedLineChannelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDedicatedChannel(t *testing.T) {
	updateDedicatedChannelRequest := &UpdateDedicatedChannelRequest{
		EtId:          util.PtrString(""),
		EtChannelId:   util.PtrString(""),
		ClientToken:   util.PtrString(""),
		Name:          util.PtrString(""),
		Description:   util.PtrString(""),
		BgpRouteLimit: util.PtrInt32(int32(0)),
	}
	err := ET_CLIENT.UpdateDedicatedChannel(updateDedicatedChannelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDedicatedChannelBfd(t *testing.T) {
	updateDedicatedChannelBfdRequest := &UpdateDedicatedChannelBfdRequest{
		EtId:             util.PtrString(""),
		EtChannelId:      util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SendInterval:     util.PtrInt32(int32(0)),
		ReceivInterval:   util.PtrInt32(int32(0)),
		DetectMultiplier: util.PtrInt32(int32(0)),
	}
	err := ET_CLIENT.UpdateDedicatedChannelBfd(updateDedicatedChannelBfdRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDedicatedChannelRouteRules(t *testing.T) {
	updateDedicatedChannelRouteRulesRequest := &UpdateDedicatedChannelRouteRulesRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		RouteRuleId: util.PtrString(""),
		ClientToken: util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := ET_CLIENT.UpdateDedicatedChannelRouteRules(updateDedicatedChannelRouteRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePhysicalDedicatedLine(t *testing.T) {
	updatePhysicalDedicatedLineRequest := &UpdatePhysicalDedicatedLineRequest{
		DcphyId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		UserName:    util.PtrString(""),
		UserPhone:   util.PtrString(""),
		UserEmail:   util.PtrString(""),
		LinkDelay:   util.PtrInt32(int32(0)),
	}
	err := ET_CLIENT.UpdatePhysicalDedicatedLine(updatePhysicalDedicatedLineRequest)
	ExpectEqual(t.Errorf, nil, err)
}
