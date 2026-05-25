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

func TestClient_AcceptPeerToPeerConnectionApplications(t *testing.T) {
	acceptPeerToPeerConnectionApplicationsRequest := &AcceptPeerToPeerConnectionApplicationsRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.AcceptPeerToPeerConnectionApplications(acceptPeerToPeerConnectionApplicationsRequest)
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
func TestClient_AddElasticNetworkCardAuxiliaryIp(t *testing.T) {
	addElasticNetworkCardAuxiliaryIpRequest := &AddElasticNetworkCardAuxiliaryIpRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		IsIpv6:           util.PtrBool(false),
		PrivateIpAddress: util.PtrString(""),
	}
	result := &AddElasticNetworkCardAuxiliaryIpResponse{}
	result, err := VPC_CLIENT.AddElasticNetworkCardAuxiliaryIp(addElasticNetworkCardAuxiliaryIpRequest)
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
func TestClient_AddIpAddressGroupToIpAddressFamily(t *testing.T) {
	addIpAddressGroupToIpAddressFamilyRequest := &AddIpAddressGroupToIpAddressFamilyRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpSetIds:    []*string{},
	}
	err := VPC_CLIENT.AddIpAddressGroupToIpAddressFamily(addIpAddressGroupToIpAddressFamilyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddIpAddressesToTheIpAddressGroup(t *testing.T) {
	addIpAddressesToTheIpAddressGroupRequest := &AddIpAddressesToTheIpAddressGroupRequest{
		IpSetId:       util.PtrString(""),
		ClientToken:   util.PtrString(""),
		IpAddressInfo: []*TemplateIpAddressInfo{},
	}
	err := VPC_CLIENT.AddIpAddressesToTheIpAddressGroup(addIpAddressesToTheIpAddressGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddIpv6OnlyOutboundAndNoInboundPolicy(t *testing.T) {
	addIpv6OnlyOutboundAndNoInboundPolicyRequest := &AddIpv6OnlyOutboundAndNoInboundPolicyRequest{
		GatewayId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		Cidr:        util.PtrString(""),
	}
	result := &AddIpv6OnlyOutboundAndNoInboundPolicyResponse{}
	result, err := VPC_CLIENT.AddIpv6OnlyOutboundAndNoInboundPolicy(addIpv6OnlyOutboundAndNoInboundPolicyRequest)
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
func TestClient_AuthorizeRegularSecurityGroupRulesV2(t *testing.T) {
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
	authorizeRegularSecurityGroupRulesV2Request := &AuthorizeRegularSecurityGroupRulesV2Request{
		SecurityGroupId: util.PtrString(""),
		SgVersion:       util.PtrInt64(int64(0)),
		ClientToken:     util.PtrString(""),
		Rule:            Rule,
	}
	err := VPC_CLIENT.AuthorizeRegularSecurityGroupRulesV2(authorizeRegularSecurityGroupRulesV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AuthorizedEnterpriseSecurityGroupRules(t *testing.T) {
	authorizedEnterpriseSecurityGroupRulesRequest := &AuthorizedEnterpriseSecurityGroupRulesRequest{
		EnterpriseSecurityGroupId: util.PtrString(""),
		ClientToken:               util.PtrString(""),
		Rules:                     []*EnterpriseSecurityGroupRuleModel{},
	}
	err := VPC_CLIENT.AuthorizedEnterpriseSecurityGroupRules(authorizedEnterpriseSecurityGroupRulesRequest)
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
func TestClient_BatchDeleteElasticNetworkCardIntranetIp(t *testing.T) {
	batchDeleteElasticNetworkCardIntranetIpRequest := &BatchDeleteElasticNetworkCardIntranetIpRequest{
		EniId:              util.PtrString(""),
		ClientToken:        util.PtrString(""),
		PrivateIpAddresses: []*string{},
	}
	err := VPC_CLIENT.BatchDeleteElasticNetworkCardIntranetIp(batchDeleteElasticNetworkCardIntranetIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchIncreaseElasticNetworkCardIntranetIp(t *testing.T) {
	batchIncreaseElasticNetworkCardIntranetIpRequest := &BatchIncreaseElasticNetworkCardIntranetIpRequest{
		EniId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		IsIpv6:                util.PtrBool(false),
		PrivateIpAddresses:    []*string{},
		PrivateIpAddressCount: util.PtrInt32(int32(0)),
	}
	result := &BatchIncreaseElasticNetworkCardIntranetIpResponse{}
	result, err := VPC_CLIENT.BatchIncreaseElasticNetworkCardIntranetIp(batchIncreaseElasticNetworkCardIntranetIpRequest)
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
func TestClient_BindEip(t *testing.T) {
	bindEipRequest := &BindEipRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Eip:         util.PtrString(""),
	}
	err := VPC_CLIENT.BindEip(bindEipRequest)
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
func TestClient_ClosePeerToPeerConnectionToSynchronizeDns(t *testing.T) {
	closePeerToPeerConnectionToSynchronizeDnsRequest := &ClosePeerToPeerConnectionToSynchronizeDnsRequest{
		PeerConnId:  util.PtrString(""),
		Role:        util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ClosePeerToPeerConnectionToSynchronizeDns(closePeerToPeerConnectionToSynchronizeDnsRequest)
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
func TestClient_CreateAHighlyAvailableVirtualIp(t *testing.T) {
	createAHighlyAvailableVirtualIpRequest := &CreateAHighlyAvailableVirtualIpRequest{
		ClientToken:      util.PtrString(""),
		Name:             util.PtrString(""),
		SubnetId:         util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		Description:      util.PtrString(""),
	}
	result := &CreateAHighlyAvailableVirtualIpResponse{}
	result, err := VPC_CLIENT.CreateAHighlyAvailableVirtualIp(createAHighlyAvailableVirtualIpRequest)
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
func TestClient_CreateAPeerToPeerConnection(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createAPeerToPeerConnectionRequest := &CreateAPeerToPeerConnectionRequest{
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
	result := &CreateAPeerToPeerConnectionResponse{}
	result, err := VPC_CLIENT.CreateAPeerToPeerConnection(createAPeerToPeerConnectionRequest)
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
func TestClient_CreateARegularSecurityGroupV2(t *testing.T) {
	createARegularSecurityGroupV2Request := &CreateARegularSecurityGroupV2Request{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		VpcId:       util.PtrString(""),
		Desc:        util.PtrString(""),
		Rules:       []*SecurityGroupRuleModel{},
		Tags:        []*TagModel{},
	}
	result := &CreateARegularSecurityGroupV2Response{}
	result, err := VPC_CLIENT.CreateARegularSecurityGroupV2(createARegularSecurityGroupV2Request)
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
func TestClient_CreateAnIpv6Gateway(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createAnIpv6GatewayRequest := &CreateAnIpv6GatewayRequest{
		ClientToken:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		Name:            util.PtrString(""),
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Billing:         Billing,
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
		DeleteProtect:   util.PtrBool(false),
	}
	result := &CreateAnIpv6GatewayResponse{}
	result, err := VPC_CLIENT.CreateAnIpv6Gateway(createAnIpv6GatewayRequest)
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
func TestClient_CreateElasticNetworkCard(t *testing.T) {
	createElasticNetworkCardRequest := &CreateElasticNetworkCardRequest{
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
	result := &CreateElasticNetworkCardResponse{}
	result, err := VPC_CLIENT.CreateElasticNetworkCard(createElasticNetworkCardRequest)
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
func TestClient_CreateIpAddressFamily(t *testing.T) {
	createIpAddressFamilyRequest := &CreateIpAddressFamilyRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		IpVersion:   util.PtrString(""),
		IpSetIds:    []*string{},
		Description: util.PtrString(""),
	}
	result := &CreateIpAddressFamilyResponse{}
	result, err := VPC_CLIENT.CreateIpAddressFamily(createIpAddressFamilyRequest)
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
func TestClient_CreateIpAddressGroup(t *testing.T) {
	createIpAddressGroupRequest := &CreateIpAddressGroupRequest{
		ClientToken:   util.PtrString(""),
		Name:          util.PtrString(""),
		IpVersion:     util.PtrString(""),
		IpAddressInfo: []*TemplateIpAddressInfo{},
		Description:   util.PtrString(""),
	}
	result := &CreateIpAddressGroupResponse{}
	result, err := VPC_CLIENT.CreateIpAddressGroup(createIpAddressGroupRequest)
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
func TestClient_CreateIpv6GatewaySpeedLimitPolicy(t *testing.T) {
	createIpv6GatewaySpeedLimitPolicyRequest := &CreateIpv6GatewaySpeedLimitPolicyRequest{
		GatewayId:              util.PtrString(""),
		ClientToken:            util.PtrString(""),
		Ipv6Address:            util.PtrString(""),
		IngressBandwidthInMbps: util.PtrInt32(int32(0)),
		EgressBandwidthInMbps:  util.PtrInt32(int32(0)),
	}
	result := &CreateIpv6GatewaySpeedLimitPolicyResponse{}
	result, err := VPC_CLIENT.CreateIpv6GatewaySpeedLimitPolicy(createIpv6GatewaySpeedLimitPolicyRequest)
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
func TestClient_CreateNetworkDetection(t *testing.T) {
	createNetworkDetectionRequest := &CreateNetworkDetectionRequest{
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
	result := &CreateNetworkDetectionResponse{}
	result, err := VPC_CLIENT.CreateNetworkDetection(createNetworkDetectionRequest)
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
func TestClient_DeleteElasticNetworkCardAuxiliaryIp(t *testing.T) {
	deleteElasticNetworkCardAuxiliaryIpRequest := &DeleteElasticNetworkCardAuxiliaryIpRequest{
		EniId:            util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		ClientToken:      util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteElasticNetworkCardAuxiliaryIp(deleteElasticNetworkCardAuxiliaryIpRequest)
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
func TestClient_DeleteHighlyAvailableVirtualIp(t *testing.T) {
	deleteHighlyAvailableVirtualIpRequest := &DeleteHighlyAvailableVirtualIpRequest{
		HaVipId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteHighlyAvailableVirtualIp(deleteHighlyAvailableVirtualIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpAddressFamily(t *testing.T) {
	deleteIpAddressFamilyRequest := &DeleteIpAddressFamilyRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpAddressFamily(deleteIpAddressFamilyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpAddressFromIpAddressGroup(t *testing.T) {
	deleteIpAddressFromIpAddressGroupRequest := &DeleteIpAddressFromIpAddressGroupRequest{
		IpSetId:       util.PtrString(""),
		ClientToken:   util.PtrString(""),
		IpAddressInfo: []*string{},
	}
	err := VPC_CLIENT.DeleteIpAddressFromIpAddressGroup(deleteIpAddressFromIpAddressGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpAddressGroup(t *testing.T) {
	deleteIpAddressGroupRequest := &DeleteIpAddressGroupRequest{
		IpSetId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpAddressGroup(deleteIpAddressGroupRequest)
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
func TestClient_DeleteIpv6Gateway(t *testing.T) {
	deleteIpv6GatewayRequest := &DeleteIpv6GatewayRequest{
		GatewayId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpv6Gateway(deleteIpv6GatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpv6GatewaySpeedLimitPolicy(t *testing.T) {
	deleteIpv6GatewaySpeedLimitPolicyRequest := &DeleteIpv6GatewaySpeedLimitPolicyRequest{
		GatewayId:       util.PtrString(""),
		RateLimitRuleId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpv6GatewaySpeedLimitPolicy(deleteIpv6GatewaySpeedLimitPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpv6OnlyAccessPolicy(t *testing.T) {
	deleteIpv6OnlyAccessPolicyRequest := &DeleteIpv6OnlyAccessPolicyRequest{
		GatewayId:        util.PtrString(""),
		EgressOnlyRuleId: util.PtrString(""),
		ClientToken:      util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpv6OnlyAccessPolicy(deleteIpv6OnlyAccessPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteNetworkDetection(t *testing.T) {
	deleteNetworkDetectionRequest := &DeleteNetworkDetectionRequest{
		ProbeId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteNetworkDetection(deleteNetworkDetectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRegularSecurityGroupRulesV2(t *testing.T) {
	deleteRegularSecurityGroupRulesV2Request := &DeleteRegularSecurityGroupRulesV2Request{
		SecurityGroupRuleId: util.PtrString(""),
		ClientToken:         util.PtrString(""),
		SgVersion:           util.PtrInt64(int64(0)),
	}
	err := VPC_CLIENT.DeleteRegularSecurityGroupRulesV2(deleteRegularSecurityGroupRulesV2Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRegularSecurityGroupV2(t *testing.T) {
	deleteRegularSecurityGroupV2Request := &DeleteRegularSecurityGroupV2Request{
		SecurityGroupId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteRegularSecurityGroupV2(deleteRegularSecurityGroupV2Request)
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
func TestClient_ElasticNetworkCardBindingEip(t *testing.T) {
	elasticNetworkCardBindingEipRequest := &ElasticNetworkCardBindingEipRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		PublicIpAddress:  util.PtrString(""),
	}
	err := VPC_CLIENT.ElasticNetworkCardBindingEip(elasticNetworkCardBindingEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ElasticNetworkCardMountedCloudProductInstance(t *testing.T) {
	elasticNetworkCardMountedCloudProductInstanceRequest := &ElasticNetworkCardMountedCloudProductInstanceRequest{
		EniId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceId:   util.PtrString(""),
		InstanceType: util.PtrString(""),
	}
	err := VPC_CLIENT.ElasticNetworkCardMountedCloudProductInstance(elasticNetworkCardMountedCloudProductInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ElasticNetworkCardUnbindingEip(t *testing.T) {
	elasticNetworkCardUnbindingEipRequest := &ElasticNetworkCardUnbindingEipRequest{
		EniId:           util.PtrString(""),
		ClientToken:     util.PtrString(""),
		PublicIpAddress: util.PtrString(""),
	}
	err := VPC_CLIENT.ElasticNetworkCardUnbindingEip(elasticNetworkCardUnbindingEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ElasticNetworkCardUninstallationCloudProductInstance(t *testing.T) {
	elasticNetworkCardUninstallationCloudProductInstanceRequest := &ElasticNetworkCardUninstallationCloudProductInstanceRequest{
		EniId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		InstanceId:  util.PtrString(""),
	}
	err := VPC_CLIENT.ElasticNetworkCardUninstallationCloudProductInstance(elasticNetworkCardUninstallationCloudProductInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ElasticNetworkCardUpdateEnterpriseSecurityGroup(t *testing.T) {
	elasticNetworkCardUpdateEnterpriseSecurityGroupRequest := &ElasticNetworkCardUpdateEnterpriseSecurityGroupRequest{
		EniId:                      util.PtrString(""),
		ClientToken:                util.PtrString(""),
		EnterpriseSecurityGroupIds: []*string{},
	}
	err := VPC_CLIENT.ElasticNetworkCardUpdateEnterpriseSecurityGroup(elasticNetworkCardUpdateEnterpriseSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ElasticNetworkCardUpdatesRegularSecurityGroup(t *testing.T) {
	elasticNetworkCardUpdatesRegularSecurityGroupRequest := &ElasticNetworkCardUpdatesRegularSecurityGroupRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SecurityGroupIds: []*string{},
	}
	err := VPC_CLIENT.ElasticNetworkCardUpdatesRegularSecurityGroup(elasticNetworkCardUpdatesRegularSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnablePeerToPeerConnectionToSynchronizeDns(t *testing.T) {
	enablePeerToPeerConnectionToSynchronizeDnsRequest := &EnablePeerToPeerConnectionToSynchronizeDnsRequest{
		PeerConnId:  util.PtrString(""),
		Role:        util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.EnablePeerToPeerConnectionToSynchronizeDns(enablePeerToPeerConnectionToSynchronizeDnsRequest)
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
func TestClient_HighAvailabilityVirtualIpUnbindingEip(t *testing.T) {
	highAvailabilityVirtualIpUnbindingEipRequest := &HighAvailabilityVirtualIpUnbindingEipRequest{
		HaVipId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.HighAvailabilityVirtualIpUnbindingEip(highAvailabilityVirtualIpUnbindingEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HighAvailabilityVirtualIpUnbindingInstance(t *testing.T) {
	highAvailabilityVirtualIpUnbindingInstanceRequest := &HighAvailabilityVirtualIpUnbindingInstanceRequest{
		HaVipId:      util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceIds:  []*string{},
		InstanceType: util.PtrString(""),
	}
	err := VPC_CLIENT.HighAvailabilityVirtualIpUnbindingInstance(highAvailabilityVirtualIpUnbindingInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HighlyAvailableVirtualIpBindingEip(t *testing.T) {
	highlyAvailableVirtualIpBindingEipRequest := &HighlyAvailableVirtualIpBindingEipRequest{
		HaVipId:         util.PtrString(""),
		ClientToken:     util.PtrString(""),
		PublicIpAddress: util.PtrString(""),
	}
	err := VPC_CLIENT.HighlyAvailableVirtualIpBindingEip(highlyAvailableVirtualIpBindingEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HighlyAvailableVirtualIpBindingInstance(t *testing.T) {
	highlyAvailableVirtualIpBindingInstanceRequest := &HighlyAvailableVirtualIpBindingInstanceRequest{
		HaVipId:      util.PtrString(""),
		ClientToken:  util.PtrString(""),
		InstanceIds:  []*string{},
		InstanceType: util.PtrString(""),
	}
	err := VPC_CLIENT.HighlyAvailableVirtualIpBindingInstance(highlyAvailableVirtualIpBindingInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Ipv6GatewayBandwidthUpgradeAndDowngrade(t *testing.T) {
	ipv6GatewayBandwidthUpgradeAndDowngradeRequest := &Ipv6GatewayBandwidthUpgradeAndDowngradeRequest{
		GatewayId:       util.PtrString(""),
		ClientToken:     util.PtrString(""),
		BandwidthInMbps: util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.Ipv6GatewayBandwidthUpgradeAndDowngrade(ipv6GatewayBandwidthUpgradeAndDowngradeRequest)
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
func TestClient_OpenVpcRelay(t *testing.T) {
	openVpcRelayRequest := &OpenVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.OpenVpcRelay(openVpcRelayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PeerToPeerConnectionBandwidthUpgradeAndDowngrade(t *testing.T) {
	peerToPeerConnectionBandwidthUpgradeAndDowngradeRequest := &PeerToPeerConnectionBandwidthUpgradeAndDowngradeRequest{
		PeerConnId:         util.PtrString(""),
		ClientToken:        util.PtrString(""),
		NewBandwidthInMbps: util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.PeerToPeerConnectionBandwidthUpgradeAndDowngrade(peerToPeerConnectionBandwidthUpgradeAndDowngradeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PeerToPeerConnectionRenewal(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	peerToPeerConnectionRenewalRequest := &PeerToPeerConnectionRenewalRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     Billing,
	}
	err := VPC_CLIENT.PeerToPeerConnectionRenewal(peerToPeerConnectionRenewalRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PrepaidPeerToPeerConnectionUnsubscribe(t *testing.T) {
	prepaidPeerToPeerConnectionUnsubscribeRequest := &PrepaidPeerToPeerConnectionUnsubscribeRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.PrepaidPeerToPeerConnectionUnsubscribe(prepaidPeerToPeerConnectionUnsubscribeRequest)
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
func TestClient_QueryIpAddressFamilyList(t *testing.T) {
	queryIpAddressFamilyListRequest := &QueryIpAddressFamilyListRequest{
		IpVersion: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &QueryIpAddressFamilyListResponse{}
	result, err := VPC_CLIENT.QueryIpAddressFamilyList(queryIpAddressFamilyListRequest)
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
func TestClient_QueryNetworkDetectionDetails(t *testing.T) {
	queryNetworkDetectionDetailsRequest := &QueryNetworkDetectionDetailsRequest{
		ProbeId: util.PtrString(""),
	}
	result := &QueryNetworkDetectionDetailsResponse{}
	result, err := VPC_CLIENT.QueryNetworkDetectionDetails(queryNetworkDetectionDetailsRequest)
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
func TestClient_QueryNetworkDetectionList(t *testing.T) {
	queryNetworkDetectionListRequest := &QueryNetworkDetectionListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryNetworkDetectionListResponse{}
	result, err := VPC_CLIENT.QueryNetworkDetectionList(queryNetworkDetectionListRequest)
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
func TestClient_QueryTheListOfElasticNetworkCards(t *testing.T) {
	queryTheListOfElasticNetworkCardsRequest := &QueryTheListOfElasticNetworkCardsRequest{
		VpcId:            util.PtrString(""),
		InstanceId:       util.PtrString(""),
		Name:             util.PtrString(""),
		PrivateIpAddress: []*string{},
		Marker:           util.PtrString(""),
		MaxKeys:          util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfElasticNetworkCardsResponse{}
	result, err := VPC_CLIENT.QueryTheListOfElasticNetworkCards(queryTheListOfElasticNetworkCardsRequest)
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
func TestClient_QueryTheListOfEnterpriseSecurityGroups(t *testing.T) {
	queryTheListOfEnterpriseSecurityGroupsRequest := &QueryTheListOfEnterpriseSecurityGroupsRequest{
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		InstanceId: util.PtrString(""),
	}
	result := &QueryTheListOfEnterpriseSecurityGroupsResponse{}
	result, err := VPC_CLIENT.QueryTheListOfEnterpriseSecurityGroups(queryTheListOfEnterpriseSecurityGroupsRequest)
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
func TestClient_QueryTheListOfHighlyAvailableVirtualIps(t *testing.T) {
	queryTheListOfHighlyAvailableVirtualIpsRequest := &QueryTheListOfHighlyAvailableVirtualIpsRequest{
		VpcId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfHighlyAvailableVirtualIpsResponse{}
	result, err := VPC_CLIENT.QueryTheListOfHighlyAvailableVirtualIps(queryTheListOfHighlyAvailableVirtualIpsRequest)
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
func TestClient_QueryTheListOfIpAddressGroups(t *testing.T) {
	queryTheListOfIpAddressGroupsRequest := &QueryTheListOfIpAddressGroupsRequest{
		IpVersion: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfIpAddressGroupsResponse{}
	result, err := VPC_CLIENT.QueryTheListOfIpAddressGroups(queryTheListOfIpAddressGroupsRequest)
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
func TestClient_QueryTheListOfPeerConnections(t *testing.T) {
	queryTheListOfPeerConnectionsRequest := &QueryTheListOfPeerConnectionsRequest{
		VpcId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfPeerConnectionsResponse{}
	result, err := VPC_CLIENT.QueryTheListOfPeerConnections(queryTheListOfPeerConnectionsRequest)
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
func TestClient_QueryTheListOfRegularSecurityGroupsV2(t *testing.T) {
	queryTheListOfRegularSecurityGroupsV2Request := &QueryTheListOfRegularSecurityGroupsV2Request{
		Marker:           util.PtrString(""),
		MaxKeys:          util.PtrInt32(int32(0)),
		InstanceId:       util.PtrString(""),
		VpcId:            util.PtrString(""),
		SecurityGroupId:  util.PtrString(""),
		SecurityGroupIds: util.PtrString(""),
	}
	result := &QueryTheListOfRegularSecurityGroupsV2Response{}
	result, err := VPC_CLIENT.QueryTheListOfRegularSecurityGroupsV2(queryTheListOfRegularSecurityGroupsV2Request)
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
func TestClient_QueryTheListOfSpeedLimitPoliciesForIpv6Gateway(t *testing.T) {
	queryTheListOfSpeedLimitPoliciesForIpv6GatewayRequest := &QueryTheListOfSpeedLimitPoliciesForIpv6GatewayRequest{
		GatewayId: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &QueryTheListOfSpeedLimitPoliciesForIpv6GatewayResponse{}
	result, err := VPC_CLIENT.QueryTheListOfSpeedLimitPoliciesForIpv6Gateway(queryTheListOfSpeedLimitPoliciesForIpv6GatewayRequest)
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
func TestClient_QueryTheSpecifiedElasticNetworkCard(t *testing.T) {
	queryTheSpecifiedElasticNetworkCardRequest := &QueryTheSpecifiedElasticNetworkCardRequest{
		EniId: util.PtrString(""),
	}
	result := &QueryTheSpecifiedElasticNetworkCardResponse{}
	result, err := VPC_CLIENT.QueryTheSpecifiedElasticNetworkCard(queryTheSpecifiedElasticNetworkCardRequest)
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
func TestClient_QueryTheSpecifiedHighlyAvailableVirtualIp(t *testing.T) {
	queryTheSpecifiedHighlyAvailableVirtualIpRequest := &QueryTheSpecifiedHighlyAvailableVirtualIpRequest{
		HaVipId: util.PtrString(""),
	}
	result := &QueryTheSpecifiedHighlyAvailableVirtualIpResponse{}
	result, err := VPC_CLIENT.QueryTheSpecifiedHighlyAvailableVirtualIp(queryTheSpecifiedHighlyAvailableVirtualIpRequest)
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
func TestClient_QueryTheSpecifiedIpAddressFamily(t *testing.T) {
	queryTheSpecifiedIpAddressFamilyRequest := &QueryTheSpecifiedIpAddressFamilyRequest{
		IpGroupId: util.PtrString(""),
	}
	result := &QueryTheSpecifiedIpAddressFamilyResponse{}
	result, err := VPC_CLIENT.QueryTheSpecifiedIpAddressFamily(queryTheSpecifiedIpAddressFamilyRequest)
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
func TestClient_QueryTheSpecifiedIpAddressGroup(t *testing.T) {
	queryTheSpecifiedIpAddressGroupRequest := &QueryTheSpecifiedIpAddressGroupRequest{
		IpSetId: util.PtrString(""),
	}
	result := &QueryTheSpecifiedIpAddressGroupResponse{}
	result, err := VPC_CLIENT.QueryTheSpecifiedIpAddressGroup(queryTheSpecifiedIpAddressGroupRequest)
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
func TestClient_QueryTheStatusOfTheElasticNetworkCard(t *testing.T) {
	queryTheStatusOfTheElasticNetworkCardRequest := &QueryTheStatusOfTheElasticNetworkCardRequest{
		EniId: util.PtrString(""),
	}
	result := &QueryTheStatusOfTheElasticNetworkCardResponse{}
	result, err := VPC_CLIENT.QueryTheStatusOfTheElasticNetworkCard(queryTheStatusOfTheElasticNetworkCardRequest)
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
func TestClient_QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusion(t *testing.T) {
	queryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionRequest := &QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionRequest{
		GatewayId: util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
	}
	result := &QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionResponse{}
	result, err := VPC_CLIENT.QueryingTheIpv6PolicyListWithOnlyOutputAndNoInclusion(queryingTheIpv6PolicyListWithOnlyOutputAndNoInclusionRequest)
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
func TestClient_RejectPeerToPeerConnectionRequest(t *testing.T) {
	rejectPeerToPeerConnectionRequestRequest := &RejectPeerToPeerConnectionRequestRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.RejectPeerToPeerConnectionRequest(rejectPeerToPeerConnectionRequestRequest)
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
func TestClient_ReleasePeerToPeerConnection(t *testing.T) {
	releasePeerToPeerConnectionRequest := &ReleasePeerToPeerConnectionRequest{
		PeerConnId:  util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ReleasePeerToPeerConnection(releasePeerToPeerConnectionRequest)
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
func TestClient_RemoveElasticNetworkCard(t *testing.T) {
	removeElasticNetworkCardRequest := &RemoveElasticNetworkCardRequest{
		EniId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.RemoveElasticNetworkCard(removeElasticNetworkCardRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveIpAddressGroupFromIpAddressFamily(t *testing.T) {
	removeIpAddressGroupFromIpAddressFamilyRequest := &RemoveIpAddressGroupFromIpAddressFamilyRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpSetIds:    []*string{},
	}
	err := VPC_CLIENT.RemoveIpAddressGroupFromIpAddressFamily(removeIpAddressGroupFromIpAddressFamilyRequest)
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
func TestClient_RevokeRegularSecurityGroupRulesV2(t *testing.T) {
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
	revokeRegularSecurityGroupRulesV2Request := &RevokeRegularSecurityGroupRulesV2Request{
		SecurityGroupId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
		SgVersion:       util.PtrInt64(int64(0)),
		Rule:            Rule,
	}
	err := VPC_CLIENT.RevokeRegularSecurityGroupRulesV2(revokeRegularSecurityGroupRulesV2Request)
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
func TestClient_UpdateElasticNetworkCard(t *testing.T) {
	updateElasticNetworkCardRequest := &UpdateElasticNetworkCardRequest{
		EniId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateElasticNetworkCard(updateElasticNetworkCardRequest)
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
func TestClient_UpdateHighlyAvailableVirtualIp(t *testing.T) {
	updateHighlyAvailableVirtualIpRequest := &UpdateHighlyAvailableVirtualIpRequest{
		HaVipId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateHighlyAvailableVirtualIp(updateHighlyAvailableVirtualIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIpAddressFamily(t *testing.T) {
	updateIpAddressFamilyRequest := &UpdateIpAddressFamilyRequest{
		IpGroupId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateIpAddressFamily(updateIpAddressFamilyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIpAddressGroup(t *testing.T) {
	updateIpAddressGroupRequest := &UpdateIpAddressGroupRequest{
		IpSetId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateIpAddressGroup(updateIpAddressGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIpv6GatewayReleaseProtectionSwitch(t *testing.T) {
	updateIpv6GatewayReleaseProtectionSwitchRequest := &UpdateIpv6GatewayReleaseProtectionSwitchRequest{
		GatewayId:     util.PtrString(""),
		ClientToken:   util.PtrString(""),
		DeleteProtect: util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdateIpv6GatewayReleaseProtectionSwitch(updateIpv6GatewayReleaseProtectionSwitchRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIpv6GatewaySpeedLimitPolicy(t *testing.T) {
	updateIpv6GatewaySpeedLimitPolicyRequest := &UpdateIpv6GatewaySpeedLimitPolicyRequest{
		GatewayId:              util.PtrString(""),
		RateLimitRuleId:        util.PtrString(""),
		ClientToken:            util.PtrString(""),
		IngressBandwidthInMbps: util.PtrInt32(int32(0)),
		EgressBandwidthInMbps:  util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdateIpv6GatewaySpeedLimitPolicy(updateIpv6GatewaySpeedLimitPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateNetworkDetection(t *testing.T) {
	updateNetworkDetectionRequest := &UpdateNetworkDetectionRequest{
		ProbeId:     util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		DestIp:      util.PtrString(""),
		DestPort:    util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.UpdateNetworkDetection(updateNetworkDetectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePeerToPeerConnectionReleaseProtectionSwitch(t *testing.T) {
	updatePeerToPeerConnectionReleaseProtectionSwitchRequest := &UpdatePeerToPeerConnectionReleaseProtectionSwitchRequest{
		PeerConnId:    util.PtrString(""),
		DeleteProtect: util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdatePeerToPeerConnectionReleaseProtectionSwitch(updatePeerToPeerConnectionReleaseProtectionSwitchRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRegularSecurityGroupRulesV2(t *testing.T) {
	updateRegularSecurityGroupRulesV2Request := &UpdateRegularSecurityGroupRulesV2Request{
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
	err := VPC_CLIENT.UpdateRegularSecurityGroupRulesV2(updateRegularSecurityGroupRulesV2Request)
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
func TestClient_UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnections(t *testing.T) {
	updateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsRequest := &UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsRequest{
		PeerConnId:  util.PtrString(""),
		LocalIfId:   util.PtrString(""),
		Description: util.PtrString(""),
		LocalIfName: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnections(updateTheNameAndCommentsOfTheLocalInterfaceForPeerToPeerConnectionsRequest)
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
func TestClient_ViewPeerToPeerConnectionDetails(t *testing.T) {
	viewPeerToPeerConnectionDetailsRequest := &ViewPeerToPeerConnectionDetailsRequest{
		PeerConnId: util.PtrString(""),
		Role:       util.PtrString(""),
	}
	result := &ViewPeerToPeerConnectionDetailsResponse{}
	result, err := VPC_CLIENT.ViewPeerToPeerConnectionDetails(viewPeerToPeerConnectionDetailsRequest)
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
func TestClient_ViewSecurityGroupDetailsV2(t *testing.T) {
	viewSecurityGroupDetailsV2Request := &ViewSecurityGroupDetailsV2Request{
		SecurityGroupId: util.PtrString(""),
	}
	result := &ViewSecurityGroupDetailsV2Response{}
	result, err := VPC_CLIENT.ViewSecurityGroupDetailsV2(viewSecurityGroupDetailsV2Request)
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
