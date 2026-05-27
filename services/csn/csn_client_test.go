package csn

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
	CSN_CLIENT *Client
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

	CSN_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddRouteRule(t *testing.T) {
	addRouteRuleRequest := &AddRouteRuleRequest{
		CsnRtId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		AttachId:    util.PtrString(""),
		DestAddress: util.PtrString(""),
		RouteType:   util.PtrString(""),
	}
	err := CSN_CLIENT.AddRouteRule(addRouteRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AttachCsnInstance(t *testing.T) {
	attachCsnInstanceRequest := &AttachCsnInstanceRequest{
		CsnId:             util.PtrString(""),
		Action:            util.PtrString(""),
		ClientToken:       util.PtrString(""),
		InstanceType:      util.PtrString(""),
		InstanceId:        util.PtrString(""),
		InstanceRegion:    util.PtrString(""),
		InstanceAccountId: util.PtrString(""),
	}
	err := CSN_CLIENT.AttachCsnInstance(attachCsnInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindCsnBp(t *testing.T) {
	bindCsnBpRequest := &BindCsnBpRequest{
		CsnBpId:     util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		CsnId:       util.PtrString(""),
	}
	err := CSN_CLIENT.BindCsnBp(bindCsnBpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAssociationRelation(t *testing.T) {
	createAssociationRelationRequest := &CreateAssociationRelationRequest{
		CsnRtId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		AttachId:    util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := CSN_CLIENT.CreateAssociationRelation(createAssociationRelationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateCsn(t *testing.T) {
	createCsnRequest := &CreateCsnRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Tags:        []*TagModel{},
	}
	result := &CreateCsnResponse{}
	result, err := CSN_CLIENT.CreateCsn(createCsnRequest)
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
func TestClient_CreateCsnBp(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createCsnBpRequest := &CreateCsnBpRequest{
		ClientToken:   util.PtrString(""),
		Name:          util.PtrString(""),
		InterworkType: util.PtrString(""),
		Bandwidth:     util.PtrInt32(int32(0)),
		GeographicA:   util.PtrString(""),
		GeographicB:   util.PtrString(""),
		Billing:       Billing,
		Tags:          []*TagModel{},
	}
	result := &CreateCsnBpResponse{}
	result, err := CSN_CLIENT.CreateCsnBp(createCsnBpRequest)
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
func TestClient_CreateRegionBandwidth(t *testing.T) {
	createRegionBandwidthRequest := &CreateRegionBandwidthRequest{
		CsnBpId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		LocalRegion: util.PtrString(""),
		PeerRegion:  util.PtrString(""),
		Bandwidth:   util.PtrInt32(int32(0)),
	}
	err := CSN_CLIENT.CreateRegionBandwidth(createRegionBandwidthRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateStudyRelation(t *testing.T) {
	createStudyRelationRequest := &CreateStudyRelationRequest{
		CsnRtId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		AttachId:    util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := CSN_CLIENT.CreateStudyRelation(createStudyRelationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAssociationRelation(t *testing.T) {
	deleteAssociationRelationRequest := &DeleteAssociationRelationRequest{
		CsnRtId:     util.PtrString(""),
		AttachId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := CSN_CLIENT.DeleteAssociationRelation(deleteAssociationRelationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCsn(t *testing.T) {
	deleteCsnRequest := &DeleteCsnRequest{
		CsnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := CSN_CLIENT.DeleteCsn(deleteCsnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCsnBp(t *testing.T) {
	deleteCsnBpRequest := &DeleteCsnBpRequest{
		CsnBpId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := CSN_CLIENT.DeleteCsnBp(deleteCsnBpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRegionBandwidth(t *testing.T) {
	deleteRegionBandwidthRequest := &DeleteRegionBandwidthRequest{
		CsnBpId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		LocalRegion: util.PtrString(""),
		PeerRegion:  util.PtrString(""),
	}
	err := CSN_CLIENT.DeleteRegionBandwidth(deleteRegionBandwidthRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRouteRule(t *testing.T) {
	deleteRouteRuleRequest := &DeleteRouteRuleRequest{
		CsnRtId:     util.PtrString(""),
		CsnRtRuleId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := CSN_CLIENT.DeleteRouteRule(deleteRouteRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteStudyRelation(t *testing.T) {
	deleteStudyRelationRequest := &DeleteStudyRelationRequest{
		CsnRtId:     util.PtrString(""),
		AttachId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := CSN_CLIENT.DeleteStudyRelation(deleteStudyRelationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DetachCsnInstance(t *testing.T) {
	detachCsnInstanceRequest := &DetachCsnInstanceRequest{
		CsnId:             util.PtrString(""),
		Action:            util.PtrString(""),
		ClientToken:       util.PtrString(""),
		InstanceType:      util.PtrString(""),
		InstanceId:        util.PtrString(""),
		InstanceRegion:    util.PtrString(""),
		InstanceAccountId: util.PtrString(""),
	}
	err := CSN_CLIENT.DetachCsnInstance(detachCsnInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryAssociationRelation(t *testing.T) {
	queryAssociationRelationRequest := &QueryAssociationRelationRequest{
		CsnRtId: util.PtrString(""),
	}
	result := &QueryAssociationRelationResponse{}
	result, err := CSN_CLIENT.QueryAssociationRelation(queryAssociationRelationRequest)
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
func TestClient_QueryCsnBpDetail(t *testing.T) {
	queryCsnBpDetailRequest := &QueryCsnBpDetailRequest{
		CsnBpId: util.PtrString(""),
	}
	result := &QueryCsnBpDetailResponse{}
	result, err := CSN_CLIENT.QueryCsnBpDetail(queryCsnBpDetailRequest)
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
func TestClient_QueryCsnBpList(t *testing.T) {
	queryCsnBpListRequest := &QueryCsnBpListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryCsnBpListResponse{}
	result, err := CSN_CLIENT.QueryCsnBpList(queryCsnBpListRequest)
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
func TestClient_QueryCsnBpPrice(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	queryCsnBpPriceRequest := &QueryCsnBpPriceRequest{
		Name:        util.PtrString(""),
		Bandwidth:   util.PtrInt32(int32(0)),
		GeographicA: util.PtrString(""),
		GeographicB: util.PtrString(""),
		Billing:     Billing,
	}
	result := &QueryCsnBpPriceResponse{}
	result, err := CSN_CLIENT.QueryCsnBpPrice(queryCsnBpPriceRequest)
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
func TestClient_QueryCsnDetail(t *testing.T) {
	queryCsnDetailRequest := &QueryCsnDetailRequest{
		CsnId: util.PtrString(""),
	}
	result := &QueryCsnDetailResponse{}
	result, err := CSN_CLIENT.QueryCsnDetail(queryCsnDetailRequest)
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
func TestClient_QueryCsnInstance(t *testing.T) {
	queryCsnInstanceRequest := &QueryCsnInstanceRequest{
		CsnId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryCsnInstanceResponse{}
	result, err := CSN_CLIENT.QueryCsnInstance(queryCsnInstanceRequest)
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
func TestClient_QueryCsnList(t *testing.T) {
	queryCsnListRequest := &QueryCsnListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryCsnListResponse{}
	result, err := CSN_CLIENT.QueryCsnList(queryCsnListRequest)
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
func TestClient_QueryRegionBandwidth(t *testing.T) {
	queryRegionBandwidthRequest := &QueryRegionBandwidthRequest{
		CsnBpId: util.PtrString(""),
	}
	result := &QueryRegionBandwidthResponse{}
	result, err := CSN_CLIENT.QueryRegionBandwidth(queryRegionBandwidthRequest)
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
func TestClient_QueryRegionBandwidthByCsn(t *testing.T) {
	queryRegionBandwidthByCsnRequest := &QueryRegionBandwidthByCsnRequest{
		CsnId: util.PtrString(""),
	}
	result := &QueryRegionBandwidthByCsnResponse{}
	result, err := CSN_CLIENT.QueryRegionBandwidthByCsn(queryRegionBandwidthByCsnRequest)
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
func TestClient_QueryRouteRule(t *testing.T) {
	queryRouteRuleRequest := &QueryRouteRuleRequest{
		CsnRtId: util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryRouteRuleResponse{}
	result, err := CSN_CLIENT.QueryRouteRule(queryRouteRuleRequest)
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
func TestClient_QueryRouteTableList(t *testing.T) {
	queryRouteTableListRequest := &QueryRouteTableListRequest{
		CsnId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryRouteTableListResponse{}
	result, err := CSN_CLIENT.QueryRouteTableList(queryRouteTableListRequest)
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
func TestClient_QueryStudyRelation(t *testing.T) {
	queryStudyRelationRequest := &QueryStudyRelationRequest{
		CsnRtId: util.PtrString(""),
	}
	result := &QueryStudyRelationResponse{}
	result, err := CSN_CLIENT.QueryStudyRelation(queryStudyRelationRequest)
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
func TestClient_QueryTgwList(t *testing.T) {
	queryTgwListRequest := &QueryTgwListRequest{
		CsnId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryTgwListResponse{}
	result, err := CSN_CLIENT.QueryTgwList(queryTgwListRequest)
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
func TestClient_QueryTgwRoute(t *testing.T) {
	queryTgwRouteRequest := &QueryTgwRouteRequest{
		CsnId:   util.PtrString(""),
		TgwId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryTgwRouteResponse{}
	result, err := CSN_CLIENT.QueryTgwRoute(queryTgwRouteRequest)
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
func TestClient_ResizeCsnBp(t *testing.T) {
	resizeCsnBpRequest := &ResizeCsnBpRequest{
		CsnBpId:     util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Bandwidth:   util.PtrInt32(int32(0)),
	}
	err := CSN_CLIENT.ResizeCsnBp(resizeCsnBpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindCsnBp(t *testing.T) {
	unbindCsnBpRequest := &UnbindCsnBpRequest{
		CsnBpId:     util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		CsnId:       util.PtrString(""),
	}
	err := CSN_CLIENT.UnbindCsnBp(unbindCsnBpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateCsn(t *testing.T) {
	updateCsnRequest := &UpdateCsnRequest{
		CsnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := CSN_CLIENT.UpdateCsn(updateCsnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateCsnBp(t *testing.T) {
	updateCsnBpRequest := &UpdateCsnBpRequest{
		CsnBpId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
	}
	err := CSN_CLIENT.UpdateCsnBp(updateCsnBpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRegionBandwidth(t *testing.T) {
	updateRegionBandwidthRequest := &UpdateRegionBandwidthRequest{
		CsnBpId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		LocalRegion: util.PtrString(""),
		PeerRegion:  util.PtrString(""),
		Bandwidth:   util.PtrInt32(int32(0)),
	}
	err := CSN_CLIENT.UpdateRegionBandwidth(updateRegionBandwidthRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateTgw(t *testing.T) {
	updateTgwRequest := &UpdateTgwRequest{
		CsnId:       util.PtrString(""),
		TgwId:       util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := CSN_CLIENT.UpdateTgw(updateTgwRequest)
	ExpectEqual(t.Errorf, nil, err)
}
