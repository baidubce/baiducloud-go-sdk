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

func TestClient_AddAppBlbServerGroupRs(t *testing.T) {
	addAppBlbServerGroupRsRequest := &AddAppBlbServerGroupRsRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		SgId:              util.PtrString(""),
		BackendServerList: []*AppBackendServer{},
	}
	err := BLB_CLIENT.AddAppBlbServerGroupRs(addAppBlbServerGroupRsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddBlbServer(t *testing.T) {
	addBlbServerRequest := &AddBlbServerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		BackendServerList: []*BackendServerModel{},
	}
	err := BLB_CLIENT.AddBlbServer(addBlbServerRequest)
	ExpectEqual(t.Errorf, nil, err)
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
	result := &BlbInquiryResponse{}
	result, err := BLB_CLIENT.BlbInquiry(blbInquiryRequest)
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
func TestClient_CreateAppBlb(t *testing.T) {
	Billing := &BillingForCreate{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &ReservationForCreate{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	createAppBlbRequest := &CreateAppBlbRequest{
		ClientToken:                  util.PtrString(""),
		Name:                         util.PtrString(""),
		BlbType:                      util.PtrString(""),
		Desc:                         util.PtrString(""),
		SubnetId:                     util.PtrString(""),
		VpcId:                        util.PtrString(""),
		Address:                      util.PtrString(""),
		Eip:                          util.PtrString(""),
		Tags:                         []*TagModel{},
		Billing:                      Billing,
		PerformanceLevel:             util.PtrString(""),
		AutoRenewLength:              util.PtrInt32(int32(0)),
		AutoRenewTimeUnit:            util.PtrString(""),
		ResourceGroupId:              util.PtrString(""),
		AllowDelete:                  util.PtrBool(false),
		AllowModify:                  util.PtrBool(false),
		ModificationProtectionReason: util.PtrString(""),
		AllocateIpv6:                 util.PtrBool(false),
	}
	result := &CreateAppBlbResponse{}
	result, err := BLB_CLIENT.CreateAppBlb(createAppBlbRequest)
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
func TestClient_CreateAppBlbHttpListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	createAppBlbHttpListenerRequest := &CreateAppBlbHttpListenerRequest{
		BlbId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		ListenerPort:          util.PtrInt32(int32(0)),
		Scheduler:             util.PtrString(""),
		KeepSession:           util.PtrBool(false),
		KeepSessionType:       util.PtrString(""),
		KeepSessionTimeout:    util.PtrInt32(int32(0)),
		KeepSessionCookieName: util.PtrString(""),
		XForwardedFor:         util.PtrBool(false),
		XForwardedProto:       util.PtrBool(false),
		AdditionalAttributes:  AdditionalAttributes,
		ServerTimeout:         util.PtrInt32(int32(0)),
		RedirectPort:          util.PtrInt32(int32(0)),
		Description:           util.PtrString(""),
	}
	err := BLB_CLIENT.CreateAppBlbHttpListener(createAppBlbHttpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAppBlbHttpsListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	createAppBlbHttpsListenerRequest := &CreateAppBlbHttpsListenerRequest{
		BlbId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		ListenerPort:          util.PtrInt32(int32(0)),
		Scheduler:             util.PtrString(""),
		KeepSession:           util.PtrBool(false),
		KeepSessionType:       util.PtrString(""),
		KeepSessionTimeout:    util.PtrInt32(int32(0)),
		KeepSessionCookieName: util.PtrString(""),
		XForwardedFor:         util.PtrBool(false),
		XForwardedProto:       util.PtrBool(false),
		AdditionalAttributes:  AdditionalAttributes,
		ServerTimeout:         util.PtrInt32(int32(0)),
		CertIds:               []*string{},
		EncryptionType:        util.PtrString(""),
		EncryptionProtocols:   []*string{},
		AppliedCiphers:        util.PtrString(""),
		DualAuth:              util.PtrBool(false),
		ClientCertIds:         []*string{},
		AdditionalCertDomains: []*AdditionalCertDomain{},
		Description:           util.PtrString(""),
	}
	err := BLB_CLIENT.CreateAppBlbHttpsListener(createAppBlbHttpsListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAppBlbPolicy(t *testing.T) {
	createAppBlbPolicyRequest := &CreateAppBlbPolicyRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		AppPolicyVos: []*CreateAppPolicy{},
		BlbType:      util.PtrString(""),
	}
	err := BLB_CLIENT.CreateAppBlbPolicy(createAppBlbPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAppBlbServerGroup(t *testing.T) {
	createAppBlbServerGroupRequest := &CreateAppBlbServerGroupRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		Name:              util.PtrString(""),
		Desc:              util.PtrString(""),
		BackendServerList: []*AppBackendServer{},
	}
	result := &CreateAppBlbServerGroupResponse{}
	result, err := BLB_CLIENT.CreateAppBlbServerGroup(createAppBlbServerGroupRequest)
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
func TestClient_CreateAppBlbServerGroupPort(t *testing.T) {
	createAppBlbServerGroupPortRequest := &CreateAppBlbServerGroupPortRequest{
		BlbId:                       util.PtrString(""),
		ClientToken:                 util.PtrString(""),
		SgId:                        util.PtrString(""),
		Port:                        util.PtrInt32(int32(0)),
		BlbType:                     util.PtrString(""),
		EnableHealthCheck:           util.PtrBool(false),
		HealthCheck:                 util.PtrString(""),
		HealthCheckPort:             util.PtrInt32(int32(0)),
		HealthCheckUrlPath:          util.PtrString(""),
		HealthCheckTimeoutInSecond:  util.PtrInt32(int32(0)),
		HealthCheckIntervalInSecond: util.PtrInt32(int32(0)),
		HealthCheckDownRetry:        util.PtrInt32(int32(0)),
		HealthCheckUpRetry:          util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:     util.PtrString(""),
		HealthCheckHost:             util.PtrString(""),
		UdpHealthCheckString:        util.PtrString(""),
	}
	result := &CreateAppBlbServerGroupPortResponse{}
	result, err := BLB_CLIENT.CreateAppBlbServerGroupPort(createAppBlbServerGroupPortRequest)
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
func TestClient_CreateAppBlbSslListener(t *testing.T) {
	createAppBlbSslListenerRequest := &CreateAppBlbSslListenerRequest{
		BlbId:               util.PtrString(""),
		ClientToken:         util.PtrString(""),
		ListenerPort:        util.PtrInt32(int32(0)),
		Scheduler:           util.PtrString(""),
		CertIds:             []*string{},
		EncryptionType:      util.PtrString(""),
		EncryptionProtocols: []*string{},
		AppliedCiphers:      util.PtrString(""),
		DualAuth:            util.PtrBool(false),
		ClientCertIds:       []*string{},
		Description:         util.PtrString(""),
	}
	err := BLB_CLIENT.CreateAppBlbSslListener(createAppBlbSslListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAppBlbTcpListener(t *testing.T) {
	createAppBlbTcpListenerRequest := &CreateAppBlbTcpListenerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		ListenerPort:      util.PtrInt32(int32(0)),
		Scheduler:         util.PtrString(""),
		TcpSessionTimeout: util.PtrInt32(int32(0)),
		Description:       util.PtrString(""),
	}
	err := BLB_CLIENT.CreateAppBlbTcpListener(createAppBlbTcpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAppBlbUdpListener(t *testing.T) {
	createAppBlbUdpListenerRequest := &CreateAppBlbUdpListenerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		ListenerPort:      util.PtrInt32(int32(0)),
		Scheduler:         util.PtrString(""),
		UdpSessionTimeout: util.PtrInt32(int32(0)),
		Description:       util.PtrString(""),
	}
	err := BLB_CLIENT.CreateAppBlbUdpListener(createAppBlbUdpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateBlb(t *testing.T) {
	Billing := &BillingForCreate{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &ReservationForCreate{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	createBlbRequest := &CreateBlbRequest{
		ClientToken:                  util.PtrString(""),
		Name:                         util.PtrString(""),
		Desc:                         util.PtrString(""),
		VpcId:                        util.PtrString(""),
		SubnetId:                     util.PtrString(""),
		Address:                      util.PtrString(""),
		BlbType:                      util.PtrString(""),
		Eip:                          util.PtrString(""),
		Tags:                         []*TagModel{},
		Billing:                      Billing,
		PerformanceLevel:             util.PtrString(""),
		AutoRenewLength:              util.PtrInt32(int32(0)),
		AutoRenewTimeUnit:            util.PtrString(""),
		ResourceGroupId:              util.PtrString(""),
		AllowDelete:                  util.PtrBool(false),
		AllowModify:                  util.PtrBool(false),
		ModificationProtectionReason: util.PtrString(""),
		AllocateIpv6:                 util.PtrBool(false),
	}
	result := &CreateBlbResponse{}
	result, err := BLB_CLIENT.CreateBlb(createBlbRequest)
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
func TestClient_CreateBlbHttpListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	createBlbHttpListenerRequest := &CreateBlbHttpListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		KeepSession:                util.PtrBool(false),
		KeepSessionType:            util.PtrString(""),
		KeepSessionDuration:        util.PtrInt32(int32(0)),
		KeepSessionCookieName:      util.PtrString(""),
		XForwardFor:                util.PtrBool(false),
		XForwardedProto:            util.PtrBool(false),
		AdditionalAttributes:       AdditionalAttributes,
		HealthCheckType:            util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckURI:             util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:    util.PtrString(""),
		HealthCheckHost:            util.PtrString(""),
		ServerTimeout:              util.PtrInt32(int32(0)),
		RedirectPort:               util.PtrInt32(int32(0)),
	}
	err := BLB_CLIENT.CreateBlbHttpListener(createBlbHttpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateBlbHttpsListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	createBlbHttpsListenerRequest := &CreateBlbHttpsListenerRequest{
		BlbId:                      util.PtrString(""),
		ClientToken:                util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		KeepSession:                util.PtrBool(false),
		KeepSessionType:            util.PtrString(""),
		KeepSessionDuration:        util.PtrInt32(int32(0)),
		KeepSessionCookieName:      util.PtrInt32(int32(0)),
		XForwardFor:                util.PtrBool(false),
		XForwardedProto:            util.PtrBool(false),
		AdditionalAttributes:       AdditionalAttributes,
		HealthCheckType:            util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckURI:             util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:    util.PtrString(""),
		HealthCheckHost:            util.PtrString(""),
		ServerTimeout:              util.PtrInt32(int32(0)),
		CertIds:                    []*string{},
		EncryptionType:             util.PtrString(""),
		EncryptionProtocols:        []*string{},
		AppliedCiphers:             util.PtrString(""),
		DualAuth:                   util.PtrBool(false),
		ClientCertIds:              []*string{},
		AdditionalCertDomains:      []*AdditionalCertDomain{},
	}
	err := BLB_CLIENT.CreateBlbHttpsListener(createBlbHttpsListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateBlbSslListener(t *testing.T) {
	createBlbSslListenerRequest := &CreateBlbSslListenerRequest{
		BlbId:                      util.PtrString(""),
		ClientToken:                util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		HealthCheckType:            util.PtrString(""),
		ServerTimeout:              util.PtrInt32(int32(0)),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		CertIds:                    []*string{},
		EncryptionType:             util.PtrString(""),
		EncryptionProtocols:        []*string{},
		AppliedCiphers:             util.PtrString(""),
		DualAuth:                   util.PtrBool(false),
		ClientCertIds:              []*string{},
	}
	err := BLB_CLIENT.CreateBlbSslListener(createBlbSslListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateBlbTcpListener(t *testing.T) {
	createBlbTcpListenerRequest := &CreateBlbTcpListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		TcpSessionTimeout:          util.PtrInt32(int32(0)),
		HealthCheckType:            util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
	}
	err := BLB_CLIENT.CreateBlbTcpListener(createBlbTcpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateBlbUdpListener(t *testing.T) {
	createBlbUdpListenerRequest := &CreateBlbUdpListenerRequest{
		BlbId:                      util.PtrString(""),
		ClientToken:                util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		HealthCheckType:            util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckString:          util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		UdpSessionTimeout:          util.PtrInt32(int32(0)),
	}
	err := BLB_CLIENT.CreateBlbUdpListener(createBlbUdpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAppBlbListener(t *testing.T) {
	deleteAppBlbListenerRequest := &DeleteAppBlbListenerRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		PortList:     []*int32{},
		PortTypeList: []*PortTypeModel{},
	}
	err := BLB_CLIENT.DeleteAppBlbListener(deleteAppBlbListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAppBlbPolicy(t *testing.T) {
	deleteAppBlbPolicyRequest := &DeleteAppBlbPolicyRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Port:         util.PtrInt32(int32(0)),
		PolicyIdList: []*string{},
		BlbType:      util.PtrString(""),
	}
	err := BLB_CLIENT.DeleteAppBlbPolicy(deleteAppBlbPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAppBlbServerGroup(t *testing.T) {
	deleteAppBlbServerGroupRequest := &DeleteAppBlbServerGroupRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		SgId:        util.PtrString(""),
	}
	err := BLB_CLIENT.DeleteAppBlbServerGroup(deleteAppBlbServerGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAppBlbServerGroupPort(t *testing.T) {
	deleteAppBlbServerGroupPortRequest := &DeleteAppBlbServerGroupPortRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		SgId:        util.PtrString(""),
		PortIdList:  []*string{},
	}
	err := BLB_CLIENT.DeleteAppBlbServerGroupPort(deleteAppBlbServerGroupPortRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAppBlbServerGroupRs(t *testing.T) {
	deleteAppBlbServerGroupRsRequest := &DeleteAppBlbServerGroupRsRequest{
		BlbId:               util.PtrString(""),
		ClientToken:         util.PtrString(""),
		SgId:                util.PtrString(""),
		BackendServerIdList: []*string{},
	}
	err := BLB_CLIENT.DeleteAppBlbServerGroupRs(deleteAppBlbServerGroupRsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteBlbListener(t *testing.T) {
	deleteBlbListenerRequest := &DeleteBlbListenerRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		PortList:     []*int32{},
		PortTypeList: []*PortTypeModel{},
	}
	err := BLB_CLIENT.DeleteBlbListener(deleteBlbListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteBlbServer(t *testing.T) {
	deleteBlbServerRequest := &DeleteBlbServerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		BackendServerList: []*string{},
	}
	err := BLB_CLIENT.DeleteBlbServer(deleteBlbServerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeAppBlb(t *testing.T) {
	describeAppBlbRequest := &DescribeAppBlbRequest{
		BlbId: util.PtrString(""),
	}
	result := &DescribeAppBlbResponse{}
	result, err := BLB_CLIENT.DescribeAppBlb(describeAppBlbRequest)
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
func TestClient_DescribeAppBlbHttpListener(t *testing.T) {
	describeAppBlbHttpListenerRequest := &DescribeAppBlbHttpListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbHttpListenerResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbHttpListener(describeAppBlbHttpListenerRequest)
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
func TestClient_DescribeAppBlbHttpsListener(t *testing.T) {
	describeAppBlbHttpsListenerRequest := &DescribeAppBlbHttpsListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbHttpsListenerResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbHttpsListener(describeAppBlbHttpsListenerRequest)
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
func TestClient_DescribeAppBlbListener(t *testing.T) {
	describeAppBlbListenerRequest := &DescribeAppBlbListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbListenerResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbListener(describeAppBlbListenerRequest)
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
func TestClient_DescribeAppBlbPolicy(t *testing.T) {
	describeAppBlbPolicyRequest := &DescribeAppBlbPolicyRequest{
		BlbId:   util.PtrString(""),
		Port:    util.PtrInt32(int32(0)),
		Type:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbPolicyResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbPolicy(describeAppBlbPolicyRequest)
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
func TestClient_DescribeAppBlbServerGroup(t *testing.T) {
	describeAppBlbServerGroupRequest := &DescribeAppBlbServerGroupRequest{
		BlbId:        util.PtrString(""),
		Name:         util.PtrString(""),
		ExactlyMatch: util.PtrBool(false),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbServerGroupResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbServerGroup(describeAppBlbServerGroupRequest)
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
func TestClient_DescribeAppBlbServerGroupMountRs(t *testing.T) {
	describeAppBlbServerGroupMountRsRequest := &DescribeAppBlbServerGroupMountRsRequest{
		BlbId: util.PtrString(""),
		SgId:  util.PtrString(""),
	}
	result := &DescribeAppBlbServerGroupMountRsResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbServerGroupMountRs(describeAppBlbServerGroupMountRsRequest)
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
func TestClient_DescribeAppBlbServerGroupRs(t *testing.T) {
	describeAppBlbServerGroupRsRequest := &DescribeAppBlbServerGroupRsRequest{
		BlbId:   util.PtrString(""),
		SgId:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbServerGroupRsResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbServerGroupRs(describeAppBlbServerGroupRsRequest)
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
func TestClient_DescribeAppBlbServerGroupUnmountRs(t *testing.T) {
	describeAppBlbServerGroupUnmountRsRequest := &DescribeAppBlbServerGroupUnmountRsRequest{
		BlbId: util.PtrString(""),
		SgId:  util.PtrString(""),
	}
	result := &DescribeAppBlbServerGroupUnmountRsResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbServerGroupUnmountRs(describeAppBlbServerGroupUnmountRsRequest)
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
func TestClient_DescribeAppBlbSslListener(t *testing.T) {
	describeAppBlbSslListenerRequest := &DescribeAppBlbSslListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbSslListenerResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbSslListener(describeAppBlbSslListenerRequest)
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
func TestClient_DescribeAppBlbTcpListener(t *testing.T) {
	describeAppBlbTcpListenerRequest := &DescribeAppBlbTcpListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbTcpListenerResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbTcpListener(describeAppBlbTcpListenerRequest)
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
func TestClient_DescribeAppBlbUdpListener(t *testing.T) {
	describeAppBlbUdpListenerRequest := &DescribeAppBlbUdpListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbUdpListenerResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbUdpListener(describeAppBlbUdpListenerRequest)
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
func TestClient_DescribeAppBlbs(t *testing.T) {
	describeAppBlbsRequest := &DescribeAppBlbsRequest{
		Address:      util.PtrString(""),
		Name:         util.PtrString(""),
		BlbId:        util.PtrString(""),
		BccId:        util.PtrString(""),
		ExactlyMatch: util.PtrBool(false),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeAppBlbsResponse{}
	result, err := BLB_CLIENT.DescribeAppBlbs(describeAppBlbsRequest)
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
func TestClient_DescribeBlb(t *testing.T) {
	describeBlbRequest := &DescribeBlbRequest{
		BlbId: util.PtrString(""),
		Type:  util.PtrString(""),
	}
	result := &DescribeBlbResponse{}
	result, err := BLB_CLIENT.DescribeBlb(describeBlbRequest)
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
func TestClient_DescribeBlbHttpListener(t *testing.T) {
	describeBlbHttpListenerRequest := &DescribeBlbHttpListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbHttpListenerResponse{}
	result, err := BLB_CLIENT.DescribeBlbHttpListener(describeBlbHttpListenerRequest)
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
func TestClient_DescribeBlbHttpsListener(t *testing.T) {
	describeBlbHttpsListenerRequest := &DescribeBlbHttpsListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbHttpsListenerResponse{}
	result, err := BLB_CLIENT.DescribeBlbHttpsListener(describeBlbHttpsListenerRequest)
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
func TestClient_DescribeBlbListener(t *testing.T) {
	describeBlbListenerRequest := &DescribeBlbListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbListenerResponse{}
	result, err := BLB_CLIENT.DescribeBlbListener(describeBlbListenerRequest)
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
func TestClient_DescribeBlbServerHealth(t *testing.T) {
	describeBlbServerHealthRequest := &DescribeBlbServerHealthRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbServerHealthResponse{}
	result, err := BLB_CLIENT.DescribeBlbServerHealth(describeBlbServerHealthRequest)
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
func TestClient_DescribeBlbServers(t *testing.T) {
	describeBlbServersRequest := &DescribeBlbServersRequest{
		BlbId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbServersResponse{}
	result, err := BLB_CLIENT.DescribeBlbServers(describeBlbServersRequest)
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
func TestClient_DescribeBlbSslListener(t *testing.T) {
	describeBlbSslListenerRequest := &DescribeBlbSslListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbSslListenerResponse{}
	result, err := BLB_CLIENT.DescribeBlbSslListener(describeBlbSslListenerRequest)
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
func TestClient_DescribeBlbTcpListener(t *testing.T) {
	describeBlbTcpListenerRequest := &DescribeBlbTcpListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbTcpListenerResponse{}
	result, err := BLB_CLIENT.DescribeBlbTcpListener(describeBlbTcpListenerRequest)
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
func TestClient_DescribeBlbUdpListener(t *testing.T) {
	describeBlbUdpListenerRequest := &DescribeBlbUdpListenerRequest{
		BlbId:        util.PtrString(""),
		ListenerPort: util.PtrInt32(int32(0)),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &DescribeBlbUdpListenerResponse{}
	result, err := BLB_CLIENT.DescribeBlbUdpListener(describeBlbUdpListenerRequest)
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
func TestClient_DescribeBlbs(t *testing.T) {
	describeBlbsRequest := &DescribeBlbsRequest{
		Address:      util.PtrString(""),
		Name:         util.PtrString(""),
		BlbId:        util.PtrString(""),
		BccId:        util.PtrString(""),
		ExactlyMatch: util.PtrBool(false),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
		Type:         util.PtrString(""),
	}
	result := &DescribeBlbsResponse{}
	result, err := BLB_CLIENT.DescribeBlbs(describeBlbsRequest)
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
func TestClient_RefundBlb(t *testing.T) {
	refundBlbRequest := &RefundBlbRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := BLB_CLIENT.RefundBlb(refundBlbRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseAppBlb(t *testing.T) {
	releaseAppBlbRequest := &ReleaseAppBlbRequest{
		BlbId: util.PtrString(""),
	}
	err := BLB_CLIENT.ReleaseAppBlb(releaseAppBlbRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseBlb(t *testing.T) {
	releaseBlbRequest := &ReleaseBlbRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := BLB_CLIENT.ReleaseBlb(releaseBlbRequest)
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
func TestClient_UpdateAppBlb(t *testing.T) {
	updateAppBlbRequest := &UpdateAppBlbRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Name:         util.PtrString(""),
		Desc:         util.PtrString(""),
		AllowDelete:  util.PtrBool(false),
		AllocateIpv6: util.PtrBool(false),
	}
	err := BLB_CLIENT.UpdateAppBlb(updateAppBlbRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbHttpListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	updateAppBlbHttpListenerRequest := &UpdateAppBlbHttpListenerRequest{
		BlbId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		ListenerPort:          util.PtrInt32(int32(0)),
		Scheduler:             util.PtrString(""),
		KeepSession:           util.PtrBool(false),
		KeepSessionType:       util.PtrString(""),
		KeepSessionTimeout:    util.PtrInt32(int32(0)),
		KeepSessionCookieName: util.PtrString(""),
		XForwardedFor:         util.PtrBool(false),
		XForwardedProto:       util.PtrBool(false),
		AdditionalAttributes:  AdditionalAttributes,
		ServerTimeout:         util.PtrInt32(int32(0)),
		RedirectPort:          util.PtrInt32(int32(0)),
		Description:           util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbHttpListener(updateAppBlbHttpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbHttpsListener(t *testing.T) {
	updateAppBlbHttpsListenerRequest := &UpdateAppBlbHttpsListenerRequest{
		BlbId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		ListenerPort:          util.PtrInt32(int32(0)),
		Scheduler:             util.PtrString(""),
		KeepSession:           util.PtrBool(false),
		KeepSessionType:       util.PtrString(""),
		KeepSessionTimeout:    util.PtrInt32(int32(0)),
		KeepSessionCookieName: util.PtrString(""),
		XForwardedFor:         util.PtrBool(false),
		XForwardedProto:       util.PtrBool(false),
		ServerTimeout:         util.PtrInt32(int32(0)),
		CertIds:               []*string{},
		EncryptionType:        util.PtrString(""),
		EncryptionProtocols:   []*string{},
		AppliedCiphers:        util.PtrString(""),
		DualAuth:              util.PtrBool(false),
		ClientCertIds:         []*string{},
		AdditionalCertDomains: []*AdditionalCertDomain{},
		Description:           util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbHttpsListener(updateAppBlbHttpsListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbPolicy(t *testing.T) {
	updateAppBlbPolicyRequest := &UpdateAppBlbPolicyRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Port:        util.PtrInt32(int32(0)),
		BlbType:     util.PtrString(""),
		PolicyList:  []*AppPolicyForUpdate{},
	}
	err := BLB_CLIENT.UpdateAppBlbPolicy(updateAppBlbPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbServerGroup(t *testing.T) {
	updateAppBlbServerGroupRequest := &UpdateAppBlbServerGroupRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		SgId:        util.PtrString(""),
		Name:        util.PtrString(""),
		Desc:        util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbServerGroup(updateAppBlbServerGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbServerGroupPort(t *testing.T) {
	updateAppBlbServerGroupPortRequest := &UpdateAppBlbServerGroupPortRequest{
		BlbId:                       util.PtrString(""),
		ClientToken:                 util.PtrString(""),
		SgId:                        util.PtrString(""),
		PortId:                      util.PtrString(""),
		EnableHealthCheck:           util.PtrBool(false),
		HealthCheck:                 util.PtrString(""),
		HealthCheckPort:             util.PtrInt32(int32(0)),
		HealthCheckUrlPath:          util.PtrString(""),
		HealthCheckTimeoutInSecond:  util.PtrInt32(int32(0)),
		HealthCheckIntervalInSecond: util.PtrInt32(int32(0)),
		HealthCheckDownRetry:        util.PtrInt32(int32(0)),
		HealthCheckUpRetry:          util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:     util.PtrString(""),
		HealthCheckHost:             util.PtrString(""),
		UdpHealthCheckString:        util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbServerGroupPort(updateAppBlbServerGroupPortRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbServerGroupRs(t *testing.T) {
	updateAppBlbServerGroupRsRequest := &UpdateAppBlbServerGroupRsRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		SgId:              util.PtrString(""),
		BackendServerList: []*AppBackendServer{},
	}
	err := BLB_CLIENT.UpdateAppBlbServerGroupRs(updateAppBlbServerGroupRsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbSslListener(t *testing.T) {
	updateAppBlbSslListenerRequest := &UpdateAppBlbSslListenerRequest{
		BlbId:               util.PtrString(""),
		ClientToken:         util.PtrString(""),
		ListenerPort:        util.PtrInt32(int32(0)),
		Scheduler:           util.PtrString(""),
		CertIds:             []*string{},
		EncryptionType:      util.PtrString(""),
		EncryptionProtocols: []*string{},
		AppliedCiphers:      util.PtrString(""),
		DualAuth:            util.PtrBool(false),
		ClientCertIds:       []*string{},
		Description:         util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbSslListener(updateAppBlbSslListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbTcpListener(t *testing.T) {
	updateAppBlbTcpListenerRequest := &UpdateAppBlbTcpListenerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		ListenerPort:      util.PtrInt32(int32(0)),
		Scheduler:         util.PtrString(""),
		TcpSessionTimeout: util.PtrInt32(int32(0)),
		Description:       util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbTcpListener(updateAppBlbTcpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAppBlbUdpListener(t *testing.T) {
	updateAppBlbUdpListenerRequest := &UpdateAppBlbUdpListenerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		ListenerPort:      util.PtrInt32(int32(0)),
		Scheduler:         util.PtrString(""),
		UdpSessionTimeout: util.PtrInt32(int32(0)),
		Description:       util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateAppBlbUdpListener(updateAppBlbUdpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlb(t *testing.T) {
	updateBlbRequest := &UpdateBlbRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Name:         util.PtrString(""),
		Desc:         util.PtrString(""),
		AllowDelete:  util.PtrBool(false),
		AllocateIpv6: util.PtrBool(false),
	}
	err := BLB_CLIENT.UpdateBlb(updateBlbRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbAcl(t *testing.T) {
	updateBlbAclRequest := &UpdateBlbAclRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		SupportAcl:  util.PtrBool(false),
	}
	err := BLB_CLIENT.UpdateBlbAcl(updateBlbAclRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbHttpListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	updateBlbHttpListenerRequest := &UpdateBlbHttpListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		KeepSession:                util.PtrBool(false),
		KeepSessionType:            util.PtrString(""),
		KeepSessionDuration:        util.PtrInt32(int32(0)),
		KeepSessionCookieName:      util.PtrString(""),
		XForwardFor:                util.PtrBool(false),
		XForwardedProto:            util.PtrBool(false),
		AdditionalAttributes:       AdditionalAttributes,
		HealthCheckType:            util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckURI:             util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:    util.PtrString(""),
		HealthCheckHost:            util.PtrString(""),
		ServerTimeout:              util.PtrInt32(int32(0)),
		RedirectPort:               util.PtrInt32(int32(0)),
	}
	err := BLB_CLIENT.UpdateBlbHttpListener(updateBlbHttpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbHttpsListener(t *testing.T) {
	AdditionalAttributes := &AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	updateBlbHttpsListenerRequest := &UpdateBlbHttpsListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		KeepSession:                util.PtrBool(false),
		KeepSessionType:            util.PtrString(""),
		KeepSessionDuration:        util.PtrInt32(int32(0)),
		KeepSessionCookieName:      util.PtrString(""),
		XForwardFor:                util.PtrBool(false),
		XForwardedProto:            util.PtrBool(false),
		AdditionalAttributes:       AdditionalAttributes,
		HealthCheckType:            util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckURI:             util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:    util.PtrString(""),
		HealthCheckHost:            util.PtrString(""),
		ServerTimeout:              util.PtrInt32(int32(0)),
		CertIds:                    []*string{},
		AdditionalCertDomains:      []*AdditionalCertDomain{},
		EncryptionType:             util.PtrString(""),
		EncryptionProtocols:        []*string{},
		AppliedCiphers:             util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateBlbHttpsListener(updateBlbHttpsListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbModifyProtection(t *testing.T) {
	updateBlbModifyProtectionRequest := &UpdateBlbModifyProtectionRequest{
		BlbId:                        util.PtrString(""),
		ClientToken:                  util.PtrString(""),
		AllowModify:                  util.PtrBool(false),
		ModificationProtectionReason: util.PtrString(""),
	}
	err := BLB_CLIENT.UpdateBlbModifyProtection(updateBlbModifyProtectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbServer(t *testing.T) {
	updateBlbServerRequest := &UpdateBlbServerRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		BackendServerList: []*BackendServerModel{},
	}
	err := BLB_CLIENT.UpdateBlbServer(updateBlbServerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbSslListener(t *testing.T) {
	updateBlbSslListenerRequest := &UpdateBlbSslListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		HealthCheckType:            util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		CertIds:                    []*string{},
		EncryptionType:             util.PtrString(""),
		EncryptionProtocols:        []*string{},
		AppliedCiphers:             util.PtrString(""),
		DualAuth:                   util.PtrBool(false),
		ClientCertIds:              []*string{},
	}
	err := BLB_CLIENT.UpdateBlbSslListener(updateBlbSslListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbTcpListener(t *testing.T) {
	updateBlbTcpListenerRequest := &UpdateBlbTcpListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		HealthCheckType:            util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
	}
	err := BLB_CLIENT.UpdateBlbTcpListener(updateBlbTcpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateBlbUdpListener(t *testing.T) {
	updateBlbUdpListenerRequest := &UpdateBlbUdpListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		HealthCheckType:            util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		HealthCheckString:          util.PtrString(""),
		UdpSessionTimeout:          util.PtrInt32(int32(0)),
	}
	err := BLB_CLIENT.UpdateBlbUdpListener(updateBlbUdpListenerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
