package blb

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
	BLB_CLIENT *Client
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

	BLB_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_BillingChangeCancelToPostBlb(t *testing.T) {
	billingChangeCancelToPostBlbRequest := &BillingChangeCancelToPostBlbRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := BLB_CLIENT.BillingChangeCancelToPostBlb(billingChangeCancelToPostBlbRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BillingChangePostToPreBlb(t *testing.T) {
	billingChangePostToPreBlbRequest := &BillingChangePostToPreBlbRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		BillingMethod:     util.PtrString(""),
		PerformanceLevel:  util.PtrString(""),
		ReservationLength: util.PtrInt32(int32(0)),
	}
	result := &BillingChangePostToPreBlbResponse{}
	result, err := BLB_CLIENT.BillingChangePostToPreBlb(billingChangePostToPreBlbRequest)
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
func TestClient_BillingChangePreToPostBlb(t *testing.T) {
	billingChangePreToPostBlbRequest := &BillingChangePreToPostBlbRequest{
		BlbId:                util.PtrString(""),
		ClientToken:          util.PtrString(""),
		BillingMethod:        util.PtrString(""),
		PerformanceLevel:     util.PtrString(""),
		EffectiveImmediately: util.PtrBool(false),
	}
	result := &BillingChangePreToPostBlbResponse{}
	result, err := BLB_CLIENT.BillingChangePreToPostBlb(billingChangePreToPostBlbRequest)
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
func TestClient_BlbInquiry(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	blbInquiryRequest := &BlbInquiryRequest{
		BlbType:          util.PtrString(""),
		PerformanceLevel: util.PtrString(""),
		Count:            util.PtrInt32(int32(0)),
		Billing:          Billing,
	}
	err := BLB_CLIENT.BlbInquiry(blbInquiryRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RefundBlb(t *testing.T) {
	refundBlbRequest := &RefundBlbRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := BLB_CLIENT.RefundBlb(refundBlbRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeBlb(t *testing.T) {
	resizeBlbRequest := &ResizeBlbRequest{
		BlbId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		PerformanceLevel: util.PtrString(""),
	}
	result := &ResizeBlbResponse{}
	result, err := BLB_CLIENT.ResizeBlb(resizeBlbRequest)
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
