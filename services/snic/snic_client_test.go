package snic

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
	SNIC_CLIENT *Client
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

	SNIC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateSnic(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createSnicRequest := &CreateSnicRequest{
		ClientToken:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		Name:            util.PtrString(""),
		SubnetId:        util.PtrString(""),
		Service:         util.PtrString(""),
		Description:     util.PtrString(""),
		IpAddress:       util.PtrString(""),
		Bandwidth:       util.PtrInt32(int32(0)),
		Billing:         Billing,
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
	}
	result := &CreateSnicResponse{}
	result, err := SNIC_CLIENT.CreateSnic(createSnicRequest)
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
func TestClient_DeleteSnic(t *testing.T) {
	deleteSnicRequest := &DeleteSnicRequest{
		EndpointId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := SNIC_CLIENT.DeleteSnic(deleteSnicRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeSnic(t *testing.T) {
	describeSnicRequest := &DescribeSnicRequest{
		EndpointId: util.PtrString(""),
	}
	result := &DescribeSnicResponse{}
	result, err := SNIC_CLIENT.DescribeSnic(describeSnicRequest)
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
func TestClient_ListSnic(t *testing.T) {
	listSnicRequest := &ListSnicRequest{
		VpcId:     util.PtrString(""),
		Name:      util.PtrString(""),
		IpAddress: util.PtrString(""),
		Status:    util.PtrString(""),
		SubnetId:  util.PtrString(""),
		Service:   util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &ListSnicResponse{}
	result, err := SNIC_CLIENT.ListSnic(listSnicRequest)
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
func TestClient_QueryAvailablePublicServices(t *testing.T) {
	result := &QueryAvailablePublicServicesResponse{}
	result, err := SNIC_CLIENT.QueryAvailablePublicServices()
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
func TestClient_UpdateSnic(t *testing.T) {
	updateSnicRequest := &UpdateSnicRequest{
		EndpointId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := SNIC_CLIENT.UpdateSnic(updateSnicRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSnicEsg(t *testing.T) {
	updateSnicEsgRequest := &UpdateSnicEsgRequest{
		EndpointId:                 util.PtrString(""),
		ClientToken:                util.PtrString(""),
		EnterpriseSecurityGroupIds: []*string{},
	}
	err := SNIC_CLIENT.UpdateSnicEsg(updateSnicEsgRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSnicSg(t *testing.T) {
	updateSnicSgRequest := &UpdateSnicSgRequest{
		EndpointId:       util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SecurityGroupIds: []*string{},
	}
	err := SNIC_CLIENT.UpdateSnicSg(updateSnicSgRequest)
	ExpectEqual(t.Errorf, nil, err)
}
