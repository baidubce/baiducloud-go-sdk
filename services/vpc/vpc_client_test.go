package vpc

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
	VPC_CLIENT *Client
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

	VPC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AcceptPeerConn(t *testing.T) {
	acceptPeerConnRequest := &AcceptPeerConnRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.AcceptPeerConn(acceptPeerConnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ActiveStandbySwitchover(t *testing.T) {
	activeStandbySwitchoverRequest := &ActiveStandbySwitchoverRequest{
		RouteRuleId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ActiveStandbySwitchover(activeStandbySwitchoverRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddAclRule(t *testing.T) {
	addAclRuleRequest := &AddAclRuleRequest{
		ClientToken: util.PtrString(""),
		AclRules:    []*AclRuleRequest{},
	}
	err := VPC_CLIENT.AddAclRule(addAclRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddEniIp(t *testing.T) {
	addEniIpRequest := &AddEniIpRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		IsIpv6:           util.PtrBool(false),
		PrivateIpAddress: util.PtrString(""),
	}
	result := &AddEniIpResponse{}
	result, err := VPC_CLIENT.AddEniIp(addEniIpRequest)
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
func TestClient_AddIpAddressToIpGroup(t *testing.T) {
	addIpAddressToIpGroupRequest := &AddIpAddressToIpGroupRequest{
		IpSetId:       util.PtrString(""),
		ClientToken:   util.PtrString(""),
		IpAddressInfo: []*TemplateIpAddressInfo{},
	}
	err := VPC_CLIENT.AddIpAddressToIpGroup(addIpAddressToIpGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddIpGroupToIpSet(t *testing.T) {
	addIpGroupToIpSetRequest := &AddIpGroupToIpSetRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpSetIds:    []*string{},
	}
	err := VPC_CLIENT.AddIpGroupToIpSet(addIpGroupToIpSetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AttachEniInstance(t *testing.T) {
	attachEniInstanceRequest := &AttachEniInstanceRequest{
		EniId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceId:   util.PtrString(""),
		InstanceType: util.PtrString(""),
	}
	err := VPC_CLIENT.AttachEniInstance(attachEniInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AuthorizeEnterpriseSecurityGroupRules(t *testing.T) {
	authorizeEnterpriseSecurityGroupRulesRequest := &AuthorizeEnterpriseSecurityGroupRulesRequest{
		EnterpriseSecurityGroupId: util.PtrString(""),
		ClientToken:               util.PtrString(""),
		Rules:                     []*EnterpriseSecurityGroupRuleModel{},
	}
	err := VPC_CLIENT.AuthorizeEnterpriseSecurityGroupRules(authorizeEnterpriseSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AuthorizeSecurityGroupRules(t *testing.T) {
	Rule := &SecurityGroupRuleModel{
		Remark:              util.PtrString(""),
		Direction:           util.PtrString(""),
		Ethertype:           util.PtrString(""),
		PortRange:           util.PtrString(""),
		Protocol:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		SourceIp:            util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		DestIp:              util.PtrString(""),
		SecurityGroupId:     util.PtrString(""),
		SecurityGroupRuleId: util.PtrString(""),
		CreatedTime:         util.PtrString(""),
		UpdatedTime:         util.PtrString(""),
	}
	authorizeSecurityGroupRulesRequest := &AuthorizeSecurityGroupRulesRequest{
		SecurityGroupId: util.PtrString(""),
		SgVersion:       util.PtrInt64(int64(0)),
		ClientToken:     util.PtrString(""),
		Rule:            Rule,
	}
	err := VPC_CLIENT.AuthorizeSecurityGroupRules(authorizeSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchAddDnatRules(t *testing.T) {
	batchAddDnatRulesRequest := &BatchAddDnatRulesRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Rules:       []*DnatRuleRequest{},
	}
	result := &BatchAddDnatRulesResponse{}
	result, err := VPC_CLIENT.BatchAddDnatRules(batchAddDnatRulesRequest)
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
func TestClient_BatchAddEniIp(t *testing.T) {
	batchAddEniIpRequest := &BatchAddEniIpRequest{
		EniId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		IsIpv6:                util.PtrBool(false),
		PrivateIpAddresses:    []*string{},
		PrivateIpAddressCount: util.PtrInt32(int32(0)),
	}
	result := &BatchAddEniIpResponse{}
	result, err := VPC_CLIENT.BatchAddEniIp(batchAddEniIpRequest)
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
func TestClient_BatchAddSnatRules(t *testing.T) {
	batchAddSnatRulesRequest := &BatchAddSnatRulesRequest{
		ClientToken: util.PtrString(""),
		NatId:       util.PtrString(""),
		SnatRules:   []*SnatRuleRequest{},
	}
	result := &BatchAddSnatRulesResponse{}
	result, err := VPC_CLIENT.BatchAddSnatRules(batchAddSnatRulesRequest)
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
func TestClient_BatchCreateSslVpnUsers(t *testing.T) {
	batchCreateSslVpnUsersRequest := &BatchCreateSslVpnUsersRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		SslVpnUsers: []*SslVpnUser{},
	}
	result := &BatchCreateSslVpnUsersResponse{}
	result, err := VPC_CLIENT.BatchCreateSslVpnUsers(batchCreateSslVpnUsersRequest)
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
func TestClient_BatchDeleteEniIp(t *testing.T) {
	batchDeleteEniIpRequest := &BatchDeleteEniIpRequest{
		EniId:              util.PtrString(""),
		ClientToken:        util.PtrString(""),
		PrivateIpAddresses: []*string{},
	}
	err := VPC_CLIENT.BatchDeleteEniIp(batchDeleteEniIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindEip(t *testing.T) {
	bindEipRequest := &BindEipRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Eip:         util.PtrString(""),
	}
	err := VPC_CLIENT.BindEip(bindEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindEniEip(t *testing.T) {
	bindEniEipRequest := &BindEniEipRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		PublicIpAddress:  util.PtrString(""),
	}
	err := VPC_CLIENT.BindEniEip(bindEniEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindHaVipEip(t *testing.T) {
	bindHaVipEipRequest := &BindHaVipEipRequest{
		HaVipId:         util.PtrString(""),
		ClientToken:     util.PtrString(""),
		PublicIpAddress: util.PtrString(""),
	}
	err := VPC_CLIENT.BindHaVipEip(bindHaVipEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindHaVipInstance(t *testing.T) {
	bindHaVipInstanceRequest := &BindHaVipInstanceRequest{
		HaVipId:      util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceIds:  []*string{},
		InstanceType: util.PtrString(""),
	}
	err := VPC_CLIENT.BindHaVipInstance(bindHaVipInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindPhysicalDedicatedLine(t *testing.T) {
	bindPhysicalDedicatedLineRequest := &BindPhysicalDedicatedLineRequest{
		EtGatewayId: util.PtrString(""),
		ClientToken: util.PtrString(""),
		EtId:        util.PtrString(""),
		ChannelId:   util.PtrString(""),
		LocalCidrs:  []*string{},
	}
	err := VPC_CLIENT.BindPhysicalDedicatedLine(bindPhysicalDedicatedLineRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ClosePeerConnSyncDns(t *testing.T) {
	closePeerConnSyncDnsRequest := &ClosePeerConnSyncDnsRequest{
		PeerConnId:  util.PtrString(""),
		Role:        util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ClosePeerConnSyncDns(closePeerConnSyncDnsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CloseVpcRelay(t *testing.T) {
	closeVpcRelayRequest := &CloseVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.CloseVpcRelay(closeVpcRelayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateDedicatedGateway(t *testing.T) {
	createDedicatedGatewayRequest := &CreateDedicatedGatewayRequest{
		ClientToken:     util.PtrString(""),
		Name:            util.PtrString(""),
		VpcId:           util.PtrString(""),
		Speed:           util.PtrInt32(int32(0)),
		Description:     util.PtrString(""),
		EtId:            util.PtrString(""),
		ChannelId:       util.PtrString(""),
		LocalCidrs:      []*string{},
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
	}
	result := &CreateDedicatedGatewayResponse{}
	result, err := VPC_CLIENT.CreateDedicatedGateway(createDedicatedGatewayRequest)
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
func TestClient_CreateDedicatedGatewayHealthCheck(t *testing.T) {
	createDedicatedGatewayHealthCheckRequest := &CreateDedicatedGatewayHealthCheckRequest{
		EtGatewayId:           util.PtrString(""),
		ClientToken:           util.PtrString(""),
		HealthCheckSourceIp:   util.PtrString(""),
		HealthCheckType:       util.PtrString(""),
		HealthCheckPort:       util.PtrInt32(int32(0)),
		HealthCheckInterval:   util.PtrInt32(int32(0)),
		HealthThreshold:       util.PtrInt32(int32(0)),
		UnhealthThreshold:     util.PtrInt32(int32(0)),
		AutoGenerateRouteRule: util.PtrBool(false),
	}
	err := VPC_CLIENT.CreateDedicatedGatewayHealthCheck(createDedicatedGatewayHealthCheckRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateDnatRule(t *testing.T) {
	createDnatRuleRequest := &CreateDnatRuleRequest{
		NatId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		RuleName:         util.PtrString(""),
		PublicIpAddress:  util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		Protocol:         util.PtrString(""),
		PublicPort:       util.PtrInt32(int32(0)),
		PrivatePort:      util.PtrInt32(int32(0)),
		PublicPortRange:  util.PtrString(""),
		PrivatePortRange: util.PtrString(""),
	}
	result := &CreateDnatRuleResponse{}
	result, err := VPC_CLIENT.CreateDnatRule(createDnatRuleRequest)
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
func TestClient_CreateEgressOnlyRule(t *testing.T) {
	createEgressOnlyRuleRequest := &CreateEgressOnlyRuleRequest{
		GatewayId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		Cidr:        util.PtrString(""),
	}
	result := &CreateEgressOnlyRuleResponse{}
	result, err := VPC_CLIENT.CreateEgressOnlyRule(createEgressOnlyRuleRequest)
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
func TestClient_CreateEni(t *testing.T) {
	createEniRequest := &CreateEniRequest{
		ClientToken:                 util.PtrString(""),
		Name:                        util.PtrString(""),
		SubnetId:                    util.PtrString(""),
		SecurityGroupIds:            []*string{},
		EnterpriseSecurityGroupIds:  []*string{},
		PrivateIpSet:                []*PrivateIP{},
		Ipv6PrivateIpSet:            []*PrivateIP{},
		Description:                 util.PtrString(""),
		NetworkInterfaceTrafficMode: util.PtrString(""),
	}
	result := &CreateEniResponse{}
	result, err := VPC_CLIENT.CreateEni(createEniRequest)
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
func TestClient_CreateEnterpriseSecurityGroup(t *testing.T) {
	createEnterpriseSecurityGroupRequest := &CreateEnterpriseSecurityGroupRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Desc:        util.PtrString(""),
		Rules:       []*EnterpriseSecurityGroupRuleModel{},
		Tags:        []*TagModel{},
	}
	result := &CreateEnterpriseSecurityGroupResponse{}
	result, err := VPC_CLIENT.CreateEnterpriseSecurityGroup(createEnterpriseSecurityGroupRequest)
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
func TestClient_CreateGatewayLimitRules(t *testing.T) {
	createGatewayLimitRulesRequest := &CreateGatewayLimitRulesRequest{
		ClientToken:    util.PtrString(""),
		IpVersion:      util.PtrString(""),
		Name:           util.PtrString(""),
		Description:    util.PtrString(""),
		ServiceType:    util.PtrString(""),
		SubServiceType: util.PtrString(""),
		PeerRegion:     util.PtrString(""),
		ResourceId:     util.PtrString(""),
		Direction:      util.PtrString(""),
		Cidr:           util.PtrString(""),
		Bandwidth:      util.PtrInt32(int32(0)),
	}
	result := &CreateGatewayLimitRulesResponse{}
	result, err := VPC_CLIENT.CreateGatewayLimitRules(createGatewayLimitRulesRequest)
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
func TestClient_CreateHaVip(t *testing.T) {
	createHaVipRequest := &CreateHaVipRequest{
		ClientToken:      util.PtrString(""),
		Name:             util.PtrString(""),
		SubnetId:         util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		Description:      util.PtrString(""),
	}
	result := &CreateHaVipResponse{}
	result, err := VPC_CLIENT.CreateHaVip(createHaVipRequest)
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
func TestClient_CreateIpGroup(t *testing.T) {
	createIpGroupRequest := &CreateIpGroupRequest{
		ClientToken:   util.PtrString(""),
		Name:          util.PtrString(""),
		IpVersion:     util.PtrString(""),
		IpAddressInfo: []*TemplateIpAddressInfo{},
		Description:   util.PtrString(""),
	}
	result := &CreateIpGroupResponse{}
	result, err := VPC_CLIENT.CreateIpGroup(createIpGroupRequest)
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
func TestClient_CreateIpReserved(t *testing.T) {
	createIpReservedRequest := &CreateIpReservedRequest{
		ClientToken: util.PtrString(""),
		SubnetId:    util.PtrString(""),
		IpCidr:      util.PtrString(""),
		IpVersion:   util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
	}
	result := &CreateIpReservedResponse{}
	result, err := VPC_CLIENT.CreateIpReserved(createIpReservedRequest)
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
func TestClient_CreateIpSet(t *testing.T) {
	createIpSetRequest := &CreateIpSetRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		IpVersion:   util.PtrString(""),
		IpSetIds:    []*string{},
		Description: util.PtrString(""),
	}
	result := &CreateIpSetResponse{}
	result, err := VPC_CLIENT.CreateIpSet(createIpSetRequest)
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
func TestClient_CreateIpv6Gateway(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createIpv6GatewayRequest := &CreateIpv6GatewayRequest{
		ClientToken:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		Name:            util.PtrString(""),
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Billing:         Billing,
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
		DeleteProtect:   util.PtrBool(false),
	}
	result := &CreateIpv6GatewayResponse{}
	result, err := VPC_CLIENT.CreateIpv6Gateway(createIpv6GatewayRequest)
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
func TestClient_CreateNat(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	SessionConfig := &SessionConfig{
		TcpTimeout:  util.PtrInt32(int32(0)),
		UdpTimeout:  util.PtrInt32(int32(0)),
		IcmpTimeout: util.PtrInt32(int32(0)),
	}
	createNatRequest := &CreateNatRequest{
		ClientToken:     util.PtrString(""),
		Name:            util.PtrString(""),
		VpcId:           util.PtrString(""),
		CuNum:           util.PtrInt32(int32(0)),
		IpVersion:       util.PtrString(""),
		BindEips:        []*string{},
		Billing:         Billing,
		SessionConfig:   SessionConfig,
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
		DeleteProtect:   util.PtrBool(false),
	}
	result := &CreateNatResponse{}
	result, err := VPC_CLIENT.CreateNat(createNatRequest)
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
func TestClient_CreatePeerConn(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createPeerConnRequest := &CreatePeerConnRequest{
		ClientToken:     util.PtrString(""),
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Description:     util.PtrString(""),
		LocalIfName:     util.PtrString(""),
		LocalVpcId:      util.PtrString(""),
		PeerAccountId:   util.PtrString(""),
		PeerVpcId:       util.PtrString(""),
		PeerRegion:      util.PtrString(""),
		PeerIfName:      util.PtrString(""),
		Billing:         Billing,
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
		DeleteProtect:   util.PtrBool(false),
	}
	result := &CreatePeerConnResponse{}
	result, err := VPC_CLIENT.CreatePeerConn(createPeerConnRequest)
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
func TestClient_CreateProbe(t *testing.T) {
	createProbeRequest := &CreateProbeRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		VpcId:       util.PtrString(""),
		SubnetId:    util.PtrString(""),
		Protocol:    util.PtrString(""),
		Frequency:   util.PtrInt32(int32(0)),
		SourceIps:   []*string{},
		SourceIpNum: util.PtrInt32(int32(0)),
		DestIp:      util.PtrString(""),
		DestPort:    util.PtrInt32(int32(0)),
		Payload:     util.PtrString(""),
	}
	result := &CreateProbeResponse{}
	result, err := VPC_CLIENT.CreateProbe(createProbeRequest)
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
func TestClient_CreateRateLimitRule(t *testing.T) {
	createRateLimitRuleRequest := &CreateRateLimitRuleRequest{
		GatewayId:              util.PtrString(""),
		ClientToken:            util.PtrString(""),
		Ipv6Address:            util.PtrString(""),
		IngressBandwidthInMbps: util.PtrInt32(int32(0)),
		EgressBandwidthInMbps:  util.PtrInt32(int32(0)),
	}
	result := &CreateRateLimitRuleResponse{}
	result, err := VPC_CLIENT.CreateRateLimitRule(createRateLimitRuleRequest)
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
func TestClient_CreateRoutingRules(t *testing.T) {
	createRoutingRulesRequest := &CreateRoutingRulesRequest{
		ClientToken:        util.PtrString(""),
		RouteTableId:       util.PtrString(""),
		SourceAddress:      util.PtrString(""),
		DestinationAddress: util.PtrString(""),
		NexthopId:          util.PtrString(""),
		NexthopType:        util.PtrString(""),
		NextHopList:        []*NextHop{},
		Description:        util.PtrString(""),
	}
	result := &CreateRoutingRulesResponse{}
	result, err := VPC_CLIENT.CreateRoutingRules(createRoutingRulesRequest)
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
func TestClient_CreateSecurityGroup(t *testing.T) {
	createSecurityGroupRequest := &CreateSecurityGroupRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		VpcId:       util.PtrString(""),
		Desc:        util.PtrString(""),
		Rules:       []*SecurityGroupRuleModel{},
		Tags:        []*TagModel{},
	}
	result := &CreateSecurityGroupResponse{}
	result, err := VPC_CLIENT.CreateSecurityGroup(createSecurityGroupRequest)
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
func TestClient_CreateSnatRule(t *testing.T) {
	createSnatRuleRequest := &CreateSnatRuleRequest{
		NatId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		RuleName:         util.PtrString(""),
		PublicIpsAddress: []*string{},
		SourceCIDR:       util.PtrString(""),
	}
	result := &CreateSnatRuleResponse{}
	result, err := VPC_CLIENT.CreateSnatRule(createSnatRuleRequest)
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
func TestClient_CreateSslVpnServer(t *testing.T) {
	createSslVpnServerRequest := &CreateSslVpnServerRequest{
		VpnId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SslVpnServerName: util.PtrString(""),
		InterfaceType:    util.PtrString(""),
		LocalSubnets:     []*string{},
		RemoteSubnet:     util.PtrString(""),
		ClientDns:        util.PtrString(""),
	}
	result := &CreateSslVpnServerResponse{}
	result, err := VPC_CLIENT.CreateSslVpnServer(createSslVpnServerRequest)
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
func TestClient_CreateSubnet(t *testing.T) {
	createSubnetRequest := &CreateSubnetRequest{
		ClientToken:      util.PtrString(""),
		Name:             util.PtrString(""),
		EnableIpv6:       util.PtrBool(false),
		ZoneName:         util.PtrString(""),
		Cidr:             util.PtrString(""),
		VpcId:            util.PtrString(""),
		VpcSecondaryCidr: util.PtrString(""),
		SubnetType:       util.PtrString(""),
		Description:      util.PtrString(""),
		Tags:             []*TagModel{},
	}
	result := &CreateSubnetResponse{}
	result, err := VPC_CLIENT.CreateSubnet(createSubnetRequest)
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
func TestClient_CreateUserGateway(t *testing.T) {
	createUserGatewayRequest := &CreateUserGatewayRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Ip:          util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateUserGatewayResponse{}
	result, err := VPC_CLIENT.CreateUserGateway(createUserGatewayRequest)
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
func TestClient_CreateVpc(t *testing.T) {
	createVpcRequest := &CreateVpcRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Cidr:        util.PtrString(""),
		EnableIpv6:  util.PtrBool(false),
		Tags:        []*TagModel{},
	}
	result := &CreateVpcResponse{}
	result, err := VPC_CLIENT.CreateVpc(createVpcRequest)
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
func TestClient_CreateVpn(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createVpnRequest := &CreateVpnRequest{
		ClientToken:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		SubnetId:        util.PtrString(""),
		VpnName:         util.PtrString(""),
		VpcType:         util.PtrString(""),
		Description:     util.PtrString(""),
		Eip:             util.PtrString(""),
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
		Billing:         Billing,
		MaxConnection:   util.PtrInt32(int32(0)),
		DeleteProtect:   util.PtrBool(false),
	}
	result := &CreateVpnResponse{}
	result, err := VPC_CLIENT.CreateVpn(createVpnRequest)
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
func TestClient_CreateVpnTunnel(t *testing.T) {
	IkeConfig := &IkeConfig{
		IkeVersion:  util.PtrString(""),
		IkeMode:     util.PtrString(""),
		IkeEncAlg:   util.PtrString(""),
		IkeAuthAlg:  util.PtrString(""),
		IkePfs:      util.PtrString(""),
		IkeLifeTime: util.PtrString(""),
	}
	IpsecConfig := &IpsecConfig{
		IpsecEncAlg:   util.PtrString(""),
		IpsecAuthAlg:  util.PtrString(""),
		IpsecPfs:      util.PtrString(""),
		IpsecLifetime: util.PtrString(""),
	}
	createVpnTunnelRequest := &CreateVpnTunnelRequest{
		VpnId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		SecretKey:     util.PtrString(""),
		LocalSubnets:  []*string{},
		CgwId:         util.PtrString(""),
		RemoteSubnets: []*string{},
		Description:   util.PtrString(""),
		VpnConnName:   util.PtrString(""),
		IkeConfig:     IkeConfig,
		IpsecConfig:   IpsecConfig,
	}
	result := &CreateVpnTunnelResponse{}
	result, err := VPC_CLIENT.CreateVpnTunnel(createVpnTunnelRequest)
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
func TestClient_DeleteAclRule(t *testing.T) {
	deleteAclRuleRequest := &DeleteAclRuleRequest{
		AclRuleId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteAclRule(deleteAclRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDnatRule(t *testing.T) {
	deleteDnatRuleRequest := &DeleteDnatRuleRequest{
		NatId:       util.PtrString(""),
		RuleId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteDnatRule(deleteDnatRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteEniIp(t *testing.T) {
	deleteEniIpRequest := &DeleteEniIpRequest{
		EniId:            util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		ClientToken:      util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteEniIp(deleteEniIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteEnterpriseSecurityGroup(t *testing.T) {
	deleteEnterpriseSecurityGroupRequest := &DeleteEnterpriseSecurityGroupRequest{
		EnterpriseSecurityGroupId: util.PtrString(""),
		ClientToken:               util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteEnterpriseSecurityGroup(deleteEnterpriseSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteEnterpriseSecurityGroupRules(t *testing.T) {
	deleteEnterpriseSecurityGroupRulesRequest := &DeleteEnterpriseSecurityGroupRulesRequest{
		EnterpriseSecurityGroupRuleId: util.PtrString(""),
		ClientToken:                   util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteEnterpriseSecurityGroupRules(deleteEnterpriseSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteGatewayLimitRule(t *testing.T) {
	deleteGatewayLimitRuleRequest := &DeleteGatewayLimitRuleRequest{
		GlrId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteGatewayLimitRule(deleteGatewayLimitRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteHaVip(t *testing.T) {
	deleteHaVipRequest := &DeleteHaVipRequest{
		HaVipId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteHaVip(deleteHaVipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpGroup(t *testing.T) {
	deleteIpGroupRequest := &DeleteIpGroupRequest{
		IpSetId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpGroup(deleteIpGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpReserve(t *testing.T) {
	deleteIpReserveRequest := &DeleteIpReserveRequest{
		IpReserveId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpReserve(deleteIpReserveRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpSet(t *testing.T) {
	deleteIpSetRequest := &DeleteIpSetRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpSet(deleteIpSetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpv6Gateway(t *testing.T) {
	deleteIpv6GatewayRequest := &DeleteIpv6GatewayRequest{
		GatewayId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpv6Gateway(deleteIpv6GatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpv6GatewayEgressOnlyRule(t *testing.T) {
	deleteIpv6GatewayEgressOnlyRuleRequest := &DeleteIpv6GatewayEgressOnlyRuleRequest{
		GatewayId:        util.PtrString(""),
		EgressOnlyRuleId: util.PtrString(""),
		ClientToken:      util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpv6GatewayEgressOnlyRule(deleteIpv6GatewayEgressOnlyRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpv6GatewayRateLimitRule(t *testing.T) {
	deleteIpv6GatewayRateLimitRuleRequest := &DeleteIpv6GatewayRateLimitRuleRequest{
		GatewayId:       util.PtrString(""),
		RateLimitRuleId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpv6GatewayRateLimitRule(deleteIpv6GatewayRateLimitRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteProbe(t *testing.T) {
	deleteProbeRequest := &DeleteProbeRequest{
		ProbeId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteProbe(deleteProbeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRoutingRules(t *testing.T) {
	deleteRoutingRulesRequest := &DeleteRoutingRulesRequest{
		RouteRuleId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteRoutingRules(deleteRoutingRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSecurityGroup(t *testing.T) {
	deleteSecurityGroupRequest := &DeleteSecurityGroupRequest{
		SecurityGroupId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSecurityGroup(deleteSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSecurityGroupRules(t *testing.T) {
	deleteSecurityGroupRulesRequest := &DeleteSecurityGroupRulesRequest{
		SecurityGroupRuleId: util.PtrString(""),
		ClientToken:         util.PtrString(""),
		SgVersion:           util.PtrInt64(int64(0)),
	}
	err := VPC_CLIENT.DeleteSecurityGroupRules(deleteSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSnatRule(t *testing.T) {
	deleteSnatRuleRequest := &DeleteSnatRuleRequest{
		NatId:       util.PtrString(""),
		RuleId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSnatRule(deleteSnatRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSslVpnServer(t *testing.T) {
	deleteSslVpnServerRequest := &DeleteSslVpnServerRequest{
		VpnId:          util.PtrString(""),
		SslVpnServerId: util.PtrString(""),
		ClientToken:    util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSslVpnServer(deleteSslVpnServerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSslVpnUser(t *testing.T) {
	deleteSslVpnUserRequest := &DeleteSslVpnUserRequest{
		VpnId:       util.PtrString(""),
		UserId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSslVpnUser(deleteSslVpnUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSubnet(t *testing.T) {
	deleteSubnetRequest := &DeleteSubnetRequest{
		SubnetId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSubnet(deleteSubnetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteUserGateway(t *testing.T) {
	deleteUserGatewayRequest := &DeleteUserGatewayRequest{
		CgwId: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteUserGateway(deleteUserGatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteVpc(t *testing.T) {
	deleteVpcRequest := &DeleteVpcRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteVpc(deleteVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteVpnTunnel(t *testing.T) {
	deleteVpnTunnelRequest := &DeleteVpnTunnelRequest{
		VpnConnId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteVpnTunnel(deleteVpnTunnelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DetachEniInstance(t *testing.T) {
	detachEniInstanceRequest := &DetachEniInstanceRequest{
		EniId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
	}
	err := VPC_CLIENT.DetachEniInstance(detachEniInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetEniDetail(t *testing.T) {
	getEniDetailRequest := &GetEniDetailRequest{
		EniId: util.PtrString(""),
	}
	result := &GetEniDetailResponse{}
	result, err := VPC_CLIENT.GetEniDetail(getEniDetailRequest)
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
func TestClient_GetEniStatus(t *testing.T) {
	getEniStatusRequest := &GetEniStatusRequest{
		EniId: util.PtrString(""),
	}
	result := &GetEniStatusResponse{}
	result, err := VPC_CLIENT.GetEniStatus(getEniStatusRequest)
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
func TestClient_GetHaVipDetail(t *testing.T) {
	getHaVipDetailRequest := &GetHaVipDetailRequest{
		HaVipId: util.PtrString(""),
	}
	result := &GetHaVipDetailResponse{}
	result, err := VPC_CLIENT.GetHaVipDetail(getHaVipDetailRequest)
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
func TestClient_GetNat(t *testing.T) {
	getNatRequest := &GetNatRequest{
		NatId: util.PtrString(""),
	}
	result := &GetNatResponse{}
	result, err := VPC_CLIENT.GetNat(getNatRequest)
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
func TestClient_GetPeerConn(t *testing.T) {
	getPeerConnRequest := &GetPeerConnRequest{
		PeerConnId: util.PtrString(""),
		Role:       util.PtrString(""),
	}
	result := &GetPeerConnResponse{}
	result, err := VPC_CLIENT.GetPeerConn(getPeerConnRequest)
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
func TestClient_GetProbeDetail(t *testing.T) {
	getProbeDetailRequest := &GetProbeDetailRequest{
		ProbeId: util.PtrString(""),
	}
	result := &GetProbeDetailResponse{}
	result, err := VPC_CLIENT.GetProbeDetail(getProbeDetailRequest)
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
func TestClient_GetSecurityGroupDetails(t *testing.T) {
	getSecurityGroupDetailsRequest := &GetSecurityGroupDetailsRequest{
		SecurityGroupId: util.PtrString(""),
	}
	result := &GetSecurityGroupDetailsResponse{}
	result, err := VPC_CLIENT.GetSecurityGroupDetails(getSecurityGroupDetailsRequest)
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
func TestClient_GetVpcResourceIpInfo(t *testing.T) {
	getVpcResourceIpInfoRequest := &GetVpcResourceIpInfoRequest{
		VpcId:        util.PtrString(""),
		SubnetId:     util.PtrString(""),
		ResourceType: util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &GetVpcResourceIpInfoResponse{}
	result, err := VPC_CLIENT.GetVpcResourceIpInfo(getVpcResourceIpInfoRequest)
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
func TestClient_ListDnatRule(t *testing.T) {
	listDnatRuleRequest := &ListDnatRuleRequest{
		NatId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListDnatRuleResponse{}
	result, err := VPC_CLIENT.ListDnatRule(listDnatRuleRequest)
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
func TestClient_ListEgressOnlyRule(t *testing.T) {
	listEgressOnlyRuleRequest := &ListEgressOnlyRuleRequest{
		GatewayId: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &ListEgressOnlyRuleResponse{}
	result, err := VPC_CLIENT.ListEgressOnlyRule(listEgressOnlyRuleRequest)
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
func TestClient_ListEni(t *testing.T) {
	listEniRequest := &ListEniRequest{
		VpcId:            util.PtrString(""),
		InstanceId:       util.PtrString(""),
		Name:             util.PtrString(""),
		PrivateIpAddress: []*string{},
		Marker:           util.PtrString(""),
		MaxKeys:          util.PtrInt32(int32(0)),
	}
	result := &ListEniResponse{}
	result, err := VPC_CLIENT.ListEni(listEniRequest)
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
func TestClient_ListHaVip(t *testing.T) {
	listHaVipRequest := &ListHaVipRequest{
		VpcId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListHaVipResponse{}
	result, err := VPC_CLIENT.ListHaVip(listHaVipRequest)
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
func TestClient_ListIpReserve(t *testing.T) {
	listIpReserveRequest := &ListIpReserveRequest{
		SubnetId: util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
	}
	result := &ListIpReserveResponse{}
	result, err := VPC_CLIENT.ListIpReserve(listIpReserveRequest)
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
func TestClient_ListNat(t *testing.T) {
	listNatRequest := &ListNatRequest{
		VpcId:   util.PtrString(""),
		NatId:   util.PtrString(""),
		Name:    util.PtrString(""),
		Ip:      util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListNatResponse{}
	result, err := VPC_CLIENT.ListNat(listNatRequest)
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
func TestClient_ListPeerConn(t *testing.T) {
	listPeerConnRequest := &ListPeerConnRequest{
		VpcId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListPeerConnResponse{}
	result, err := VPC_CLIENT.ListPeerConn(listPeerConnRequest)
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
func TestClient_ListProbes(t *testing.T) {
	listProbesRequest := &ListProbesRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListProbesResponse{}
	result, err := VPC_CLIENT.ListProbes(listProbesRequest)
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
func TestClient_ListRateLimitRule(t *testing.T) {
	listRateLimitRuleRequest := &ListRateLimitRuleRequest{
		GatewayId: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &ListRateLimitRuleResponse{}
	result, err := VPC_CLIENT.ListRateLimitRule(listRateLimitRuleRequest)
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
func TestClient_ListSnatRule(t *testing.T) {
	listSnatRuleRequest := &ListSnatRuleRequest{
		NatId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListSnatRuleResponse{}
	result, err := VPC_CLIENT.ListSnatRule(listSnatRuleRequest)
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
func TestClient_ModifyGatewayLimitRules(t *testing.T) {
	modifyGatewayLimitRulesRequest := &ModifyGatewayLimitRulesRequest{
		GlrId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Bandwidth:   util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.ModifyGatewayLimitRules(modifyGatewayLimitRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyNat(t *testing.T) {
	modifyNatRequest := &ModifyNatRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
	}
	err := VPC_CLIENT.ModifyNat(modifyNatRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_NatBindEip(t *testing.T) {
	natBindEipRequest := &NatBindEipRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		BindEips:    []*string{},
	}
	err := VPC_CLIENT.NatBindEip(natBindEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_NatUnBindEip(t *testing.T) {
	natUnBindEipRequest := &NatUnBindEipRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		BindEips:    []*string{},
	}
	err := VPC_CLIENT.NatUnBindEip(natUnBindEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_OpenPeerConnSyncDns(t *testing.T) {
	openPeerConnSyncDnsRequest := &OpenPeerConnSyncDnsRequest{
		PeerConnId:  util.PtrString(""),
		Role:        util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.OpenPeerConnSyncDns(openPeerConnSyncDnsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_OpenVpcRelay(t *testing.T) {
	openVpcRelayRequest := &OpenVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.OpenVpcRelay(openVpcRelayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PurchaseReservedNat(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	purchaseReservedNatRequest := &PurchaseReservedNatRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     Billing,
	}
	err := VPC_CLIENT.PurchaseReservedNat(purchaseReservedNatRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryAcl(t *testing.T) {
	queryAclRequest := &QueryAclRequest{
		VpcId: util.PtrString(""),
	}
	result := &QueryAclResponse{}
	result, err := VPC_CLIENT.QueryAcl(queryAclRequest)
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
func TestClient_QueryAclRules(t *testing.T) {
	queryAclRulesRequest := &QueryAclRulesRequest{
		SubnetId: util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
	}
	result := &QueryAclRulesResponse{}
	result, err := VPC_CLIENT.QueryAclRules(queryAclRulesRequest)
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
func TestClient_QueryEnterpriseSecurityGroupList(t *testing.T) {
	queryEnterpriseSecurityGroupListRequest := &QueryEnterpriseSecurityGroupListRequest{
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		InstanceId: util.PtrString(""),
	}
	result := &QueryEnterpriseSecurityGroupListResponse{}
	result, err := VPC_CLIENT.QueryEnterpriseSecurityGroupList(queryEnterpriseSecurityGroupListRequest)
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
func TestClient_QueryIpGroupDetail(t *testing.T) {
	queryIpGroupDetailRequest := &QueryIpGroupDetailRequest{
		IpSetId: util.PtrString(""),
	}
	result := &QueryIpGroupDetailResponse{}
	result, err := VPC_CLIENT.QueryIpGroupDetail(queryIpGroupDetailRequest)
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
func TestClient_QueryIpGroupList(t *testing.T) {
	queryIpGroupListRequest := &QueryIpGroupListRequest{
		IpVersion: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &QueryIpGroupListResponse{}
	result, err := VPC_CLIENT.QueryIpGroupList(queryIpGroupListRequest)
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
func TestClient_QueryIpSetDetail(t *testing.T) {
	queryIpSetDetailRequest := &QueryIpSetDetailRequest{
		IpGroupId: util.PtrString(""),
	}
	result := &QueryIpSetDetailResponse{}
	result, err := VPC_CLIENT.QueryIpSetDetail(queryIpSetDetailRequest)
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
func TestClient_QueryIpSetList(t *testing.T) {
	queryIpSetListRequest := &QueryIpSetListRequest{
		IpVersion: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &QueryIpSetListResponse{}
	result, err := VPC_CLIENT.QueryIpSetList(queryIpSetListRequest)
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
func TestClient_QueryIpv6Gateway(t *testing.T) {
	queryIpv6GatewayRequest := &QueryIpv6GatewayRequest{
		VpcId: util.PtrString(""),
	}
	result := &QueryIpv6GatewayResponse{}
	result, err := VPC_CLIENT.QueryIpv6Gateway(queryIpv6GatewayRequest)
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
func TestClient_QueryRoutingRules(t *testing.T) {
	queryRoutingRulesRequest := &QueryRoutingRulesRequest{
		RouteTableId: util.PtrString(""),
		VpcId:        util.PtrString(""),
		Marker:       util.PtrString(""),
		MaxKeys:      util.PtrInt32(int32(0)),
	}
	result := &QueryRoutingRulesResponse{}
	result, err := VPC_CLIENT.QueryRoutingRules(queryRoutingRulesRequest)
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
func TestClient_QueryRoutingTable(t *testing.T) {
	queryRoutingTableRequest := &QueryRoutingTableRequest{
		RouteTableId: util.PtrString(""),
		VpcId:        util.PtrString(""),
	}
	result := &QueryRoutingTableResponse{}
	result, err := VPC_CLIENT.QueryRoutingTable(queryRoutingTableRequest)
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
func TestClient_QuerySecurityGroupsList(t *testing.T) {
	querySecurityGroupsListRequest := &QuerySecurityGroupsListRequest{
		Marker:           util.PtrString(""),
		MaxKeys:          util.PtrInt32(int32(0)),
		InstanceId:       util.PtrString(""),
		VpcId:            util.PtrString(""),
		SecurityGroupId:  util.PtrString(""),
		SecurityGroupIds: util.PtrString(""),
	}
	result := &QuerySecurityGroupsListResponse{}
	result, err := VPC_CLIENT.QuerySecurityGroupsList(querySecurityGroupsListRequest)
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
func TestClient_QuerySpecifiedSubnet(t *testing.T) {
	querySpecifiedSubnetRequest := &QuerySpecifiedSubnetRequest{
		SubnetId: util.PtrString(""),
	}
	result := &QuerySpecifiedSubnetResponse{}
	result, err := VPC_CLIENT.QuerySpecifiedSubnet(querySpecifiedSubnetRequest)
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
func TestClient_QuerySpecifiedVpc(t *testing.T) {
	querySpecifiedVpcRequest := &QuerySpecifiedVpcRequest{
		VpcId: util.PtrString(""),
	}
	result := &QuerySpecifiedVpcResponse{}
	result, err := VPC_CLIENT.QuerySpecifiedVpc(querySpecifiedVpcRequest)
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
func TestClient_QuerySslVpnServer(t *testing.T) {
	querySslVpnServerRequest := &QuerySslVpnServerRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	result := &QuerySslVpnServerResponse{}
	result, err := VPC_CLIENT.QuerySslVpnServer(querySslVpnServerRequest)
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
func TestClient_QuerySslVpnUsers(t *testing.T) {
	querySslVpnUsersRequest := &QuerySslVpnUsersRequest{
		VpnId:    util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
		UserName: util.PtrString(""),
	}
	result := &QuerySslVpnUsersResponse{}
	result, err := VPC_CLIENT.QuerySslVpnUsers(querySslVpnUsersRequest)
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
func TestClient_QuerySubnetList(t *testing.T) {
	querySubnetListRequest := &QuerySubnetListRequest{
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		VpcId:      util.PtrString(""),
		ZoneName:   util.PtrString(""),
		SubnetType: util.PtrString(""),
		SubnetIds:  util.PtrString(""),
	}
	result := &QuerySubnetListResponse{}
	result, err := VPC_CLIENT.QuerySubnetList(querySubnetListRequest)
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
func TestClient_QueryTheDetailsOfTheDedicatedGateway(t *testing.T) {
	queryTheDetailsOfTheDedicatedGatewayRequest := &QueryTheDetailsOfTheDedicatedGatewayRequest{
		EtGatewayId: util.PtrString(""),
	}
	result := &QueryTheDetailsOfTheDedicatedGatewayResponse{}
	result, err := VPC_CLIENT.QueryTheDetailsOfTheDedicatedGateway(queryTheDetailsOfTheDedicatedGatewayRequest)
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
func TestClient_QueryTheListOfDedicatedLineGateways(t *testing.T) {
	queryTheListOfDedicatedLineGatewaysRequest := &QueryTheListOfDedicatedLineGatewaysRequest{
		VpcId:       util.PtrString(""),
		EtGatewayId: util.PtrString(""),
		Name:        util.PtrString(""),
		Status:      util.PtrString(""),
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfDedicatedLineGatewaysResponse{}
	result, err := VPC_CLIENT.QueryTheListOfDedicatedLineGateways(queryTheListOfDedicatedLineGatewaysRequest)
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
func TestClient_QueryVpcIntranetIp(t *testing.T) {
	queryVpcIntranetIpRequest := &QueryVpcIntranetIpRequest{
		VpcId:              util.PtrString(""),
		PrivateIpAddresses: []*string{},
		PrivateIpRange:     util.PtrString(""),
	}
	result := &QueryVpcIntranetIpResponse{}
	result, err := VPC_CLIENT.QueryVpcIntranetIp(queryVpcIntranetIpRequest)
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
func TestClient_QueryVpcList(t *testing.T) {
	queryVpcListRequest := &QueryVpcListRequest{
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		IsDefault: util.PtrBool(false),
		VpcIds:    util.PtrString(""),
	}
	result := &QueryVpcListResponse{}
	result, err := VPC_CLIENT.QueryVpcList(queryVpcListRequest)
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
func TestClient_QueryVpnList(t *testing.T) {
	queryVpnListRequest := &QueryVpnListRequest{
		VpcId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		Eip:     util.PtrString(""),
		Type:    util.PtrString(""),
	}
	result := &QueryVpnListResponse{}
	result, err := VPC_CLIENT.QueryVpnList(queryVpnListRequest)
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
func TestClient_RefundPeerConn(t *testing.T) {
	refundPeerConnRequest := &RefundPeerConnRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.RefundPeerConn(refundPeerConnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RejectPeerConn(t *testing.T) {
	rejectPeerConnRequest := &RejectPeerConnRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.RejectPeerConn(rejectPeerConnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseDedicatedGateway(t *testing.T) {
	releaseDedicatedGatewayRequest := &ReleaseDedicatedGatewayRequest{
		EtGatewayId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ReleaseDedicatedGateway(releaseDedicatedGatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseNat(t *testing.T) {
	releaseNatRequest := &ReleaseNatRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ReleaseNat(releaseNatRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleasePeerConn(t *testing.T) {
	releasePeerConnRequest := &ReleasePeerConnRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ReleasePeerConn(releasePeerConnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseVpn(t *testing.T) {
	releaseVpnRequest := &ReleaseVpnRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ReleaseVpn(releaseVpnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveEni(t *testing.T) {
	removeEniRequest := &RemoveEniRequest{
		EniId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.RemoveEni(removeEniRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveIpAddressFromIpGroup(t *testing.T) {
	removeIpAddressFromIpGroupRequest := &RemoveIpAddressFromIpGroupRequest{
		IpSetId:       util.PtrString(""),
		ClientToken:   util.PtrString(""),
		IpAddressInfo: []*string{},
	}
	err := VPC_CLIENT.RemoveIpAddressFromIpGroup(removeIpAddressFromIpGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveIpGroupFromIpSet(t *testing.T) {
	removeIpGroupFromIpSetRequest := &RemoveIpGroupFromIpSetRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpSetIds:    []*string{},
	}
	err := VPC_CLIENT.RemoveIpGroupFromIpSet(removeIpGroupFromIpSetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RenewPeerConn(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	renewPeerConnRequest := &RenewPeerConnRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     Billing,
	}
	err := VPC_CLIENT.RenewPeerConn(renewPeerConnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RenewVpn(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	renewVpnRequest := &RenewVpnRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     Billing,
	}
	err := VPC_CLIENT.RenewVpn(renewVpnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeIpv6Gateway(t *testing.T) {
	resizeIpv6GatewayRequest := &ResizeIpv6GatewayRequest{
		GatewayId:       util.PtrString(""),
		ClientToken:     util.PtrString(""),
		BandwidthInMbps: util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.ResizeIpv6Gateway(resizeIpv6GatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeNat(t *testing.T) {
	resizeNatRequest := &ResizeNatRequest{
		NatId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		CuNum:       util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.ResizeNat(resizeNatRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RevokeSecurityGroupRules(t *testing.T) {
	Rule := &SecurityGroupRuleModel{
		Remark:              util.PtrString(""),
		Direction:           util.PtrString(""),
		Ethertype:           util.PtrString(""),
		PortRange:           util.PtrString(""),
		Protocol:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		SourceIp:            util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		DestIp:              util.PtrString(""),
		SecurityGroupId:     util.PtrString(""),
		SecurityGroupRuleId: util.PtrString(""),
		CreatedTime:         util.PtrString(""),
		UpdatedTime:         util.PtrString(""),
	}
	revokeSecurityGroupRulesRequest := &RevokeSecurityGroupRulesRequest{
		SecurityGroupId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
		SgVersion:       util.PtrInt64(int64(0)),
		Rule:            Rule,
	}
	err := VPC_CLIENT.RevokeSecurityGroupRules(revokeSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SearchForVpnDetails(t *testing.T) {
	searchForVpnDetailsRequest := &SearchForVpnDetailsRequest{
		VpnId: util.PtrString(""),
	}
	result := &SearchForVpnDetailsResponse{}
	result, err := VPC_CLIENT.SearchForVpnDetails(searchForVpnDetailsRequest)
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
func TestClient_SearchVpnTunnel(t *testing.T) {
	searchVpnTunnelRequest := &SearchVpnTunnelRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	result := &SearchVpnTunnelResponse{}
	result, err := VPC_CLIENT.SearchVpnTunnel(searchVpnTunnelRequest)
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
func TestClient_UnbindEip(t *testing.T) {
	unbindEipRequest := &UnbindEipRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.UnbindEip(unbindEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindEniEip(t *testing.T) {
	unbindEniEipRequest := &UnbindEniEipRequest{
		EniId:           util.PtrString(""),
		ClientToken:     util.PtrString(""),
		PublicIpAddress: util.PtrString(""),
	}
	err := VPC_CLIENT.UnbindEniEip(unbindEniEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindHaVipEip(t *testing.T) {
	unbindHaVipEipRequest := &UnbindHaVipEipRequest{
		HaVipId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.UnbindHaVipEip(unbindHaVipEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindHaVipInstance(t *testing.T) {
	unbindHaVipInstanceRequest := &UnbindHaVipInstanceRequest{
		HaVipId:      util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceIds:  []*string{},
		InstanceType: util.PtrString(""),
	}
	err := VPC_CLIENT.UnbindHaVipInstance(unbindHaVipInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindPhysicalDedicatedLine(t *testing.T) {
	unbindPhysicalDedicatedLineRequest := &UnbindPhysicalDedicatedLineRequest{
		EtGatewayId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.UnbindPhysicalDedicatedLine(unbindPhysicalDedicatedLineRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAclRules(t *testing.T) {
	updateAclRulesRequest := &UpdateAclRulesRequest{
		AclRuleId:            util.PtrString(""),
		ClientToken:          util.PtrString(""),
		Description:          util.PtrString(""),
		Protocol:             util.PtrString(""),
		SourceIpAddress:      util.PtrString(""),
		DestinationIpAddress: util.PtrString(""),
		SourcePort:           util.PtrString(""),
		DestinationPort:      util.PtrString(""),
		Position:             util.PtrInt32(int32(0)),
		Action:               util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateAclRules(updateAclRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDedicatedGateway(t *testing.T) {
	updateDedicatedGatewayRequest := &UpdateDedicatedGatewayRequest{
		EtGatewayId:    util.PtrString(""),
		ClientToken:    util.PtrString(""),
		Name:           util.PtrString(""),
		Speed:          util.PtrInt32(int32(0)),
		Description:    util.PtrString(""),
		LocalCidrs:     []*string{},
		EnableIpv6:     util.PtrInt32(int32(0)),
		Ipv6LocalCidrs: []*string{},
	}
	err := VPC_CLIENT.UpdateDedicatedGateway(updateDedicatedGatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDeleteProtect(t *testing.T) {
	updateDeleteProtectRequest := &UpdateDeleteProtectRequest{
		GatewayId:     util.PtrString(""),
		ClientToken:   util.PtrString(""),
		DeleteProtect: util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdateDeleteProtect(updateDeleteProtectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDnatRule(t *testing.T) {
	updateDnatRuleRequest := &UpdateDnatRuleRequest{
		NatId:            util.PtrString(""),
		RuleId:           util.PtrString(""),
		ClientToken:      util.PtrString(""),
		RuleName:         util.PtrString(""),
		Protocol:         util.PtrString(""),
		PublicPort:       util.PtrInt32(int32(0)),
		PrivatePort:      util.PtrInt32(int32(0)),
		PublicPortRange:  util.PtrString(""),
		PrivatePortRange: util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		PublicIpAddress:  util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateDnatRule(updateDnatRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateEni(t *testing.T) {
	updateEniRequest := &UpdateEniRequest{
		EniId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateEni(updateEniRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateEniEnterpriseSecurityGroup(t *testing.T) {
	updateEniEnterpriseSecurityGroupRequest := &UpdateEniEnterpriseSecurityGroupRequest{
		EniId:                      util.PtrString(""),
		ClientToken:                util.PtrString(""),
		EnterpriseSecurityGroupIds: []*string{},
	}
	err := VPC_CLIENT.UpdateEniEnterpriseSecurityGroup(updateEniEnterpriseSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateEniSecurityGroup(t *testing.T) {
	updateEniSecurityGroupRequest := &UpdateEniSecurityGroupRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SecurityGroupIds: []*string{},
	}
	err := VPC_CLIENT.UpdateEniSecurityGroup(updateEniSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateEnterpriseSecurityGroupRules(t *testing.T) {
	updateEnterpriseSecurityGroupRulesRequest := &UpdateEnterpriseSecurityGroupRulesRequest{
		EnterpriseSecurityGroupRuleId: util.PtrString(""),
		ClientToken:                   util.PtrString(""),
		Remark:                        util.PtrString(""),
		PortRange:                     util.PtrString(""),
		SourcePortRange:               util.PtrString(""),
		SourceIp:                      util.PtrString(""),
		DestIp:                        util.PtrString(""),
		LocalIp:                       util.PtrString(""),
		RemoteIpSet:                   util.PtrString(""),
		RemoteIpGroup:                 util.PtrString(""),
		Action:                        util.PtrString(""),
		Priority:                      util.PtrInt32(int32(0)),
		Protocol:                      util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateEnterpriseSecurityGroupRules(updateEnterpriseSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateHaVip(t *testing.T) {
	updateHaVipRequest := &UpdateHaVipRequest{
		HaVipId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateHaVip(updateHaVipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIpGroup(t *testing.T) {
	updateIpGroupRequest := &UpdateIpGroupRequest{
		IpSetId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateIpGroup(updateIpGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIpSet(t *testing.T) {
	updateIpSetRequest := &UpdateIpSetRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateIpSet(updateIpSetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateNatReleaseProtectionSwitch(t *testing.T) {
	updateNatReleaseProtectionSwitchRequest := &UpdateNatReleaseProtectionSwitchRequest{
		NatId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		DeleteProtect: util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdateNatReleaseProtectionSwitch(updateNatReleaseProtectionSwitchRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePeerConn(t *testing.T) {
	updatePeerConnRequest := &UpdatePeerConnRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
		LocalIfId:   util.PtrString(""),
		Description: util.PtrString(""),
		LocalIfName: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdatePeerConn(updatePeerConnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePeerConnBandwidth(t *testing.T) {
	updatePeerConnBandwidthRequest := &UpdatePeerConnBandwidthRequest{
		PeerConnId:         util.PtrString(""),
		ClientToken:        util.PtrString(""),
		NewBandwidthInMbps: util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdatePeerConnBandwidth(updatePeerConnBandwidthRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePeerConnDeleteProtect(t *testing.T) {
	updatePeerConnDeleteProtectRequest := &UpdatePeerConnDeleteProtectRequest{
		PeerConnId:    util.PtrString(""),
		ClientToken:   util.PtrString(""),
		DeleteProtect: util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdatePeerConnDeleteProtect(updatePeerConnDeleteProtectRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateProbe(t *testing.T) {
	updateProbeRequest := &UpdateProbeRequest{
		ProbeId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		DestIp:      util.PtrString(""),
		DestPort:    util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdateProbe(updateProbeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRateLimitRule(t *testing.T) {
	updateRateLimitRuleRequest := &UpdateRateLimitRuleRequest{
		GatewayId:              util.PtrString(""),
		RateLimitRuleId:        util.PtrString(""),
		ClientToken:            util.PtrString(""),
		IngressBandwidthInMbps: util.PtrInt32(int32(0)),
		EgressBandwidthInMbps:  util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdateRateLimitRule(updateRateLimitRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRoutingRules(t *testing.T) {
	updateRoutingRulesRequest := &UpdateRoutingRulesRequest{
		RouteRuleId:        util.PtrString(""),
		ClientToken:        util.PtrString(""),
		SourceAddress:      util.PtrString(""),
		DestinationAddress: util.PtrString(""),
		NexthopId:          util.PtrString(""),
		Description:        util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateRoutingRules(updateRoutingRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSecurityGroupRules(t *testing.T) {
	updateSecurityGroupRulesRequest := &UpdateSecurityGroupRulesRequest{
		ClientToken:         util.PtrString(""),
		SgVersion:           util.PtrInt64(int64(0)),
		SecurityGroupRuleId: util.PtrString(""),
		Remark:              util.PtrString(""),
		PortRange:           util.PtrString(""),
		SourceIp:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		DestIp:              util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		Protocol:            util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateSecurityGroupRules(updateSecurityGroupRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSnatRule(t *testing.T) {
	updateSnatRuleRequest := &UpdateSnatRuleRequest{
		NatId:            util.PtrString(""),
		RuleId:           util.PtrString(""),
		ClientToken:      util.PtrString(""),
		RuleName:         util.PtrString(""),
		SourceCIDR:       util.PtrString(""),
		PublicIpsAddress: []*string{},
	}
	err := VPC_CLIENT.UpdateSnatRule(updateSnatRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSslVpnServer(t *testing.T) {
	updateSslVpnServerRequest := &UpdateSslVpnServerRequest{
		VpnId:            util.PtrString(""),
		SslVpnServerId:   util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SslVpnServerName: util.PtrString(""),
		LocalSubnets:     []*string{},
		RemoteSubnet:     util.PtrString(""),
		ClientDns:        util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateSslVpnServer(updateSslVpnServerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSslVpnUsers(t *testing.T) {
	updateSslVpnUsersRequest := &UpdateSslVpnUsersRequest{
		VpnId:       util.PtrString(""),
		UserId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		ClientName:  util.PtrString(""),
		Password:    util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateSslVpnUsers(updateSslVpnUsersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSubnet(t *testing.T) {
	updateSubnetRequest := &UpdateSubnetRequest{
		SubnetId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		EnableIpv6:  util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdateSubnet(updateSubnetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateUserGateway(t *testing.T) {
	updateUserGatewayRequest := &UpdateUserGatewayRequest{
		CgwId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateUserGateway(updateUserGatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpc(t *testing.T) {
	updateVpcRequest := &UpdateVpcRequest{
		VpcId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		Name:          util.PtrString(""),
		Description:   util.PtrString(""),
		EnableIpv6:    util.PtrBool(false),
		SecondaryCidr: []*string{},
	}
	err := VPC_CLIENT.UpdateVpc(updateVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpn(t *testing.T) {
	updateVpnRequest := &UpdateVpnRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		VpnName:     util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateVpn(updateVpnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpnReleaseProtection(t *testing.T) {
	updateVpnReleaseProtectionRequest := &UpdateVpnReleaseProtectionRequest{
		VpnId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		DeleteProtect: util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdateVpnReleaseProtection(updateVpnReleaseProtectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpnTunnel(t *testing.T) {
	IkeConfig := &IkeConfig{
		IkeVersion:  util.PtrString(""),
		IkeMode:     util.PtrString(""),
		IkeEncAlg:   util.PtrString(""),
		IkeAuthAlg:  util.PtrString(""),
		IkePfs:      util.PtrString(""),
		IkeLifeTime: util.PtrString(""),
	}
	IpsecConfig := &IpsecConfig{
		IpsecEncAlg:   util.PtrString(""),
		IpsecAuthAlg:  util.PtrString(""),
		IpsecPfs:      util.PtrString(""),
		IpsecLifetime: util.PtrString(""),
	}
	updateVpnTunnelRequest := &UpdateVpnTunnelRequest{
		VpnConnId:     util.PtrString(""),
		ClientToken:   util.PtrString(""),
		VpnId:         util.PtrString(""),
		SecretKey:     util.PtrString(""),
		LocalSubnets:  []*string{},
		CgwId:         util.PtrString(""),
		RemoteSubnets: []*string{},
		Description:   util.PtrString(""),
		VpnConnName:   util.PtrString(""),
		IkeConfig:     IkeConfig,
		IpsecConfig:   IpsecConfig,
	}
	err := VPC_CLIENT.UpdateVpnTunnel(updateVpnTunnelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserGatewayDetails(t *testing.T) {
	userGatewayDetailsRequest := &UserGatewayDetailsRequest{
		CgwId: util.PtrString(""),
	}
	result := &UserGatewayDetailsResponse{}
	result, err := VPC_CLIENT.UserGatewayDetails(userGatewayDetailsRequest)
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
func TestClient_UserGatewayList(t *testing.T) {
	userGatewayListRequest := &UserGatewayListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &UserGatewayListResponse{}
	result, err := VPC_CLIENT.UserGatewayList(userGatewayListRequest)
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
func TestClient_ViewGatewayLimitRules(t *testing.T) {
	viewGatewayLimitRulesRequest := &ViewGatewayLimitRulesRequest{
		ServiceType: util.PtrString(""),
		Name:        util.PtrString(""),
		GlrId:       util.PtrString(""),
		ResourceId:  util.PtrString(""),
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrString(""),
	}
	result := &ViewGatewayLimitRulesResponse{}
	result, err := VPC_CLIENT.ViewGatewayLimitRules(viewGatewayLimitRulesRequest)
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
